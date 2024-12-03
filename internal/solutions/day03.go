package solutions

import (
	"strconv"
)

type MulCommand struct {
	X int64
	Y int64
}
func (command MulCommand) multiply() int64 {
	return command.X * command.Y
}

type Day03Part01 struct {
	MulCommands []MulCommand
}

type Day03Part02 struct {
	MulCommands []MulCommand
}

func (data *Day03Part01) Exec() (string, error) {
	var result int64
	result = 0
	for _, mulCommand := range data.MulCommands {
		result += mulCommand.multiply()
	}
	return strconv.FormatInt(result, 10), nil
}

func (data *Day03Part02) Exec() (string, error) {
	var result int64
	result = 0
	for _, mulCommand := range data.MulCommands {
		result += mulCommand.multiply()
	}
	return strconv.FormatInt(result, 10), nil
}
