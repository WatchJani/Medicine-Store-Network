package filter

import (
	"root/model/market"
	"root/model/medicine"
)

func Find(market *[]market.Markt, filters []func(*medicine.Medicine) bool) []medicine.Medicine {

	res := []medicine.Medicine{}

	//too many foor loops
	for _, market := range *market {
		for _, medicine := range market.Medicine {
			match := true
			for _, value := range filters {
				if !value(&medicine) {
					match = false
					break
				}
			}
			if match {
				res = append(res, medicine)
			}
		}
	}

	return res
}
