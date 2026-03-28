package biomarkers

type Category string

const (
	CategoryIron         Category = "iron"
	CategoryLipid        Category = "lipid"
	CategoryMetabolic    Category = "metabolic"
	CategoryLiver        Category = "liver"
	CategoryKidney       Category = "kidney"
	CategoryElectrolyte  Category = "electrolyte"
	CategoryHormonal     Category = "hormonal"
	CategoryThyroid      Category = "thyroid"
	CategoryInflammatory Category = "inflammatory"
	CategoryCBC          Category = "cbc"
	CategoryVitamin      Category = "vitamin"
	CategoryMineral      Category = "mineral"
	CategoryCardiac      Category = "cardiac"
	CategoryProtein      Category = "protein"
)

type RefRangeType string

const (
	RefRangeMinMax      RefRangeType = "min_max"       // 50-180
	RefRangeLessThan    RefRangeType = "less_than"      // <200
	RefRangeGreaterThan RefRangeType = "greater_than"   // >= 40
	RefRangeLessEqual   RefRangeType = "less_equal"     // <= 13.5
	RefRangeTiered      RefRangeType = "tiered"         // optimal/moderate/high
)

type Flag string

const (
	FlagNormal   Flag = "normal"
	FlagLow      Flag = "low"
	FlagHigh     Flag = "high"
	FlagCritical Flag = "critical"
)

type Biomarker struct {
	Name           string       `json:"name"`
	LOINCCode      string       `json:"loinc_code,omitempty"`
	Category       Category     `json:"category"`
	Unit           string       `json:"unit"`
	RefRangeType   RefRangeType `json:"ref_range_type"`
	RefRangeLow    *float64     `json:"ref_range_low,omitempty"`
	RefRangeHigh   *float64     `json:"ref_range_high,omitempty"`
	OptimalLow     *float64     `json:"optimal_low,omitempty"`
	OptimalHigh    *float64     `json:"optimal_high,omitempty"`
	PeptideRelevance string     `json:"peptide_relevance,omitempty"`
	Aliases        []string     `json:"aliases,omitempty"`
}

type LabResult struct {
	BiomarkerName string   `json:"biomarker_name"`
	LOINCCode     string   `json:"loinc_code,omitempty"`
	Category      Category `json:"category"`
	Value         float64  `json:"value"`
	Unit          string   `json:"unit"`
	RefRangeLow   *float64 `json:"ref_range_low,omitempty"`
	RefRangeHigh  *float64 `json:"ref_range_high,omitempty"`
	Flag          Flag     `json:"flag"`
	PreviousValue *float64 `json:"previous_value,omitempty"`
	PreviousDate  string   `json:"previous_date,omitempty"`
}

type LabReport struct {
	LabProvider    string      `json:"lab_provider"`
	CollectionDate string     `json:"collection_date"`
	ReportDate     string     `json:"report_date,omitempty"`
	PanelName      string     `json:"panel_name,omitempty"`
	Fasting        bool       `json:"fasting"`
	Results        []LabResult `json:"results"`
}

func (b *Biomarker) Evaluate(value float64) Flag {
	switch b.RefRangeType {
	case RefRangeMinMax:
		if b.RefRangeLow != nil && value < *b.RefRangeLow {
			return FlagLow
		}
		if b.RefRangeHigh != nil && value > *b.RefRangeHigh {
			return FlagHigh
		}
	case RefRangeLessThan, RefRangeLessEqual:
		if b.RefRangeHigh != nil && value > *b.RefRangeHigh {
			return FlagHigh
		}
	case RefRangeGreaterThan:
		if b.RefRangeLow != nil && value < *b.RefRangeLow {
			return FlagLow
		}
	case RefRangeTiered:
		if b.OptimalHigh != nil && value > *b.OptimalHigh {
			return FlagHigh
		}
		if b.OptimalLow != nil && value < *b.OptimalLow {
			return FlagLow
		}
	}
	return FlagNormal
}
