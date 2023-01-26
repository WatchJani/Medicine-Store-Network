package main

import (
	"fmt"
	"root/filter"
	"root/model/market"
	"root/model/medicine"
	"root/os"
	"testing"
)

func TestFind(t *testing.T) {
	market := market.Empty()
	os.ReadFile(market, "./data/medicineMarket.json")

	filters := []func(*medicine.Medicine) bool{
		medicine.LevelOfEffectiveness(2),
	}

	LevelOfEffectiveness := filter.Find(market, filters)

	if len(LevelOfEffectiveness) == 2 {
		for _, value := range LevelOfEffectiveness {
			fmt.Println(value)
		}
	} else {
		t.Error("Something is wrong")
	}
}
