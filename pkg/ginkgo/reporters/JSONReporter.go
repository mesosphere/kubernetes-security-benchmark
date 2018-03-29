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

package reporters

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
)

// Produces JSON like:
//
// {
//   "name": "testsuitename",
//   "total": 123,
//   "failures": 15,
//   "skipped": 3,
//   "pending": 2,
//   "time": 125.8,
//   "specs": [{
//     "name": "[1] something [1.2] something else [1.1.3] the spec!",
//     "result": "passed|failed|skipped|pending",
//     "message": "it failed..",
//     "time": 10.5,
//     "systemOut": "full output"
//   }]
// }

type JSONTestSuite struct {
	Name     string         `json:"name"`
	Specs    []JSONTestCase `json:"specs"`
	Total    int            `json:"total"`
	Running  int            `json:"running"`
	Failures int            `json:"failures"`
	Skipped  int            `json:"skipped"`
	Pending  int            `json:"pending"`
	Time     float64        `json:"time"`
}

type SpecResult string

const (
	SpecResultFailed   SpecResult = "failed"
	SpecResultPassed   SpecResult = "passed"
	SpecResultSkipped  SpecResult = "skipped"
	SpecResultPending  SpecResult = "pending"
	SpecResultTimeout  SpecResult = "timeout"
	SpecResultPanicked SpecResult = "panicked"
	SpecResultInvalid  SpecResult = "invalid"
)

type JSONTestCase struct {
	Name              string     `json:"name"`
	Message           string     `json:"message,omitempty"`
	Result            SpecResult `json:"result"`
	Time              float64    `json:"time"`
	SystemOut         string     `json:"systemOut,omitempty"`
	Location          string     `json:"location,omitempty"`
	ComponentLocation string     `json:"componentLocation,omitempty"`
}

func NewJSONReporter(filename string) *JSONReporter {
	return &JSONReporter{
		filename: filename,
	}
}

type JSONReporter struct {
	filename string
	suite    JSONTestSuite
}

func (r *JSONReporter) SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary) {
	r.suite = JSONTestSuite{Name: summary.SuiteDescription}
}

func (r *JSONReporter) failureTypeForState(state types.SpecState) SpecResult {
	switch state {
	case types.SpecStatePassed:
		return SpecResultPassed
	case types.SpecStateFailed:
		return SpecResultFailed
	case types.SpecStateSkipped:
		return SpecResultSkipped
	case types.SpecStatePending:
		return SpecResultPending
	case types.SpecStateTimedOut:
		return SpecResultTimeout
	case types.SpecStatePanicked:
		return SpecResultPanicked
	case types.SpecStateInvalid:
		return SpecResultInvalid
	default:
		return ""
	}
}

func (r *JSONReporter) handleSetupSummary(name string, setupSummary *types.SetupSummary) {
	if setupSummary.State != types.SpecStatePassed {
		spec := JSONTestCase{
			Name:              name,
			Result:            r.failureTypeForState(setupSummary.State),
			Message:           setupSummary.Failure.Message,
			SystemOut:         setupSummary.CapturedOutput,
			Location:          setupSummary.Failure.Location.String(),
			ComponentLocation: setupSummary.Failure.ComponentCodeLocation.String(),
			Time:              setupSummary.RunTime.Seconds(),
		}

		r.suite.Specs = append(r.suite.Specs, spec)
	}
}

func (r *JSONReporter) BeforeSuiteDidRun(setupSummary *types.SetupSummary) {
	r.handleSetupSummary("BeforeSuite", setupSummary)
}

func (r *JSONReporter) SpecWillRun(specSummary *types.SpecSummary) {

}

var (
	matchLeadingAndTrailingSpaces = regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	matchDoubleSpaces             = regexp.MustCompile(`[\s\p{Zs}]{2,}`)
)

func (r *JSONReporter) normalizeSpaces(input string) string {
	final := matchLeadingAndTrailingSpaces.ReplaceAllString(input, "")
	final = matchDoubleSpaces.ReplaceAllString(final, " ")
	return final
}

func (r *JSONReporter) SpecDidComplete(specSummary *types.SpecSummary) {
	spec := JSONTestCase{
		Name:              r.normalizeSpaces(strings.Join(specSummary.ComponentTexts[1:], " ")),
		Result:            r.failureTypeForState(specSummary.State),
		Message:           specSummary.Failure.Message,
		SystemOut:         specSummary.CapturedOutput,
		Location:          specSummary.Failure.Location.String(),
		ComponentLocation: specSummary.Failure.ComponentCodeLocation.String(),
		Time:              specSummary.RunTime.Seconds(),
	}
	r.suite.Specs = append(r.suite.Specs, spec)
}

func (r *JSONReporter) AfterSuiteDidRun(setupSummary *types.SetupSummary) {
	r.handleSetupSummary("AfterSuite", setupSummary)
}

func (r *JSONReporter) SpecSuiteDidEnd(summary *types.SuiteSummary) {
	r.suite.Total = summary.NumberOfTotalSpecs
	r.suite.Running = summary.NumberOfSpecsThatWillBeRun
	r.suite.Failures = summary.NumberOfFailedSpecs
	r.suite.Skipped = summary.NumberOfSkippedSpecs
	r.suite.Pending = summary.NumberOfPendingSpecs
	r.suite.Time = summary.RunTime.Seconds()
	file, err := os.Create(r.filename)
	if err != nil {
		fmt.Printf("Failed to create JSON report file: %s\n\t%s", r.filename, err.Error())
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(r.suite)
	if err != nil {
		fmt.Printf("Failed to generate JSON report\n\t%s", err.Error())
	}
}
