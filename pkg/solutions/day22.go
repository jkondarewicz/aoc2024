package solutions

import (
	"fmt"
	"strconv"

	"github.com/jkondarewicz/aoc2024/pkg/utils"
)

type Day22Part01 struct {
	Secrets []int64
}

type Day22Part02 struct {
	Secrets []int64
}

func (data *Day22Part01) Exec() (string, error) {
	var res int64 = 0
	for _, secret := range data.Secrets {
		result, _ := calculateNthSecret(secret, 2000)
		res += result
	}
	return strconv.FormatInt(res, 10), nil
}

func (data *Day22Part02) Exec() (string, error) {
	bananaPatterns := make([]bananaPrices, 0)
	for _, secret := range data.Secrets {
		_, bp := calculateNthSecret(secret, 2000)
		bananaPatterns = append(bananaPatterns, bp)
	}
	allMaxBananas := 0
	checkedChanges := utils.NewSet[bananaPriceChange]()
	winningPattern := bananaPriceChange{}
	for first, bp1 := range bananaPatterns {
		for pattern, price := range bp1 {
			if checkedChanges.Exists(pattern) {
				continue
			}
			maxBananas := price
			for second, bp2 := range bananaPatterns {
				if first == second {
					continue
				}
				var maxBananasGot int = bp2[pattern]
				maxBananas += maxBananasGot
			}
			if maxBananas > allMaxBananas {
				allMaxBananas = maxBananas
				winningPattern = pattern
			}
			checkedChanges.Add(pattern)
		}
	}

	fmt.Println(winningPattern)
	return strconv.Itoa(allMaxBananas), nil
}

var mod int64 = ((1 << 24) - 1)
func calculateNthSecret(secret int64, nth int) (int64, bananaPrices) {
	priceChanges := make([][]int, 0)
	for i := 0; i < nth; i++ {
		prevSecret := secret
		secret = calculateNextSecret(secret)
		prevBananas := prevSecret % 10
		bananas := secret % 10
		priceChanges = append(priceChanges, []int{int(bananas), int(bananas - prevBananas)})
	}
	prices := make(map[bananaPriceChange]int)
	for i := 0; i < nth-3; i++ {
		priceChange := bananaPriceChange{
			first:  priceChanges[i][1],
			second: priceChanges[i+1][1],
			third:  priceChanges[i+2][1],
			fourth: priceChanges[i+3][1],
		}
		newPrice := priceChanges[i+3][0]
		prevPrice := prices[priceChange]
		if prevPrice < newPrice {
			prices[priceChange] = newPrice
		}
	}
	return secret, bananaPrices(prices)
}

func calculateNextSecret(secret int64) int64 {
	ns := ((secret << 6) ^ secret) & mod
	ns = ((ns >> 5) ^ ns) & mod
	ns = ((ns << 11) ^ ns) & mod
	return ns
}

type bananaPrices map[bananaPriceChange]int
type bananaPriceChange struct {
	first, second, third, fourth int
}
