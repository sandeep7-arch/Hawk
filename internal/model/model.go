package model

type Finding struct {
	Scanner    string
	RuleID     string
	RuleName   string
	Severity   string
	Confidence string
	Message    string
	File       string
	Line       int
}
