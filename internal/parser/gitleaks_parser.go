package parser

import ("encoding/json" ; "github.com/sandeep7-arch/Hawk/internal/model" ; "time")

type GLReport struct {
	Results []GLResult 
}

type GLResult struct {
	Description string    `json:"Description"`
	StartLine   int       `json:"StartLine"`
	EndLine     int       `json:"EndLine"`
	StartColumn int       `json:"StartColumn"`
	EndColumn   int       `json:"EndColumn"`
	Match       string    `json:"Match"`
	Secret      string    `json:"Secret"`
	File        string    `json:"File"`
	SymlinkFile string    `json:"SymlinkFile"`
	Commit      string    `json:"Commit"`
	Author      string    `json:"Author"`
	Email       string    `json:"Email"`
	Date        time.Time `json:"Date"`
	Message     string    `json:"Message"`
	Tags        []string  `json:"Tags"`
	RuleID      string    `json:"RuleID"`
	Fingerprint string    `json:"Fingerprint"`
}

func GLParser(data []byte)(GLReport , error) {
	var report GLReport
	err := json.Unmarshal(data , &report)

	if err!=nil {
		return GLReport{} , err
	}
	return report , nil
}

func ConvGLResult(result GLResult) (model.Finding) {
	finding := model.Finding{
		Scanner: "GitLeaks",
		RuleID: result.RuleID,
		RuleName: result.Description,
		Message: result.Message,
		File: result.File,
		Line: result.StartLine,
	}	
	return finding
}
