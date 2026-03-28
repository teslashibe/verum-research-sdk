package anonymize

import (
	"fmt"

	"github.com/teslashibe/verum-research-sdk/protocols"
	"github.com/teslashibe/verum-research-sdk/schema"
)

func Prepare(protocol protocols.Protocol, outcomes protocols.OutcomeDeltas, demographics Demographics, config Config) (*AnonymizedContribution, error) {
	contribution := &AnonymizedContribution{
		CompoundID:     protocol.CompoundID,
		CompoundName:   protocol.CompoundName,
		DoseBucket:     BucketDose(protocol.Dose, protocol.DoseUnit),
		Frequency:      protocol.Frequency,
		Route:          protocol.Route,
		DurationBucket: BucketDuration(protocol.DurationDays()),
		GoalCategory:   protocol.GoalCategory,
		AgeRange:       demographics.AgeRange,
		Sex:            demographics.Sex,
		BMIRange:       demographics.BMIRange,
		OutcomeDeltas:  flattenDeltas(outcomes),
		SelfRating:     outcomes.SelfRating,
		SchemaVersion:  schema.Version,
		ContributedAt:  MonthPrecision(),
	}

	contribution.Hash = ContributionHash(contribution)

	if err := ValidateNoLeaks(contribution); err != nil {
		return nil, fmt.Errorf("anonymization validation failed: %w", err)
	}

	return contribution, nil
}

func flattenDeltas(d protocols.OutcomeDeltas) map[string]float64 {
	m := make(map[string]float64)
	if d.HRVPctChange != nil {
		m["hrv_pct_change"] = *d.HRVPctChange
	}
	if d.RecoveryPctChange != nil {
		m["recovery_pct_change"] = *d.RecoveryPctChange
	}
	if d.RestingHRChange != nil {
		m["resting_hr_change"] = *d.RestingHRChange
	}
	if d.SleepPctChange != nil {
		m["sleep_pct_change"] = *d.SleepPctChange
	}
	if d.DeepSleepPctChange != nil {
		m["deep_sleep_pct_change"] = *d.DeepSleepPctChange
	}
	if d.StrainPctChange != nil {
		m["strain_pct_change"] = *d.StrainPctChange
	}
	if d.VO2MaxPctChange != nil {
		m["vo2_max_pct_change"] = *d.VO2MaxPctChange
	}
	if d.LeanMassPctChange != nil {
		m["lean_mass_pct_change"] = *d.LeanMassPctChange
	}
	if d.FatMassPctChange != nil {
		m["fat_mass_pct_change"] = *d.FatMassPctChange
	}
	if d.BodyFatPctChange != nil {
		m["body_fat_pct_change"] = *d.BodyFatPctChange
	}
	if d.CaloriePctChange != nil {
		m["calorie_pct_change"] = *d.CaloriePctChange
	}
	if d.ProteinPctChange != nil {
		m["protein_pct_change"] = *d.ProteinPctChange
	}
	for _, ld := range d.LabDeltas {
		key := fmt.Sprintf("lab_%s_pct_change", ld.BiomarkerName)
		m[key] = ld.PctChange
	}
	return m
}
