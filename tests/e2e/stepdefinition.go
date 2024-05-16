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
	bag *bag
}

func NewStepDefinitions(t *testing.T) *StepDefinitions {
	return &StepDefinitions{
		T:   t,
		bag: nil,
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
		// setting a new bag, before every scenario
		s.bag = newbag()
		return ctx, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		// resetting the bag, before every scenario
		s.bag = nil
		return ctx, nil
	})

	ctx.Step(`^I have the first number (\d+)$`, s.iHaveTheFirstNumber)
	ctx.Step(`^I have the second number (\d+)$`, s.iHaveTheSecondNumber)
	ctx.Step(`^I add the numbers$`, s.iAddTheNumbers)
	ctx.Step(`^I subtract the numbers$`, s.iSubtractTheNumbers)
	ctx.Step(`^I multiply the numbers$`, s.iMultiplyTheNumbers)
	ctx.Step(`^I divide the numbers$`, s.iDivideTheNumbers)
	ctx.Step(`^the result should be (\d+)$`, s.theResultShouldBe)

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
	// Perform the addition operation
	s.bag.res = calculator.Add(s.bag.inp1, s.bag.inp2)
	return nil
}

func (s *StepDefinitions) iSubtractTheNumbers() error {
	if s.bag == nil {
		s.T.Fatal("bag is nil")
	}
	// Perform the subtraction operation
	s.bag.res = calculator.Subtract(s.bag.inp1, s.bag.inp2)
	return nil
}

func (s *StepDefinitions) iMultiplyTheNumbers() error {
	if s.bag == nil {
		s.T.Fatal("bag is nil")
	}
	// Perform the multiplication operation
	s.bag.res = calculator.Multiply(s.bag.inp1, s.bag.inp2)
	return nil
}

func (s *StepDefinitions) iDivideTheNumbers() error {
	if s.bag == nil {
		s.T.Fatal("bag is nil")
	}
	// Perform the division operation
	s.bag.res, s.bag.err = calculator.Divide(s.bag.inp1, s.bag.inp2)
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
