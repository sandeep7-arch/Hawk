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



	scanResult:=model.ScanResult{}
	Pyfile := detector.DetectPy(target)
	if Pyfile {
		fmt.Println("Python Project Found")
		fmt.Println("Bandit is Being run")
		output_bandit , err := runner.RunBandit(target)
		if err!=nil {
			fmt.Println("Bandit Found security Issues")
		}
		bandit_report , err2 := parser.Parser(output_bandit)
		if err2 != nil {
			fmt.Println("Bandit Parsing Error:" , err2)
		}

		for _, result := range bandit_report.Results {
		finding := parser.ConvBanditResult(result)
		scanResult.Findings = append(scanResult.Findings , finding)
		}
			}	else {
		fmt.Println("NO Python code detected")
	}




	fmt.Println()
	fmt.Println("Trivy is being run")

	output_trivy , err3 := runner.RunTrivy(target)

	if err3!= nil {
		fmt.Println("Trivy Found Security Issues")
	}
	
	trivy_report , err4:= parser.TrivyParser(output_trivy)
	if err4 != nil {
		fmt.Println("Trivy Parsing Error:" , err4)
	}

	for _, run := range trivy_report.Runs {

    for _, result := range run.Results {

        for _, rule := range run.Tool.Driver.Rules {

            if rule.ID == result.RuleID {

                finding := parser.ConvTrivyResult(result, rule)

                scanResult.Findings = append(
                    scanResult.Findings,
                    finding,
                )
                break
            }
        }
    }
}
	fmt.Println("Findings Of Bandit and Trivy") 
	for _, finding := range scanResult.Findings {
				fmt.Println()
				fmt.Println("Scanner:", finding.Scanner)
				fmt.Println("Rule:", finding.RuleID)
				fmt.Println("Severity:", finding.Severity)
				fmt.Println("Confidence:", finding.Confidence)
				fmt.Println("File:", finding.File)
				fmt.Println("Line:", finding.Line)
				fmt.Println("Message:", finding.Message)
				fmt.Println()
			}		

	fmt.Println("\nFiles:")

	for _, entry := range Files {
		fmt.Println(" " , entry)
	}

}
