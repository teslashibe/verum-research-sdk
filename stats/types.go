package stats

type CohortAnalysis struct {
	CompoundID   string            `json:"compound_id"`
	CompoundName string            `json:"compound_name"`
	DoseBucket   string            `json:"dose_bucket,omitempty"`
	SampleSize   int               `json:"sample_size"`
	Metrics      map[string]Metric `json:"metrics"`
}

type Metric struct {
	Name       string  `json:"name"`
	Mean       float64 `json:"mean"`
	Median     float64 `json:"median"`
	StdDev     float64 `json:"std_dev"`
	N          int     `json:"n"`
	EffectSize float64 `json:"effect_size"`
	CI95Lower  float64 `json:"ci_95_lower"`
	CI95Upper  float64 `json:"ci_95_upper"`
	PValue     float64 `json:"p_value"`
}

type DoseResponseCurve struct {
	CompoundID   string             `json:"compound_id"`
	CompoundName string             `json:"compound_name"`
	MetricName   string             `json:"metric_name"`
	Points       []DoseResponsePoint `json:"points"`
}

type DoseResponsePoint struct {
	DoseBucket string  `json:"dose_bucket"`
	Mean       float64 `json:"mean"`
	N          int     `json:"n"`
	CI95Lower  float64 `json:"ci_95_lower"`
	CI95Upper  float64 `json:"ci_95_upper"`
}

type ConfoundReport struct {
	Warnings      []string            `json:"warnings"`
	Demographics  map[string]int      `json:"demographics"`
	SkewedFields  []string            `json:"skewed_fields"`
	IsReliable    bool                `json:"is_reliable"`
}

type StatisticalResult struct {
	StudyTitle   string          `json:"study_title"`
	Cohort       CohortAnalysis  `json:"cohort"`
	Confounds    ConfoundReport  `json:"confounds"`
	DoseResponse *DoseResponseCurve `json:"dose_response,omitempty"`
}
