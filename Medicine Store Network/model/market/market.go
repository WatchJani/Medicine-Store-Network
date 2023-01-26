package market

import "root/model/medicine"

type Markt struct {
	Name     string              `json:"Name"`
	Medicine []medicine.Medicine `json:"Medicine"`
}

func Empty() *[]Markt {
	return &[]Markt{}
}
