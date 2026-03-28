package protocols

import (
	"context"
	"time"
)

const BaselineDays = 14
const PostDays = 14

type HealthDataProvider interface {
	GetHRVAverage(ctx context.Context, userID string, start, end time.Time) (*float64, error)
	GetRecoveryAverage(ctx context.Context, userID string, start, end time.Time) (*float64, error)
	GetRestingHRAverage(ctx context.Context, userID string, start, end time.Time) (*float64, error)
	GetSleepScoreAverage(ctx context.Context, userID string, start, end time.Time) (*float64, error)
	GetDeepSleepAverage(ctx context.Context, userID string, start, end time.Time) (*float64, error)
	GetStrainAverage(ctx context.Context, userID string, start, end time.Time) (*float64, error)
	GetVO2Max(ctx context.Context, userID string, start, end time.Time) (*float64, error)
	GetBodyComposition(ctx context.Context, userID string, at time.Time) (*BodyComposition, error)
	GetNutritionAverage(ctx context.Context, userID string, start, end time.Time) (*NutritionAverage, error)
	GetLabValue(ctx context.Context, userID string, biomarker string, start, end time.Time) (*float64, error)
}

func ComputeOutcomes(ctx context.Context, provider HealthDataProvider, userID string, protocol Protocol) (*OutcomeSnapshot, error) {
	baselineStart := protocol.StartDate.AddDate(0, 0, -BaselineDays)
	baselineEnd := protocol.StartDate.AddDate(0, 0, -1)

	duringStart := protocol.StartDate
	duringEnd := time.Now()
	if protocol.EndDate != nil {
		duringEnd = *protocol.EndDate
	}

	baseline, err := fetchPeriodMetrics(ctx, provider, userID, baselineStart, baselineEnd)
	if err != nil {
		return nil, err
	}

	during, err := fetchPeriodMetrics(ctx, provider, userID, duringStart, duringEnd)
	if err != nil {
		return nil, err
	}

	snapshot := &OutcomeSnapshot{
		Baseline: baseline,
		During:   during,
		Deltas:   computeDeltas(baseline, during),
	}

	if protocol.EndDate != nil {
		postStart := protocol.EndDate.AddDate(0, 0, 1)
		postEnd := protocol.EndDate.AddDate(0, 0, PostDays)
		if postEnd.Before(time.Now()) {
			post, err := fetchPeriodMetrics(ctx, provider, userID, postStart, postEnd)
			if err == nil {
				snapshot.Post = post
			}
		}
	}

	return snapshot, nil
}

func fetchPeriodMetrics(ctx context.Context, p HealthDataProvider, userID string, start, end time.Time) (*PeriodMetrics, error) {
	m := &PeriodMetrics{
		PeriodStart: start,
		PeriodEnd:   end,
	}

	m.HRVAvg, _ = p.GetHRVAverage(ctx, userID, start, end)
	m.RestingHR, _ = p.GetRestingHRAverage(ctx, userID, start, end)
	m.Recovery, _ = p.GetRecoveryAverage(ctx, userID, start, end)
	m.SleepScore, _ = p.GetSleepScoreAverage(ctx, userID, start, end)
	m.DeepSleep, _ = p.GetDeepSleepAverage(ctx, userID, start, end)
	m.StrainAvg, _ = p.GetStrainAverage(ctx, userID, start, end)
	m.VO2Max, _ = p.GetVO2Max(ctx, userID, start, end)

	if comp, err := p.GetBodyComposition(ctx, userID, end); err == nil && comp != nil {
		m.LeanMass = &comp.LeanMassKg
		m.FatMass = &comp.FatMassKg
		m.BodyFatPct = &comp.BodyFatPct
	}

	if nutr, err := p.GetNutritionAverage(ctx, userID, start, end); err == nil && nutr != nil {
		m.CalorieAvg = &nutr.CaloriesAvg
		m.ProteinAvg = &nutr.ProteinAvg
	}

	return m, nil
}

func computeDeltas(baseline, during *PeriodMetrics) OutcomeDeltas {
	var d OutcomeDeltas
	d.HRVPctChange = pctChange(baseline.HRVAvg, during.HRVAvg)
	d.RecoveryPctChange = pctChange(baseline.Recovery, during.Recovery)
	d.RestingHRChange = absChange(baseline.RestingHR, during.RestingHR)
	d.SleepPctChange = pctChange(baseline.SleepScore, during.SleepScore)
	d.DeepSleepPctChange = pctChange(baseline.DeepSleep, during.DeepSleep)
	d.StrainPctChange = pctChange(baseline.StrainAvg, during.StrainAvg)
	d.VO2MaxPctChange = pctChange(baseline.VO2Max, during.VO2Max)
	d.LeanMassPctChange = pctChange(baseline.LeanMass, during.LeanMass)
	d.FatMassPctChange = pctChange(baseline.FatMass, during.FatMass)
	d.BodyFatPctChange = pctChange(baseline.BodyFatPct, during.BodyFatPct)
	d.CaloriePctChange = pctChange(baseline.CalorieAvg, during.CalorieAvg)
	d.ProteinPctChange = pctChange(baseline.ProteinAvg, during.ProteinAvg)
	return d
}

func pctChange(before, after *float64) *float64 {
	if before == nil || after == nil || *before == 0 {
		return nil
	}
	v := ((*after - *before) / *before) * 100
	return &v
}

func absChange(before, after *float64) *float64 {
	if before == nil || after == nil {
		return nil
	}
	v := *after - *before
	return &v
}
