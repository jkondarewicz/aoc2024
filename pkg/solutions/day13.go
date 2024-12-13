package solutions

import (
	"fmt"
	"strconv"

	aocmath "github.comjkondarewicz/aoc2024/pkg/math"
)

type Day13Part01 struct {
	Machines []Day13Machine
}
type Day13Part02 struct {
	Machines []Day13Machine
}

type Day13Machine struct {
	ButtonA         aocmath.Vertex
	ButtonB         aocmath.Vertex
	WinningPosition aocmath.Vertex
}

func (data *Day13Part01) Exec() (string, error) {
	result := 0
	for _, machine := range data.Machines {
		result += machine.calculateRequiredTokensToWin(false)
	}
	return strconv.Itoa(result), nil
}

func (data *Day13Part02) Exec() (string, error) {
	result := 0
	for _, machine := range data.Machines {
		result += machine.calculateRequiredTokensToWin(true)
	}
	return strconv.Itoa(result), nil
}

func (machine *Day13Machine) calculateRequiredTokensToWin(part2 bool) int {
	winningPosition := machine.WinningPosition
	if part2 {
		winningPosition.X += 10000000000000
		winningPosition.Y += 10000000000000
	}
	overallDeterminant := machine.ButtonA.X*machine.ButtonB.Y - machine.ButtonB.X*machine.ButtonA.Y
	aDeterminant := winningPosition.X*machine.ButtonB.Y - machine.ButtonB.X*winningPosition.Y
	bDeterminant := machine.ButtonA.X*winningPosition.Y - winningPosition.X*machine.ButtonA.Y
	if overallDeterminant == 0 {
		return 0
	}
	howManyAButtonPush := aDeterminant / overallDeterminant
	howManyBButtonPush := bDeterminant / overallDeterminant
	if howManyAButtonPush*overallDeterminant != aDeterminant ||
		howManyBButtonPush*overallDeterminant != bDeterminant ||
		(!part2 && howManyAButtonPush > 100) ||
		(!part2 && howManyBButtonPush > 100) {
		return 0
	}
	return 3*howManyAButtonPush + howManyBButtonPush
}

func (machine Day13Machine) String() string {
	return fmt.Sprintf("A=(%d, %d), B=(%d, %d), Win=(%d, %d)",
		machine.ButtonA.X,
		machine.ButtonA.Y,
		machine.ButtonB.X,
		machine.ButtonB.Y,
		machine.WinningPosition.X,
		machine.WinningPosition.Y,
	)
}
