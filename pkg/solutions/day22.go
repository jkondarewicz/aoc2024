package solutions

import (
	"strconv"
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
		result := calculateNthSecret(secret, 2000)
		res += result
	}
	return strconv.FormatInt(res, 10), nil
}

func (data *Day22Part02) Exec() (string, error) {
	priceChangeMaxBananas := make(map[bananaPriceChange]int)
	for _, secret := range data.Secrets {
		calculateMaxBananas(secret, 2000, priceChangeMaxBananas)
	}
	mb := 0
	for _, bananas := range priceChangeMaxBananas {
		if mb < bananas {
			mb = bananas
		}
	}
	return strconv.Itoa(mb), nil
}

var mod int64 = ((1 << 24) - 1)

func calculateNthSecret(secret int64, nth int) int64 {
	for i := 0; i < nth; i++ {
		secret = calculateNextSecret(secret)
	}
	return secret
}

func calculateMaxBananas(secret int64, nth int, priceChangeMaxBananas map[bananaPriceChange]int) {
	priceChanges := make([]struct {
		price       int
		priceChange int
	}, 0)
	cbp := make(map[bananaPriceChange]int)
	for i := 0; i < nth; i++ {
		prevSecret := secret
		secret = calculateNextSecret(secret)
		prevBananas := prevSecret % 10
		bananas := secret % 10
		priceChanges = append(priceChanges, struct {
			price       int
			priceChange int
		}{price: int(bananas), priceChange: int(bananas - prevBananas)})
		if i >= 3 {
			priceChange := bananaPriceChange{
				first:  priceChanges[i-3].priceChange,
				second: priceChanges[i-2].priceChange,
				third:  priceChanges[i-1].priceChange,
				fourth: priceChanges[i-0].priceChange,
			}
			if _, e := cbp[priceChange]; !e {
				cbp[priceChange] = priceChanges[i].price
				priceChangeMaxBananas[priceChange] = priceChangeMaxBananas[priceChange] + priceChanges[i].price 
			}
		}
	}
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
