package solutions

import "fmt"

type Day01Part01 struct {
	Nums []int
}

type Day01Part02 struct {

}

func (data *Day01Part01) Exec() (string, error) {
	result := 0
	for i := 0; i < len(data.Nums); i++ {
		result += data.Nums[i]
	}
	return fmt.Sprint(result), nil
}

func (data *Day01Part02) Exec() (string, error) {
	return "", nil
}
