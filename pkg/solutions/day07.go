package solutions

import (
	"fmt"
	"strconv"
)

type Day07Equation struct {
	Result int
	Nums   []int
}

func (equation Day07Equation) canBeAchieved(index int, currentResult int, concatetionAllowed bool) bool {
	if index >= len(equation.Nums) {
		if currentResult == equation.Result {
			return true
		} else {
			return false
		}
	}
	num := equation.Nums[index]
	result1 := equation.canBeAchieved(index+1, num*currentResult, concatetionAllowed)
	result2 := equation.canBeAchieved(index+1, num+currentResult, concatetionAllowed)
	result3 := false
	if concatetionAllowed {
		concated, _ := strconv.Atoi(fmt.Sprintf("%d%d", currentResult, num))
		result3 = equation.canBeAchieved(index+1, concated, concatetionAllowed)
	}
	return result1 || result2 || result3
}

type Day07Part01 struct {
	Equations []Day07Equation
}
type Day07Part02 struct {
	Equations []Day07Equation
}

func (data *Day07Part01) Exec() (string, error) {
	var result int64 = 0
	for _, equation := range data.Equations {
		if equation.canBeAchieved(1, equation.Nums[0], false) {
			result += int64(equation.Result)
		}
	}
	return strconv.FormatInt(int64(int(result)), 10), nil
}

func (data *Day07Part02) Exec() (string, error) {
	var result int64 = 0
	for _, equation := range data.Equations {
		if equation.canBeAchieved(1, equation.Nums[0], true) {
			result += int64(equation.Result)
		}
	}
	return strconv.FormatInt(int64(int(result)), 10), nil
}
