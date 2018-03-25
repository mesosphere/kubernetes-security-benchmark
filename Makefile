# Copyright © 2018 Jimmi Dyson <jimmidyson@gmail.com>
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

ROOTPKG := $(shell go list .)
PARENTPKG := $(shell dirname $(ROOTPKG))
BINARYNAME := $(shell basename $(ROOTPKG))

export GOPATH := $(CURDIR)/.go
export GOBIN := $(GOPATH)/bin
export CGO_ENABLED=0

BUILD_DATE := $(shell date -u)
VERSION ?= $(shell git describe --match 'v[0-9]*' --dirty --always)

DCOS_TASK ?= kube-apiserver-0-instance
CIS_FOCUS ?= api-server

.PHONY: build
build: out/$(BINARYNAME)

out/$(BINARYNAME): .vendor $(shell find ! -path '*/.go/*' -type f -name '*.go')
	@cd $(GOPATH)/src/$(ROOTPKG) && \
		GOOS=linux GOARCH=amd64 go build \
			-tags netgo \
			-ldflags "-extldflags \"-static\" -X $(ROOTPKG)/pkg/version.AppVersion=$(VERSION) -X '$(ROOTPKG)/pkg/version.BuildDate=$(BUILD_DATE)'" \
			-o $(CURDIR)/out/$(BINARYNAME) .

.PHONY: vendor
vendor: .vendor

.vendor: .gopath.prepare $(GOBIN)/dep Gopkg.toml Gopkg.lock
	@cd $(GOPATH)/src/$(ROOTPKG) && $(GOBIN)/dep ensure
	@touch $@

.gopath: .vendor
	@touch $@

.gopath.prepare:
	@mkdir -p $(GOPATH)/src/$(PARENTPKG)
	@ln -s $(CURDIR) $(GOPATH)/src/$(ROOTPKG)
	@touch $@

.PHONY: dep
dep: $(GOBIN)/dep

$(GOBIN)/dep:
	@go get github.com/golang/dep/cmd/dep

.PHONY: clean
clean:
	@rm -rf vendor .vendor $(GOPATH) .gopath .gopath.prepare results out

.PHONY: test.dcos
test.dcos: build $(addprefix test.dcos.,apiserver scheduler controller-manager)

.PHONY: test.dcos.remote
test.dcos.remote: build
ifndef DCOS_TASK
	$(error "Missing DCOS_TASK variable")
endif
	@mkdir -p $(CURDIR)/results/
	@echo "Copying binary to $(DCOS_TASK)"
	@cat out/$(BINARYNAME) | dcos task exec -i $(DCOS_TASK) bash -c "cat > $(BINARYNAME)"
	@dcos task exec $(DCOS_TASK) chmod +x $(BINARYNAME)
	@echo "Running tests on $(DCOS_TASK)"
	@dcos task exec $(DCOS_TASK) ./$(BINARYNAME) cis $(CIS_FOCUS) > $(CURDIR)/results/$(CIS_FOCUS).txt || true
	@echo "Retrieving junit results from $(DCOS_TASK) into $(CURDIR)/results/junit.$(CIS_FOCUS).xml"
	@dcos task exec -i $(DCOS_TASK) bash -c "cat junit.xml" > $(CURDIR)/results/junit.$(CIS_FOCUS).xml

.PHONY: test.dcos.apiserver
test.dcos.apiserver:
	@$(MAKE) DCOS_TASK=kube-apiserver-0-instance CIS_FOCUS=api-server test.dcos.remote

.PHONY: test.dcos.scheduler
test.dcos.scheduler:
	@$(MAKE) DCOS_TASK=kube-scheduler-0-instance CIS_FOCUS=scheduler test.dcos.remote

.PHONY: test.dcos.controller-manager
test.dcos.controller-manager:
	@$(MAKE) DCOS_TASK=kube-controller-manager-0-instance CIS_FOCUS=controller-manager test.dcos.remote