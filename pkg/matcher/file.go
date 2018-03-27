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
	"fmt"
	"os"
	"os/user"
	"syscall"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"
)

func BeOwnedBy(user, group string) types.GomegaMatcher {
	return &BeOwnedByMatcher{
		expectedUser:  user,
		expectedGroup: group,
	}
}

type BeOwnedByMatcher struct {
	expectedUser  string
	expectedGroup string
	actualUser    string
	actualGroup   string
}

func (matcher *BeOwnedByMatcher) Match(actual interface{}) (success bool, err error) {
	actualFilename, ok := actual.(string)
	if !ok {
		return false, fmt.Errorf("BeOwnedByMatcher matcher expects a file path")
	}

	fileInfo, err := os.Stat(actualFilename)
	if err != nil {
		return false, nil
	}

	uid := fmt.Sprint(fileInfo.Sys().(*syscall.Stat_t).Uid)
	u, err := user.LookupId(uid)
	if err != nil {
		return false, fmt.Errorf("Failed to lookup user with uid %s: %v", uid, err)
	}

	gid := fmt.Sprint(fileInfo.Sys().(*syscall.Stat_t).Gid)
	g, err := user.LookupGroupId(gid)
	if err != nil {
		if gid != "0" {
			return false, fmt.Errorf("Failed to lookup group with gid %s: %v", gid, err)
		}

		g = &user.Group{
			Name: "root",
		}
	}

	matcher.actualUser = u.Username
	matcher.actualGroup = g.Name

	return matcher.actualUser == matcher.expectedUser && matcher.actualGroup == matcher.expectedGroup, nil
}

func (matcher *BeOwnedByMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(
		matcher.actualUser+":"+matcher.actualGroup,
		fmt.Sprintf("file ownership of %s does not match", actual),
		matcher.expectedUser+":"+matcher.expectedGroup,
	)
}

func (matcher *BeOwnedByMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(
		matcher.actualUser+":"+matcher.actualGroup,
		fmt.Sprintf("file ownership of %s matches", actual),
		matcher.expectedUser+":"+matcher.expectedGroup,
	)
}

func HavePermissionsNumerically(comparator string, compareTo os.FileMode) types.GomegaMatcher {
	return &HavePermissionsNumericallyMatcher{
		comparator: comparator,
		compareTo:  compareTo,
		beNumericallyMatcher: &matchers.BeNumericallyMatcher{
			Comparator: comparator,
			CompareTo:  []interface{}{compareTo},
		},
	}
}

type HavePermissionsNumericallyMatcher struct {
	comparator           string
	compareTo            os.FileMode
	err                  error
	actualPermissions    os.FileMode
	beNumericallyMatcher *matchers.BeNumericallyMatcher
}

func (matcher *HavePermissionsNumericallyMatcher) Match(actual interface{}) (success bool, err error) {
	actualFilename, ok := actual.(string)
	if !ok {
		return false, fmt.Errorf("HavePermissionsNumericallyMatcher matcher expects a file path")
	}

	fileInfo, err := os.Stat(actualFilename)
	if err != nil {
		matcher.err = err
		return false, nil
	}

	matcher.actualPermissions = fileInfo.Mode()

	return matcher.beNumericallyMatcher.Match(fileInfo.Mode())
}

func (matcher *HavePermissionsNumericallyMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(
		fmt.Sprintf("%04o", matcher.actualPermissions),
		fmt.Sprintf("file permissions of %s to be %s", actual, matcher.comparator),
		fmt.Sprintf("%04o", matcher.compareTo),
	)
}

func (matcher *HavePermissionsNumericallyMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(
		fmt.Sprintf("%04o", matcher.actualPermissions),
		fmt.Sprintf("file permissions of %s not to be %s", actual, matcher.comparator),
		fmt.Sprintf("%04o", matcher.compareTo),
	)
}
