package report

import (
	"fmt"
	"strings"

	"github.com/teslashibe/verum-research-sdk/stats"
)

func NewReportTemplate(study Study, result stats.StatisticalResult) Report {
	return Report{
		StudyID:    study.ID,
		Version:    1,
		Title:      result.StudyTitle,
		Sections:   buildSections(study, result),
		Disclaimer: StandardDisclaimer,
	}
}

func buildSections(study Study, result stats.StatisticalResult) []Section {
	return []Section{
		{Heading: "Abstract", Content: ""},
		{Heading: "Methodology", Content: buildMethodology(study, result)},
		{Heading: "Results", Content: buildResults(result)},
		{Heading: "Discussion", Content: ""},
		{Heading: "Limitations", Content: buildLimitations(result)},
		{Heading: "Data Availability", Content: DataLicense},
		{Heading: "Disclaimer", Content: StandardDisclaimer},
	}
}

func buildMethodology(study Study, result stats.StatisticalResult) string {
	var b strings.Builder
	fmt.Fprintf(&b, "**Study type:** Retrospective observational cohort analysis\n\n")
	fmt.Fprintf(&b, "**Data source:** Verum Research anonymous contribution pool\n\n")
	fmt.Fprintf(&b, "**Sample size:** n = %d\n\n", result.Cohort.SampleSize)
	fmt.Fprintf(&b, "**Statistical methods:** One-sample t-test for significance, Cohen's d for effect size, 95%% confidence intervals\n\n")
	fmt.Fprintf(&b, "**Confound controls:** Demographic stratification (age range, sex, BMI range)\n")
	return b.String()
}

func buildResults(result stats.StatisticalResult) string {
	var b strings.Builder
	for name, m := range result.Cohort.Metrics {
		significance := "not statistically significant"
		if stats.IsSignificant(m.PValue, 0.05) {
			significance = "statistically significant"
		}
		fmt.Fprintf(&b, "**%s:** Mean change %.1f%% (95%% CI: %.1f–%.1f), Cohen's d = %.2f (%s), p = %.4f, n = %d — %s\n\n",
			name, m.Mean, m.CI95Lower, m.CI95Upper, m.EffectSize,
			stats.InterpretCohenD(m.EffectSize), m.PValue, m.N, significance)
	}
	return b.String()
}

func buildLimitations(result stats.StatisticalResult) string {
	var b strings.Builder
	b.WriteString("- Observational design — no causal claims\n")
	b.WriteString("- Self-selected population (health-tracking users)\n")
	b.WriteString("- Self-reported protocol adherence\n")
	b.WriteString("- No placebo control group\n")
	b.WriteString("- Wearable data accuracy limitations\n")

	for _, w := range result.Confounds.Warnings {
		fmt.Fprintf(&b, "- %s\n", w)
	}
	return b.String()
}
