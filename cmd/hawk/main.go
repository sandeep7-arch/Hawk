package main

import  ("fmt" ; "os" ; "github.com/sandeep7-arch/Hawk/internal/scanner" ; "github.com/sandeep7-arch/Hawk/internal/detector" ; "github.com/sandeep7-arch/Hawk/internal/runner" ; "github.com/sandeep7-arch/Hawk/internal/parser")

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
			fmt.Println("Bandit Error:", err)
			return
		}
		report , err2 := parser.Parser(output)
		if err2 != nil {
			fmt.Println("Bandit Parsing Error:" , err2)
		}
		fmt.Println("Bandit Findings:" , len(report.Results))
		for _, result := range report.Results {
			fmt.Println()
			fmt.Println("Rule:", result.TestID)
			fmt.Println("Severity:", result.Severity)
			fmt.Println("Confidence:", result.Confidence)
			fmt.Println("File:", result.Filename)
			fmt.Println("Line:", result.LineNumber)
			fmt.Println("Message:", result.IssueText)
		}
	}	else {
		fmt.Println("NO Python code detected")
	}
	fmt.Println("\nFiles:")

	for _, entry := range Files {
		fmt.Println(" " , entry)
	}

}
