package solutions

import (
	"fmt"
	"strconv"

	aocmath "github.com/jkondarewicz/aoc2024/pkg/math"
)

type Day21Part01 struct {
	Codes Codes
}

type Day21Part02 struct {
	Codes Codes
}

func (data *Day21Part01) Exec() (string, error) {
	result := 0
	for _, c := range data.Codes.code {
		result += c.code * c.calculatePath()
	}
	return strconv.Itoa(result), nil
}

func (data *Day21Part02) Exec() (string, error) {
	return "", nil
}

func (c *code) calculatePath() int {
	codeRobot := robot{position: codeKeypad['A']}
	codeRobotMoves := make([]aocmath.Vertex, 0)
	for _, pos := range c.positions {
		np := codeRobot.position.DiffBetweenVertexes(pos)
		if np.Y == 3 {
			for y := 0; y < aocmath.Abs(np.Y); y++ {
				codeRobotMoves = append(codeRobotMoves, aocmath.NewVertex(0, np.Y).Normalize())
			}
			for x := 0; x < aocmath.Abs(np.X); x++ {
				codeRobotMoves = append(codeRobotMoves, aocmath.NewVertex(np.X, 0).Normalize())
			}
		} else {
			for x := 0; x < aocmath.Abs(np.X); x++ {
				codeRobotMoves = append(codeRobotMoves, aocmath.NewVertex(np.X, 0).Normalize())
			}
			for y := 0; y < aocmath.Abs(np.Y); y++ {
				codeRobotMoves = append(codeRobotMoves, aocmath.NewVertex(0, np.Y).Normalize())
			}
		}
		codeRobotMoves = append(codeRobotMoves, aocmath.NewVertex(0, 0))
		codeRobot.position = pos
	}
	controller1 := &controller{controllerMoves: codeRobotMoves}
	controller2 := &controller{controllerMoves: controller1.toRobotMoves(newRobot())}
	m := controller2.toRobotMoves(newRobot())
	fmt.Println(len(m))
	printMoves(m)
	return len(m)
}

func printMoves(m []aocmath.Vertex) {
	for _, move := range m {
		fmt.Printf("%c", numericKeypadPrint[move])
	}
	fmt.Println()
}

type controller struct {
	controllerMoves []aocmath.Vertex
}

func (c *controller) toRobotMoves(r *robot) []aocmath.Vertex {
	robotMoves := make([]aocmath.Vertex, 0)
	for _, cm := range c.controllerMoves {
		directionPosition := numericKeypad[cm]
		np := r.position.DiffBetweenVertexes(directionPosition)
		if r.position.Y == 0 {
			for y := 0; y < aocmath.Abs(np.Y); y++ {
				robotMoves = append(robotMoves, aocmath.NewVertex(0, np.Y).Normalize())
			}
			for x := 0; x < aocmath.Abs(np.X); x++ {
				robotMoves = append(robotMoves, aocmath.NewVertex(np.X, 0).Normalize())
			}
		} else {
			for x := 0; x < aocmath.Abs(np.X); x++ {
				robotMoves = append(robotMoves, aocmath.NewVertex(np.X, 0).Normalize())
			}
			for y := 0; y < aocmath.Abs(np.Y); y++ {
				robotMoves = append(robotMoves, aocmath.NewVertex(0, np.Y).Normalize())
			}
		}
		robotMoves = append(robotMoves, aocmath.NewVertex(0, 0))
		r.position = directionPosition
	}
	return robotMoves
}

type robot struct {
	position aocmath.Vertex
}

func newRobot() *robot {
	return &robot {position: numericKeypad[aocmath.NewVertex(0, 0)]}
}

type code struct {
	code      int
	positions []aocmath.Vertex
}
type Codes struct {
	code []code
}

func CreateCodes() *Codes {
	return &Codes{code: make([]code, 0)}
}
func (c *Codes) AddCode(codeStr string) {
	num, _ := strconv.Atoi(codeStr[0 : len(codeStr)-1])
	positions := make([]aocmath.Vertex, 0)
	for _, p := range []rune(codeStr) {
		positions = append(positions, codeKeypad[p])
	}
	c.code = append(c.code, code{num, positions})
}

var codeKeypad = map[rune]aocmath.Vertex{
	'0': aocmath.NewVertex(1, 3),
	'A': aocmath.NewVertex(2, 3),
	'1': aocmath.NewVertex(0, 2),
	'2': aocmath.NewVertex(1, 2),
	'3': aocmath.NewVertex(2, 2),
	'4': aocmath.NewVertex(0, 1),
	'5': aocmath.NewVertex(1, 1),
	'6': aocmath.NewVertex(2, 1),
	'7': aocmath.NewVertex(0, 0),
	'8': aocmath.NewVertex(1, 0),
	'9': aocmath.NewVertex(2, 0),
}
var numericKeypad = map[aocmath.Vertex]aocmath.Vertex{
	aocmath.NewVertex(0, -1): aocmath.NewVertex(1, 0), // ^
	aocmath.NewVertex(0, 0):  aocmath.NewVertex(2, 0), // A
	aocmath.NewVertex(-1, 0): aocmath.NewVertex(0, 1), // <
	aocmath.NewVertex(0, 1):  aocmath.NewVertex(1, 1), // v
	aocmath.NewVertex(1, 0):  aocmath.NewVertex(2, 1), // >
}
var numericKeypadPrint = map[aocmath.Vertex]rune{
	aocmath.NewVertex(0, -1): '^', // ^
	aocmath.NewVertex(0, 0):  'A', // A
	aocmath.NewVertex(-1, 0): '<', // <
	aocmath.NewVertex(0, 1):  'v', // v
	aocmath.NewVertex(1, 0):  '>', // >
}


//
//v<<A>>^AvA^Av<A<AA>>^AAvA<^A>AAvA^Av<A>^AA<A>Av<A<A>>^AAAvA<^A>A
//v<<A>>^AvA^Av<A<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A
//
