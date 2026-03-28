package compounds

type Category string

const (
	CategoryHealingRecovery  Category = "healing_recovery"
	CategoryGHSecretagogue   Category = "gh_secretagogue"
	CategoryBodyComposition  Category = "body_composition"
	CategoryCognitive        Category = "cognitive"
	CategoryAntiAging        Category = "anti_aging"
	CategoryImmune           Category = "immune"
	CategorySexualHealth     Category = "sexual_health"
	CategoryMuscle           Category = "muscle_performance"
	CategoryGutHealth        Category = "gut_health"
	CategoryPainInflammation Category = "pain_inflammation"
	CategorySkinCosmetic     Category = "skin_cosmetic"
	CategoryEmerging         Category = "emerging_research"
)

type Status string

const (
	StatusVerified          Status = "verified"
	StatusCommunityProposed Status = "community_proposed"
	StatusUnderReview       Status = "under_review"
)

type Compound struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Aliases           []string `json:"aliases,omitempty"`
	Category          Category `json:"category"`
	Subcategory       string   `json:"subcategory,omitempty"`
	Description       string   `json:"description,omitempty"`
	Mechanism         string   `json:"mechanism,omitempty"`
	TypicalDoseMin    float64  `json:"typical_dose_min,omitempty"`
	TypicalDoseMax    float64  `json:"typical_dose_max,omitempty"`
	DoseUnit          string   `json:"dose_unit,omitempty"`
	CommonRoutes      []string `json:"common_routes,omitempty"`
	HalfLife          string   `json:"half_life,omitempty"`
	TypicalCycle      string   `json:"typical_cycle,omitempty"`
	Contraindications []string `json:"contraindications,omitempty"`
	PubMedRefs        []string `json:"pubmed_refs,omitempty"`
	Status            Status   `json:"status"`
}
