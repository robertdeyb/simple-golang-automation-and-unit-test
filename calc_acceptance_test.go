package calc

import (
	"fmt"

	"github.com/cucumber/godog"
)

type CalcSuite struct {
	Suite *godog.ScenarioContext
	calc  *Calculator
}

var testAccount *Calculator

func (cs *CalcSuite) calculatorIsCleared() error {
	cs.calc.Clear()
	return nil
}

func (cs *CalcSuite) iAdd(arg1 int) error {
	cs.calc.Add(arg1)
	return nil
}

func (cs *CalcSuite) iMultiplyBy(arg1 int) error {
	return godog.ErrPending
}

func (cs *CalcSuite) iPress(arg1 int) error {
	cs.calc.Press(arg1)
	return nil
}

func (cs *CalcSuite) iSubtract(arg1 int) error {
	return cs.calc.MultiplyBy(arg1)
}

func (cs *CalcSuite) theResultShouldBe(arg1 int) error {
	result := cs.calc.Result()
	if result == arg1 {
		return nil
	}
	return fmt.Errorf("%d doesn't match expectation: %d", result, arg1)
}

func InitializeScenario(ctx *godog.ScenarioContext) {

	s := CalcSuite{
		calc:  new(Calculator),
		Suite: ctx,
	}
	ctx.Step(`^calculator is cleared$`, s.calculatorIsCleared)
	ctx.Step(`^i add (\d+)$`, s.iAdd)
	ctx.Step(`^i multiply by (\d+)$`, s.iMultiplyBy)
	ctx.Step(`^i press (\d+)$`, s.iPress)
	ctx.Step(`^i subtract (\d+)$`, s.iSubtract)
	ctx.Step(`^the result should be (\d+)$`, s.theResultShouldBe)
}
