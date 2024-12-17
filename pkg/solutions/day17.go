package solutions

import (
	"fmt"
	"math"
	"strconv"

	"github.com/jkondarewicz/aoc2024/pkg/utils"
)

type Day17Part01 struct {
	Program Day17Program
}

type Day17Part02 struct {
	Program Day17Program
}

type Day17Program struct {
	RegisterA int
	RegisterB int
	RegisterC int
	Program   []int
	pointer   int
	out       []int
}

func (p Day17Program) newRegisterA(val int) Day17Program {
	p.RegisterA = val
	return p
}

func programOutput(programOutput []int) string {
	out := strconv.Itoa(programOutput[0])
	for i := 1; i < len(programOutput); i++ {
		out += fmt.Sprintf(",%d", programOutput[i])
	}
	return out
}

func (data *Day17Part01) Exec() (string, error) {
	p := calculateOutput(int64(data.Program.RegisterA))
	return programOutput(p), nil
}

func (data *Day17Part02) Exec() (string, error) {
	possibleRegisterA := utils.NewSet[int]().Add(0)
	for prog := len(data.Program.Program) - 1; prog >= 0; prog-- {
		searchNum := data.Program.Program[prog]
		newPossibleRegisterA := utils.NewSet[int]()
		for _, registerA := range possibleRegisterA.Get() {
			for i := 0; i < 8; i++ {
				newRegisterAValue := (registerA << 3) + i
				if programOutput, _ := getProgramOutput(int64(newRegisterAValue)); programOutput == searchNum {
					newPossibleRegisterA.Add(newRegisterAValue)
				}
			}
		}
		possibleRegisterA = newPossibleRegisterA
	}
	lowest := math.MaxInt64
	possibleRegisterA.ForEach(func (in int)  {
		if lowest > in {
			lowest = in
		}
	})
	return strconv.Itoa(lowest), nil
}

func calculateOutput(aRegister int64) []int {
	out := make([]int, 0)
	for aRegister != 0 {
		b, a := getProgramOutput(aRegister)
		out = append(out, b)
		aRegister = a
	}
	return out
}

func getProgramOutput(aRegister int64) (int, int64) {
	partial := 1 ^ (aRegister % 8)
	b := (5 ^ partial) ^ (aRegister >> partial)
	return int(b % 8), aRegister >> 3
}
