package anonymize

type Config struct {
	MinK    int     // minimum k-anonymity threshold (default: 5)
	Epsilon float64 // differential privacy epsilon (default: 1.0)
}

func DefaultConfig() Config {
	return Config{
		MinK:    5,
		Epsilon: 1.0,
	}
}

type Demographics struct {
	AgeRange string `json:"age_range"` // "20-29", "30-39", "40-49", etc.
	Sex      string `json:"sex"`       // "male", "female", "other"
	BMIRange string `json:"bmi_range"` // "18-20", "20-23", "23-25", "25-28", "28-30", "30+"
}

type AnonymizedContribution struct {
	CompoundID     string            `json:"compound_id"`
	CompoundName   string            `json:"compound_name"`
	DoseBucket     string            `json:"dose_bucket"`
	Frequency      string            `json:"frequency"`
	Route          string            `json:"route"`
	DurationBucket string            `json:"duration_bucket"`
	GoalCategory   string            `json:"goal_category,omitempty"`
	AgeRange       string            `json:"age_range,omitempty"`
	Sex            string            `json:"sex,omitempty"`
	BMIRange       string            `json:"bmi_range,omitempty"`
	OutcomeDeltas  map[string]float64 `json:"outcome_deltas"`
	SideEffects    []AnonymizedSideEffect `json:"side_effects,omitempty"`
	SelfRating     *int              `json:"self_rating,omitempty"`
	Hash           string            `json:"contribution_hash"`
	SchemaVersion  string            `json:"schema_version"`
	ContributedAt  string            `json:"contributed_at"` // month precision: "2026-03"
}

type AnonymizedSideEffect struct {
	EffectType string `json:"effect_type"`
	Severity   int    `json:"severity"` // 1-5
}
