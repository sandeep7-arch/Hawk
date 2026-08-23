package parser

import (
	"encoding/json";
	"github.com/sandeep7-arch/Hawk/internal/model"
)

type TrivyReport struct {
	Runs []TrivyRun `json:"runs"`
}

type TrivyRun struct {
	Tool    TrivyTool     `json:"tool"`
	Results []TrivyResult `json:"results"`
}

type TrivyTool struct {
	Driver TrivyDriver `json:"driver"`
}

type TrivyDriver struct {
	Rules []TrivyRule `json:"rules"`
}

type TrivyRule struct {
	ID                 string          `json:"id"`
	Name               string          `json:"name"`
	ShortDescription   TrivyText       `json:"shortDescription"`
	DefaultConfiguration TrivyConfig   `json:"defaultConfiguration"`
	Properties         TrivyProperties `json:"properties"`
}

type TrivyText struct {
	Text string `json:"text"`
}

type TrivyConfig struct {
	Level string `json:"level"`
}

type TrivyProperties struct {
	SecuritySeverity string  `json:"security-severity"`
	CVSSBaseScore    float64 `json:"cvssv3_baseScore"`
}

type TrivyResult struct {
	RuleID   string         `json:"ruleId"`
	Level    string         `json:"level"`
	Message  TrivyText      `json:"message"`
	Locations []TrivyLocation `json:"locations"`
}

type TrivyLocation struct {
	PhysicalLocation TrivyPhysicalLocation `json:"physicalLocation"`
}

type TrivyPhysicalLocation struct {
	ArtifactLocation TrivyArtifactLocation `json:"artifactLocation"`
	Region           TrivyRegion           `json:"region"`
}

type TrivyArtifactLocation struct {
	URI string `json:"uri"`
}

type TrivyRegion struct {
	StartLine int `json:"startLine"`
}

func TrivyParser(data []byte) (TrivyReport , error) {
	var report TrivyReport
	err:=json.Unmarshal(data , &report)
	if err!= nil {
		return TrivyReport{} , err
	}
	return report , nil
}

func ConvTrivyResult(result TrivyResult , rule TrivyRule) (model.Finding) {
	finding:= model.Finding{
		Scanner: "Trivy",
		RuleID: result.RuleID,
		RuleName: rule.Name,
		Message: result.Message.Text,
	}
		if rule.Properties.CVSSBaseScore <= 3.9 {
		finding.Severity = "LOW"		
	} 	else if rule.Properties.CVSSBaseScore <=6.9 {
		finding.Severity = "MEDIUM"

	} 	else if rule.Properties.CVSSBaseScore <= 8.9 {
		finding.Severity = "HIGH"
	}   else {
		finding.Severity = "CRITICAL"
	}
		if len(result.Locations) > 0 {
			finding.File = result.Locations[0].PhysicalLocation.ArtifactLocation.URI
			finding.Line = result.Locations[0].PhysicalLocation.Region.StartLine
		}
	return finding
	}


