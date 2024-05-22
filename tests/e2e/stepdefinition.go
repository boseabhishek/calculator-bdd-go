package e2e_tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/cucumber/godog"

	"github.com/boseabhishek/calculator-bdd-go/pkg/calculator"
)

type StepDefinitions struct {
	*testing.T
	cal *calculator.Calculator
	bag *bag
}

func NewStepDefinitions(t *testing.T) *StepDefinitions {
	return &StepDefinitions{
		T:   t,
		bag: nil,
		cal: nil,
	}
}

func (s *StepDefinitions) InitializeTestSuite(ctx *godog.TestSuiteContext) {
	fmt.Println("[LOG] entering InitializeTestSuite...")
	ctx.BeforeSuite(func() {
		fmt.Println("[LOG] inside BeforeSuite...")
	})
	ctx.AfterSuite(func() {
		fmt.Println("[LOG] inside AfterSuite...")
	})
	fmt.Println("[LOG] leaving InitializeTestSuite...")
}

func (s *StepDefinitions) InitializeScenario(ctx *godog.ScenarioContext) {
	fmt.Println("[LOG] entering InitializeScenario...")
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		// setting a new bag to hold input/output data, before every scenario
		s.bag = newbag()
		return ctx, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		// resetting the cal and bag, before every scenario
		s.bag = nil
		s.cal = nil 
		// So basically, we would have set the cal in ctx.Before() above (like bag) and reset here
		// but, in a more complex scenario, weher diff tasks need to be done as bg, thsi would help more!
		return ctx, nil
	})

	ctx.Step(`^I have the first number (\d+)$`, s.iHaveTheFirstNumber)
	ctx.Step(`^I have the second number (\d+)$`, s.iHaveTheSecondNumber)
	ctx.Step(`^I add the numbers$`, s.iAddTheNumbers)
	ctx.Step(`^I subtract the numbers$`, s.iSubtractTheNumbers)
	ctx.Step(`^I multiply the numbers$`, s.iMultiplyTheNumbers)
	ctx.Step(`^I divide the numbers$`, s.iDivideTheNumbers)
	ctx.Step(`^the result should be (\d+)$`, s.theResultShouldBe)
	ctx.Step(`^I set up the calculator$`, s.iSetUpTheCalculator)

	fmt.Println("[LOG] leaving InitializeScenario...")
}

type bag struct {
	inp1, inp2, res int
	err             error
}

func newbag() *bag {
	return &bag{}
}

func (s *StepDefinitions) iHaveTheFirstNumber(arg int) error {
	if s.bag == nil {
		s.T.Fatal("bag is nil")
	}
	s.bag.inp1 = arg
	return nil
}

func (s *StepDefinitions) iHaveTheSecondNumber(arg int) error {
	if s.bag == nil {
		s.T.Fatal("bag is nil")
	}
	s.bag.inp2 = arg
	return nil
}

func (s *StepDefinitions) iAddTheNumbers() error {
	if s.bag == nil {
		s.T.Fatal("bag is nil")
	}
	s.bag.res = s.cal.Add(s.bag.inp1, s.bag.inp2)
	return nil
}

func (s *StepDefinitions) iSubtractTheNumbers() error {
	if s.bag == nil {
		s.T.Fatal("bag is nil")
	}
	s.bag.res = s.cal.Subtract(s.bag.inp1, s.bag.inp2)
	return nil
}

func (s *StepDefinitions) iMultiplyTheNumbers() error {
	if s.bag == nil {
		s.T.Fatal("bag is nil")
	}
	s.bag.res = s.cal.Multiply(s.bag.inp1, s.bag.inp2)
	return nil
}

func (s *StepDefinitions) iDivideTheNumbers() error {
	if s.bag == nil {
		s.T.Fatal("bag is nil")
	}
	s.bag.res, s.bag.err = s.cal.Divide(s.bag.inp1, s.bag.inp2)
	return nil
}

func (s *StepDefinitions) theResultShouldBe(arg int) error {
	if s.bag == nil {
		s.T.Fatal("bag is nil")
	}
	if s.bag.res != arg {
		return fmt.Errorf("expected %d but got %d", arg, s.bag.res)
	}
	return nil
}

func (s *StepDefinitions) iSetUpTheCalculator() error {
	if s.cal != nil {
		s.T.Fatal("cal is already set")
	}
	s.cal = calculator.New()
	return nil
}
