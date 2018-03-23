package matcher

import (
	"strings"

	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"github.com/onsi/gomega/types"
)

func flagID(f interface{}) string {
	return strings.Split(f.(string), "=")[0]
}

func HaveFlagWithValue(name, value string) types.GomegaMatcher {
	return MatchElements(flagID, IgnoreExtras, Elements{
		name: Equal(name + "=" + value),
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
			func(s string) string {
				return strings.TrimPrefix(s, name+"=")
			},
			BeNumerically(comparator, value),
		),
	})
}
