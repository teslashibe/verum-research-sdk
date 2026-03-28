package stats

import (
	"math"
	"sort"

	"github.com/verum-research/sdk/anonymize"
)

func AggregateCohort(compoundID, compoundName string, contributions []anonymize.AnonymizedContribution) CohortAnalysis {
	analysis := CohortAnalysis{
		CompoundID:   compoundID,
		CompoundName: compoundName,
		SampleSize:   len(contributions),
		Metrics:      make(map[string]Metric),
	}

	metricValues := make(map[string][]float64)
	for _, c := range contributions {
		for key, val := range c.OutcomeDeltas {
			metricValues[key] = append(metricValues[key], val)
		}
	}

	for name, values := range metricValues {
		if len(values) < 2 {
			continue
		}
		mean := mean(values)
		sd := stdDev(values, mean)
		n := len(values)
		ci95Lower, ci95Upper := ConfidenceInterval95(mean, sd, n)

		analysis.Metrics[name] = Metric{
			Name:       name,
			Mean:       mean,
			Median:     median(values),
			StdDev:     sd,
			N:          n,
			EffectSize: CohenD(mean, sd),
			CI95Lower:  ci95Lower,
			CI95Upper:  ci95Upper,
			PValue:     PValue(mean, sd, n),
		}
	}

	return analysis
}

func mean(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

func median(values []float64) float64 {
	sorted := make([]float64, len(values))
	copy(sorted, values)
	sort.Float64s(sorted)
	n := len(sorted)
	if n%2 == 0 {
		return (sorted[n/2-1] + sorted[n/2]) / 2
	}
	return sorted[n/2]
}

func stdDev(values []float64, mean float64) float64 {
	sum := 0.0
	for _, v := range values {
		diff := v - mean
		sum += diff * diff
	}
	return math.Sqrt(sum / float64(len(values)-1))
}
