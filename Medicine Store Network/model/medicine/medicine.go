package medicine

type Purpose map[string]string

type Medicine struct {
	ID                   uint   `json:"MedicineID"`
	MedicineName         string `json:"MedicineName"`
	Purpose              string `json:"Purpose"`
	LevelOfEffectiveness uint8  `json:"LevelOfEffectiveness"`
	CostInPounds         uint   `json:"CostInPounds"`
	Availabilit          uint   `json:"Availabilit"`
}

func New(
	ID uint,
	MedicineName, Purpose string,
	LevelOfEffectiveness uint8,
	CostInPound, Availabilit uint,
) *Medicine {
	return &Medicine{
		ID,
		MedicineName,
		Purpose,
		LevelOfEffectiveness,
		CostInPound,
		Availabilit,
	}
}
