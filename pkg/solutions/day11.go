package solutions

import (
	// "fmt"
	"strconv"
)

type Day11Part01 struct {
	Nums []int
}
type Day11Part02 struct {
	Nums []int
}

type stones struct {
	nums []int
}

type stone struct {
	number int
	count  int
}

func (allstones *stones) calculateAfterBlinksPart2(blinks int) int64 {
	currentStones := make(map[int]stone)
	for _, number := range allstones.nums {
		stone := currentStones[number]
		stone.number = number
		stone.count = stone.count + 1
		currentStones[number] = stone
	}
	for blink := 0; blink < blinks; blink++ {
		newStones := make(map[int]stone)
		for _, stone := range currentStones {
			calculateNextStones(stone.count, newStones, calculateNextNums(stone.number))
		}
		nums := make([]stone, 0)
		for _, stone := range newStones {
			nums = append(nums, stone)
		}
		currentStones = newStones
	}
	var result int64 = 0
	for _, value := range currentStones {
		result += int64(value.count)
	}
	return result
}

func calculateNextStones(multiplier int, stones map[int]stone, nums []int) {
	for _, num := range nums {
		stone, exists := stones[num]
		if !exists {
			stone.number = num
			stone.count = multiplier
		} else {
			stone.count = stone.count + multiplier
		}
		stones[num] = stone
	}
}

func calculateNextNums(num int) []int {
	nNums := make([]int, 0)
	if num == 0 {
		nNums = append(nNums, 1)
	} else if numStr := strconv.Itoa(num); len(numStr)%2 == 0 {
		leftNum, _ := strconv.Atoi(numStr[:len(numStr)/2])
		rightNum, _ := strconv.Atoi(numStr[len(numStr)/2:])
		nNums = append(nNums, leftNum)
		nNums = append(nNums, rightNum)
	} else {
		nNums = append(nNums, num*2024)
	}
	return nNums
}

func (data *Day11Part01) Exec() (string, error) {
	stones := stones{nums: data.Nums}
	result := stones.calculateAfterBlinksPart2(25)
	return strconv.FormatInt(result, 10), nil
}

func (data *Day11Part02) Exec() (string, error) {
	stones := stones{nums: data.Nums}
	result := stones.calculateAfterBlinksPart2(75)
	return strconv.FormatInt(result, 10), nil
}
