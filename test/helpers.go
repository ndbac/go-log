package test

import (
	"math/rand"

	util "github.com/ndbac/go-log/src/utils"
)

func RandomOwner() string {
	return util.RandomString(6)
}

func RandomMoney() int64 {
	return util.RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "VND", "JPY", "NTD", "GBP"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
