package solutions

import (
	"strconv"
)
type MulSystem struct {
	stateEnabled bool
	active bool
	result int64
}
type MulSystemCommand interface {
	Exec(mulSystem *MulSystem)
}

type ActivateMul struct {}
func (activate ActivateMul) Exec(mulSystem *MulSystem) {
	mulSystem.active = true
}

type DeactivateMul struct {}
func (deactivate DeactivateMul) Exec(mulSystem *MulSystem) {
	mulSystem.active = false
}

type PerformCalcMul struct {
	X int64
	Y int64
}
func (performCalc PerformCalcMul) Exec(mulSystem *MulSystem) {
	if (mulSystem.active && mulSystem.stateEnabled) || !mulSystem.stateEnabled {
		mulSystem.result += performCalc.X * performCalc.Y
	}
}
type Day03Part01 struct {
	MulCommands []MulSystemCommand
}

type Day03Part02 struct {
	MulCommands []MulSystemCommand
}

func (data *Day03Part01) Exec() (string, error) {
	mulSystem := MulSystem { stateEnabled: false, result: 0 }
	for _, mulCommand := range data.MulCommands {
		mulCommand.Exec(&mulSystem)
	}
	return strconv.FormatInt(mulSystem.result, 10), nil
}

func (data *Day03Part02) Exec() (string, error) {
	mulSystem := MulSystem { stateEnabled: true, active: true, result: 0 }
	for _, mulCommand := range data.MulCommands {
		mulCommand.Exec(&mulSystem)
	}
	return strconv.FormatInt(mulSystem.result, 10), nil
}
