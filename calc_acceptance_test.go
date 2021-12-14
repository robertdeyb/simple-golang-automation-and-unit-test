package calc

import (
	"fmt"
	"os"
	"strconv"

	"github.com/cucumber/godog"
)

type CalcSuite struct {
	Suite *godog.ScenarioContext
	calc  *Calculator
}

func (cs *CalcSuite) calculatorIsCleared() error {
	cs.calc.Clear()
	return nil
}
func (cs *CalcSuite) iAdd(arg1 string) error {
	i, err := strconv.Atoi(arg1)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	cs.calc.Add(i)
	return nil
}

func (cs *CalcSuite) iPress(arg1 string) error {
	i, err := strconv.Atoi(arg1)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	cs.calc.Press(i)
	return nil
}

func (cs *CalcSuite) iSubtract(arg1 string) error {
	i, err := strconv.Atoi(arg1)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	cs.calc.Sub(i)
	return nil
}

func (cs *CalcSuite) theResultShouldBe(arg1 string) error {
	i, err := strconv.Atoi(arg1)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	result := cs.calc.Result()
	if result == i {
		return nil
	}
	return fmt.Errorf("%d doesn't match expectation: %d", result, i)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	s := CalcSuite{
		calc:  new(Calculator),
		Suite: ctx,
	}
	ctx.Step(`^calculator is cleared$`, s.calculatorIsCleared)
	ctx.Step(`^i add "([^"]*)"$`, s.iAdd)
	ctx.Step(`^i press "([^"]*)"$`, s.iPress)
	ctx.Step(`^i subtract "([^"]*)"$`, s.iSubtract)
	ctx.Step(`^the result should be "([^"]*)"$`, s.theResultShouldBe)
}
