package contribute

import (
	"github.com/verum-research/sdk/anonymize"
	"github.com/verum-research/sdk/stats"
)

type SubmitRequest struct {
	Contribution anonymize.AnonymizedContribution `json:"contribution"`
}

type SubmitResponse struct {
	ID           string `json:"id"`
	CompoundName string `json:"compound_name"`
	Message      string `json:"message"`
}

type PlatformStats struct {
	TotalContributors int            `json:"total_contributors"`
	TotalContributions int           `json:"total_contributions"`
	CompoundsStudied  int            `json:"compounds_studied"`
	ReportsPublished  int            `json:"reports_published"`
	TopCompounds      []CompoundStat `json:"top_compounds"`
}

type CompoundStat struct {
	CompoundID   string `json:"compound_id"`
	CompoundName string `json:"compound_name"`
	Count        int    `json:"count"`
}

type ReportSummary struct {
	ID       string `json:"id"`
	Slug     string `json:"slug"`
	Title    string `json:"title"`
	Abstract string `json:"abstract"`
	Status   string `json:"status"`
}

type Report struct {
	ReportSummary
	ContentMD         string                `json:"content_md"`
	StatisticalResults stats.StatisticalResult `json:"statistical_results"`
	Disclaimers       string                `json:"disclaimers"`
}

type Finding struct {
	CompoundName string       `json:"compound_name"`
	MetricName   string       `json:"metric_name"`
	EffectSize   float64      `json:"effect_size"`
	PValue       float64      `json:"p_value"`
	SampleSize   int          `json:"sample_size"`
	Summary      string       `json:"summary"`
	ReportSlug   string       `json:"report_slug"`
}
