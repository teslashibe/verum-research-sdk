package anonymize

import (
	"fmt"
	"time"
)

func BucketDose(dose float64, unit string) string {
	switch unit {
	case "mg":
		return bucketMg(dose)
	case "mcg":
		return bucketMcg(dose)
	default:
		return fmt.Sprintf("%.0f %s", dose, unit)
	}
}

func bucketMcg(dose float64) string {
	switch {
	case dose < 100:
		return "<100mcg"
	case dose < 200:
		return "100-200mcg"
	case dose < 300:
		return "200-300mcg"
	case dose < 500:
		return "300-500mcg"
	case dose < 750:
		return "500-750mcg"
	case dose < 1000:
		return "750-1000mcg"
	default:
		return ">1000mcg"
	}
}

func bucketMg(dose float64) string {
	switch {
	case dose < 0.5:
		return "<0.5mg"
	case dose < 1:
		return "0.5-1mg"
	case dose < 2:
		return "1-2mg"
	case dose < 5:
		return "2-5mg"
	case dose < 10:
		return "5-10mg"
	case dose < 25:
		return "10-25mg"
	default:
		return ">25mg"
	}
}

func BucketDuration(days int) string {
	switch {
	case days < 14:
		return "<2 weeks"
	case days < 28:
		return "2-4 weeks"
	case days < 56:
		return "4-8 weeks"
	case days < 84:
		return "8-12 weeks"
	case days < 168:
		return "12-24 weeks"
	default:
		return "24+ weeks"
	}
}

func MonthPrecision() string {
	return time.Now().Format("2006-01")
}
