package anonymize

import (
	"fmt"
	"strings"
)

func ValidateNoLeaks(c *AnonymizedContribution) error {
	if c.Hash == "" {
		return fmt.Errorf("contribution hash is empty")
	}
	if c.SchemaVersion == "" {
		return fmt.Errorf("schema version is empty")
	}
	if c.CompoundID == "" {
		return fmt.Errorf("compound_id is empty")
	}

	if !isValidAgeRange(c.AgeRange) && c.AgeRange != "" {
		return fmt.Errorf("age_range %q is not a valid bucket", c.AgeRange)
	}
	if !isValidBMIRange(c.BMIRange) && c.BMIRange != "" {
		return fmt.Errorf("bmi_range %q is not a valid bucket", c.BMIRange)
	}

	if !isValidDurationBucket(c.DurationBucket) {
		return fmt.Errorf("duration_bucket %q is not a valid bucket", c.DurationBucket)
	}

	if len(c.ContributedAt) != 7 || c.ContributedAt[4] != '-' {
		return fmt.Errorf("contributed_at %q is not month-precision (expected YYYY-MM)", c.ContributedAt)
	}

	for _, se := range c.SideEffects {
		if se.Severity < 1 || se.Severity > 5 {
			return fmt.Errorf("side effect severity %d out of range 1-5", se.Severity)
		}
		if strings.Contains(se.EffectType, " ") {
			return fmt.Errorf("side effect type %q contains free text (must be structured enum)", se.EffectType)
		}
	}

	return nil
}

var validAgeRanges = map[string]bool{
	"18-19": true, "20-29": true, "30-39": true, "40-49": true,
	"50-59": true, "60-69": true, "70-79": true, "80+": true,
}

var validBMIRanges = map[string]bool{
	"<18.5": true, "18.5-20": true, "20-23": true, "23-25": true,
	"25-28": true, "28-30": true, "30-35": true, "35+": true,
}

var validDurationBuckets = map[string]bool{
	"<2 weeks": true, "2-4 weeks": true, "4-8 weeks": true,
	"8-12 weeks": true, "12-24 weeks": true, "24+ weeks": true,
}

func isValidAgeRange(s string) bool  { return validAgeRanges[s] }
func isValidBMIRange(s string) bool  { return validBMIRanges[s] }
func isValidDurationBucket(s string) bool { return validDurationBuckets[s] }
