package main

import  ("fmt" ; "os" ; "github.com/sandeep7-arch/Hawk/internal/scanner" ; "github.com/sandeep7-arch/Hawk/internal/detector")

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
	Files , err := scanner.ScanDir(target)
	if err!=nil	{
		fmt.Println("Error:" , err)
		return
	}
	fmt.Println("\nFiles:")

	for _, entry := range Files {
		fmt.Println(" " , entry)
	}

}
