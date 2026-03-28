package anonymize

import "fmt"

type BucketCounts interface {
	CountByBucket(compoundID, doseBucket, durationBucket, ageRange, sex string) (int, error)
}

func CheckKAnonymity(c *AnonymizedContribution, counts BucketCounts, k int) (bool, error) {
	count, err := counts.CountByBucket(
		c.CompoundID,
		c.DoseBucket,
		c.DurationBucket,
		c.AgeRange,
		c.Sex,
	)
	if err != nil {
		return false, fmt.Errorf("checking k-anonymity: %w", err)
	}

	// count + 1 because this contribution would be added
	return count+1 >= k, nil
}
