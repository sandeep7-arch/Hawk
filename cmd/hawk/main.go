package main

import  ("fmt" ; "os")

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
	entries , err := os.ReadDir(target)
	if err!=nil {
		fmt.Println("Error:" , err)
		return
	}
	fmt.Println("\nFiles:")

	for _, entry := range entries {
		fmt.Println(" " , entry.Name())
	}

}
