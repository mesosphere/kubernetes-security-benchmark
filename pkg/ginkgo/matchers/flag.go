// Copyright © 2018 Jimmi Dyson <jimmidyson@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package matcher

import (
	"strconv"
	"strings"

	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"github.com/onsi/gomega/types"

	"github.com/mesosphere/kubernetes-security-benchmark/pkg/framework"
)

func flagID(f interface{}) string {
	return strings.Split(f.(string), "=")[0]
}

func ExpectProcess(f *framework.Framework) GomegaAssertion {
	return Expect(f.Process.CmdlineSlice())
}

func HaveFlagWithValue(name, value string) types.GomegaMatcher {
	return MatchElements(flagID, IgnoreExtras, Elements{
		name: Equal(name + "=" + value),
	})
}

func HaveFlagWithOptionalValue(name, value string) types.GomegaMatcher {
	return MatchElements(flagID, IgnoreExtras, Elements{
		name: MatchRegexp("^%s(?:=%s)?$", name, value),
	})
}

func HaveFlagWithDifferentValue(name, value string) types.GomegaMatcher {
	return Not(MatchElements(flagID, IgnoreExtras, Elements{
		name: Equal(name + "=" + value),
	}))
}

func HaveFlagThatDoesNotContainValue(name, value string) types.GomegaMatcher {
	return Not(MatchElements(flagID, IgnoreExtras, Elements{
		name: MatchRegexp("^%s=(?:.)*%s,*(?:.)*$", name, value),
	}))
}

func NotHaveFlagOrNotContainValue(name, value string) types.GomegaMatcher {
	return Or(
		Not(MatchElements(flagID, IgnoreExtras, Elements{
			name: Not(BeZero()),
		})),
		Not(MatchElements(flagID, IgnoreExtras, Elements{
			name: MatchRegexp("^%s=(?:.)*%s,*(?:.)*$", name, value),
		})),
	)
}

func HaveFlagThatContainsValue(name, value string) types.GomegaMatcher {
	return MatchElements(flagID, IgnoreExtras, Elements{
		name: MatchRegexp("^%s=(?:.)*%s,*(?:.)*$", name, value),
	})
}

func HaveFlagWithAnyValue(name string) types.GomegaMatcher {
	return MatchElements(flagID, IgnoreExtras, Elements{
		name: HavePrefix(name + "="),
	})
}

func NotHaveFlag(name string) types.GomegaMatcher {
	return Not(MatchElements(flagID, IgnoreExtras, Elements{
		name: Not(BeZero()),
	}))
}

func NotHaveFlagOrHaveFlagWithValue(name, value string) types.GomegaMatcher {
	return MatchElements(flagID, IgnoreMissing|IgnoreExtras, Elements{
		name: Equal(name + "=" + value),
	})
}

func NotHaveFlagOrHaveFlagWithDifferentValue(name, value string) types.GomegaMatcher {
	return MatchElements(flagID, IgnoreMissing|IgnoreExtras, Elements{
		name: Not(Equal(name + "=" + value)),
	})
}

func HaveFlagThatMatchesNumerically(name string, comparator string, value interface{}) types.GomegaMatcher {
	return MatchElements(flagID, IgnoreExtras, Elements{
		name: WithTransform(
			func(s string) int {
				value := strings.TrimSpace(strings.TrimPrefix(s, name+"="))
				i, _ := strconv.Atoi(value)
				return i
			},
			BeNumerically(comparator, value),
		),
	})
}
