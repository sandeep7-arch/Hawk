package policy

import ("github.com/sandeep7-arch/Hawk/internal/model")

func ShouldFail(result model.ScanResult) bool {
	for _, finding := range result.Findings {
		if finding.Severity == "HIGH" || finding.Severity == "CRITICAL" {
			return true
		}

	}
	return false
}
