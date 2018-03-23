ROOTPKG := $(shell go list .)
PARENTPKG := $(shell dirname $(ROOTPKG))
BINARYNAME := $(shell basename $(ROOTPKG)).test

export GOPATH := $(CURDIR)/.go
export GOBIN := $(GOPATH)/bin

DCOS_TASK ?= kube-apiserver-0-instance
GINKGO_FOCUS ?= ' 1.1 '
TEST_OUTPUT_FILE ?= junit.$(DCOS_TASK).xml

.PHONY: test
test: ginkgo.build
	@$(GOBIN)/ginkgo -notify

.PHONY: watch
watch: ginkgo.build
	@$(GOBIN)/ginkgo watch -notify

.PHONY: build
build: out/$(BINARYNAME)

out/$(BINARYNAME): ginkgo.build
	@go test -c -o $(CURDIR)/out/$(BINARYNAME) .

.PHONY: vendor
vendor: .vendor

.vendor: .gopath.prepare $(GOBIN)/dep Gopkg.toml Gopkg.lock
	@GOPATH= $(GOBIN)/dep ensure
	@touch $@

.PHONY: gopath
gopath: .gopath

.gopath: .gopath.prepare vendor
	@ln -s $(CURDIR)/vendor/github.com/* $(GOPATH)/src/github.com 2>/dev/null || true
	@ln -s $(shell find $(CURDIR)/vendor -maxdepth 1 -type d ! -wholename $(CURDIR)/vendor ! -wholename $(CURDIR)/vendor/github.com) $(GOPATH)/src/ 2>/dev/null || true
	@touch $@

.gopath.prepare:
	@mkdir -p $(GOPATH)/src/$(PARENTPKG)
	@ln -s $(CURDIR) $(GOPATH)/src/$(ROOTPKG)
	@touch $@

.PHONY: ginkgo.bootstrap
ginkgo.bootstrap: ginkgo.build
ifndef BOOTSTRAP_DIR
	$(error "Missing BOOTSTRAP_DIR variable - use make BOOTSTRAP_DIR=path/to/tests ginkgo.bootstrap")
endif
	@mkdir -p $(BOOTSTRAP_DIR) && cd $(BOOTSTRAP_DIR) && $(GOBIN)/ginkgo bootstrap

.PHONY: ginkgo.generate
ginkgo.generate: 
ifndef SUBJECT
	$(error "Missing SUBJECT variable - use make SUBJECT=subject ginkgo.generate")
endif
	@$(GOBIN)/ginkgo generate $(SUBJECT)

.PHONY: ginkgo.build
ginkgo.build: $(GOBIN)/ginkgo

$(GOBIN)/ginkgo: gopath
	@go install github.com/onsi/ginkgo/ginkgo

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
	@dcos task exec $(DCOS_TASK) ./$(BINARYNAME) -ginkgo.focus=$(GINKGO_FOCUS) -ginkgo.noisySkippings=false -ginkgo.noisyPendings=false -ginkgo.noColor > $(CURDIR)/results/$(TEST_OUTPUT_FILE_PREFIX).txt || true
	@echo "Retrieving junit results from $(DCOS_TASK) into $(CURDIR)/results/$(TEST_OUTPUT_FILE_PREFIX).xml"
	@dcos task exec -i $(DCOS_TASK) bash -c "cat junit.xml" > $(CURDIR)/results/$(TEST_OUTPUT_FILE_PREFIX).xml

.PHONY: test.dcos.apiserver
test.dcos.apiserver:
	$(MAKE) DCOS_TASK=kube-apiserver-0-instance GINKGO_FOCUS=' 1.1 ' TEST_OUTPUT_FILE_PREFIX=1.1-APIServer test.dcos.remote

.PHONY: test.dcos.scheduler
test.dcos.scheduler: 
test.dcos.scheduler: 
test.dcos.scheduler:
	$(MAKE) DCOS_TASK=kube-scheduler-0-instance GINKGO_FOCUS=' 1.2 ' TEST_OUTPUT_FILE_PREFIX=1.2-Scheduler test.dcos.remote

.PHONY: test.dcos.controller-manager
test.dcos.controller-manager:
	$(MAKE) DCOS_TASK=kube-controller-manager-0-instance GINKGO_FOCUS=' 1.3 ' TEST_OUTPUT_FILE_PREFIX=1.3-ControllerManger test.dcos.remote