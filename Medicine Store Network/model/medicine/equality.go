package medicine

func LevelOfEffectiveness(level uint8) func(*Medicine) bool {
	return func(m *Medicine) bool {
		return m.LevelOfEffectiveness == level
	}
}
