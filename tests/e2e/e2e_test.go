package e2e_tests

import (
	"os"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func Test_E2E(t *testing.T) {
	stepDef := NewStepDefinitions(t)
	status := godog.TestSuite{
		Name: "suite-e2e",
		// NOT related to any feature file/`Feature:` occurence
		// TestSuiteInitializer is used for state management, global setup/teardown, across the ENTIRE "Test Suite" run
		// ctx.BeforeSuite runs once before teh first scenario is encountered once for the entire test suite.
		// ctx.AfterSuite runs once after all test scenarios in the entire test suite have finished.
		TestSuiteInitializer: stepDef.InitializeTestSuite,
		// InitializeScenario sets up individual scenario steps.
		ScenarioInitializer:  stepDef.InitializeScenario,
		Options: &godog.Options{
			TestingT:  t,
			Format:    "pretty",
			Paths:     []string{"features"},
			Randomize: time.Now().UTC().UnixNano(),
			Output:    colors.Colored(os.Stdout),
		},
	}.Run()

	if status > 0 {
		t.Fail()
	}
}
