package main

import  ("fmt" ; "os" ; "github.com/sandeep7-arch/Hawk/internal/scanner" ; "github.com/sandeep7-arch/Hawk/internal/detector" ; "github.com/sandeep7-arch/Hawk/internal/runner" ; "github.com/sandeep7-arch/Hawk/internal/parser" ; "github.com/sandeep7-arch/Hawk/internal/model")

func main() {
	fmt.Println("Hawk Security Scanner")

	if len(os.Args) < 3 {
		fmt.Println("Usage: hawk scan <path>")
		return
	}

	command := os.Args[1]
	target := os.Args[2]
	if command != "scan" {
		fmt.Println("Command Not Found :" , command)
		return
	}
	fmt.Println("Scanning:" , target)
	Files, err := scanner.ScanDir(target)
	if err!=nil	{
		fmt.Println("Error:" , err)
		return
	}
	Pyfile := detector.DetectPy(target)
	if Pyfile {
		fmt.Println("Python Project Found")
		fmt.Println("Bandit is Being run")
		output , err := runner.RunBandit(target)
		if err!=nil {
			fmt.Println("Bandit Found security Issues")
		}
		report , err2 := parser.Parser(output)
		if err2 != nil {
			fmt.Println("Bandit Parsing Error:" , err2)
		}
		fmt.Println("Bandit Findings:" , len(report.Results))
		scanResult:=model.ScanResult{}
		for _, result := range report.Results {
		finding := parser.ConvBanditResult(result)
		scanResult.Findings = append(scanResult.Findings , finding)
		}
		fmt.Println("Bandit Findings:", len(scanResult.Findings))

		for _, finding := range scanResult.Findings {
			fmt.Println()
			fmt.Println("Scanner:", finding.Scanner)
			fmt.Println("Rule:", finding.RuleID)
			fmt.Println("Severity:", finding.Severity)
			fmt.Println("Confidence:", finding.Confidence)
			fmt.Println("File:", finding.File)
			fmt.Println("Line:", finding.Line)
			fmt.Println("Message:", finding.Message)
		}		
	}	else {
		fmt.Println("NO Python code detected")
	}
	fmt.Println("\nFiles:")

	for _, entry := range Files {
		fmt.Println(" " , entry)
	}

}
