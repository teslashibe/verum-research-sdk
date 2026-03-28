package anonymize

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func ContributionHash(c *AnonymizedContribution) string {
	salt := make([]byte, 16)
	_, _ = rand.Read(salt)

	data := fmt.Sprintf("%s|%s|%s|%s|%s|%s|%x",
		c.CompoundID,
		c.DoseBucket,
		c.DurationBucket,
		c.Route,
		c.Frequency,
		c.ContributedAt,
		salt,
	)

	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
