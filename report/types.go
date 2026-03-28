package report

type Study struct {
	ID              string   `json:"id"`
	Title           string   `json:"title"`
	StudyType       string   `json:"study_type"`
	Hypothesis      string   `json:"hypothesis,omitempty"`
	CompoundIDs     []string `json:"compound_ids"`
	Status          string   `json:"status"`
	MinSampleSize   int      `json:"min_sample_size"`
	CurrentSampleSize int   `json:"current_sample_size"`
}

type Report struct {
	ID        string    `json:"id"`
	StudyID   string    `json:"study_id"`
	Version   int       `json:"version"`
	DOISlug   string    `json:"doi_slug"`
	Title     string    `json:"title"`
	Abstract  string    `json:"abstract"`
	Sections  []Section `json:"sections"`
	Disclaimer string  `json:"disclaimer"`
}

type Section struct {
	Heading string `json:"heading"`
	Content string `json:"content"`
}
