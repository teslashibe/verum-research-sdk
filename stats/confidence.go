package stats

import "math"

const z95 = 1.96

func ConfidenceInterval95(mean, stdDev float64, n int) (lower, upper float64) {
	if n <= 0 {
		return 0, 0
	}
	se := stdDev / math.Sqrt(float64(n))
	return mean - z95*se, mean + z95*se
}
