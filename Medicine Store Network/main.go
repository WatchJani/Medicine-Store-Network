package main

import (
	"fmt"
	"root/filter"
	"root/model/market"
	"root/model/medicine"
	"root/os"
)

func main() {

	market := market.Empty()
	os.ReadFile(market, "./data/medicineMarket.json")

	filters := []func(*medicine.Medicine) bool{
		medicine.LevelOfEffectiveness(2),
	}

	LevelOfEffectiveness := filter.Find(market, filters)

	fmt.Println(LevelOfEffectiveness)
}
