package stats

import "math"

func CohenD(mean, stdDev float64) float64 {
	if stdDev == 0 {
		return 0
	}
	return mean / stdDev
}

func InterpretCohenD(d float64) string {
	abs := math.Abs(d)
	switch {
	case abs < 0.2:
		return "negligible"
	case abs < 0.5:
		return "small"
	case abs < 0.8:
		return "medium"
	default:
		return "large"
	}
}
