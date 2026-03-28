package protocols

import "time"

type ProtocolStatus string

const (
	StatusActive    ProtocolStatus = "active"
	StatusCompleted ProtocolStatus = "completed"
	StatusPaused    ProtocolStatus = "paused"
)

type Protocol struct {
	CompoundID   string         `json:"compound_id"`
	CompoundName string         `json:"compound_name,omitempty"`
	Dose         float64        `json:"dose"`
	DoseUnit     string         `json:"dose_unit"`
	Frequency    string         `json:"frequency"`
	Route        string         `json:"route"`
	StartDate    time.Time      `json:"start_date"`
	EndDate      *time.Time     `json:"end_date,omitempty"`
	Status       ProtocolStatus `json:"status"`
	Goal         string         `json:"goal,omitempty"`
	GoalCategory string         `json:"goal_category,omitempty"`
}

func (p Protocol) DurationDays() int {
	end := time.Now()
	if p.EndDate != nil {
		end = *p.EndDate
	}
	return int(end.Sub(p.StartDate).Hours() / 24)
}

type SideEffect struct {
	EffectType   string     `json:"effect_type"`
	Severity     int        `json:"severity"` // 1-5
	Description  string     `json:"description,omitempty"`
	OnsetDate    *time.Time `json:"onset_date,omitempty"`
	ResolvedDate *time.Time `json:"resolved_date,omitempty"`
}

type OutcomeDeltas struct {
	HRVPctChange       *float64 `json:"hrv_pct_change,omitempty"`
	RecoveryPctChange  *float64 `json:"recovery_pct_change,omitempty"`
	RestingHRChange    *float64 `json:"resting_hr_change,omitempty"`
	SleepPctChange     *float64 `json:"sleep_pct_change,omitempty"`
	DeepSleepPctChange *float64 `json:"deep_sleep_pct_change,omitempty"`
	StrainPctChange    *float64 `json:"strain_pct_change,omitempty"`
	VO2MaxPctChange    *float64 `json:"vo2_max_pct_change,omitempty"`
	LeanMassPctChange  *float64 `json:"lean_mass_pct_change,omitempty"`
	FatMassPctChange   *float64 `json:"fat_mass_pct_change,omitempty"`
	BodyFatPctChange   *float64 `json:"body_fat_pct_change,omitempty"`
	BoneDensityChange  *float64 `json:"bone_density_change,omitempty"`
	CaloriePctChange   *float64 `json:"calorie_pct_change,omitempty"`
	ProteinPctChange   *float64 `json:"protein_pct_change,omitempty"`
	SelfRating         *int     `json:"self_rating,omitempty"` // 1-10

	LabDeltas []LabDelta `json:"lab_deltas,omitempty"`
}

type LabDelta struct {
	BiomarkerName string  `json:"biomarker_name"`
	LOINCCode     string  `json:"loinc_code,omitempty"`
	Category      string  `json:"category"`
	BaselineValue float64 `json:"baseline_value"`
	CurrentValue  float64 `json:"current_value"`
	Unit          string  `json:"unit"`
	PctChange     float64 `json:"pct_change"`
}

type OutcomeSnapshot struct {
	Baseline *PeriodMetrics `json:"baseline,omitempty"`
	During   *PeriodMetrics `json:"during,omitempty"`
	Post     *PeriodMetrics `json:"post,omitempty"`
	Deltas   OutcomeDeltas  `json:"deltas"`
}

type PeriodMetrics struct {
	PeriodStart time.Time `json:"period_start"`
	PeriodEnd   time.Time `json:"period_end"`
	HRVAvg     *float64  `json:"hrv_avg,omitempty"`
	RestingHR  *float64  `json:"resting_hr_avg,omitempty"`
	Recovery   *float64  `json:"recovery_avg,omitempty"`
	SleepScore *float64  `json:"sleep_score_avg,omitempty"`
	DeepSleep  *float64  `json:"deep_sleep_avg,omitempty"`
	StrainAvg  *float64  `json:"strain_avg,omitempty"`
	VO2Max     *float64  `json:"vo2_max,omitempty"`
	LeanMass   *float64  `json:"lean_mass_kg,omitempty"`
	FatMass    *float64  `json:"fat_mass_kg,omitempty"`
	BodyFatPct *float64  `json:"body_fat_pct,omitempty"`
	CalorieAvg *float64  `json:"calorie_avg,omitempty"`
	ProteinAvg *float64  `json:"protein_avg,omitempty"`
}

type BodyComposition struct {
	LeanMassKg  float64 `json:"lean_mass_kg"`
	FatMassKg   float64 `json:"fat_mass_kg"`
	BodyFatPct  float64 `json:"body_fat_pct"`
	BoneDensity float64 `json:"bone_density,omitempty"`
}

type NutritionAverage struct {
	CaloriesAvg float64 `json:"calories_avg"`
	ProteinAvg  float64 `json:"protein_avg"`
	CarbsAvg    float64 `json:"carbs_avg"`
	FatAvg      float64 `json:"fat_avg"`
}
