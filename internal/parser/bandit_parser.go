package parser

import "encoding/json"

type BanditReport struct {
	Results []BanditResult `json:"Results"`
}

type BanditResult struct {
	TestID     string `json:"test_id"`
	TestName   string `json:"test_name"`
	Severity   string `json:"issue_severity"`
	Confidence string `json:"issue_confidence"`
	IssueText  string `json:"issue_text"`
	Filename   string `json:"filename"`
	LineNumber int    `json:"line_number"`
	LineRange  []int  `json:"line_range"`
}

func Parser(data []byte)(BanditReport , error) {
	var report BanditReport
	err:= json.Unmarshal(data , &report)
	if err!= nil {
		return BanditReport{} , err
	}
	return report , nil
}
