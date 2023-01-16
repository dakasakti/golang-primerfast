package feature

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func CountMoney(nominal int64) map[string]int64 {
	denom := []int64{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	output := make(map[string]int64)

	for i, val := range denom {
		if nominal >= val {
			output[formatRupiah(val)] += nominal / val
			nominal %= val
		}

		if len(denom)-1 == i && nominal > 0 {
			output[formatRupiah(val)]++
			nominal -= nominal
		}
	}

	return output
}

func formatRupiah(nominal int64) string {
	id := language.Indonesian
	return message.NewPrinter(id).Sprintf("Rp. %d", nominal)
}
