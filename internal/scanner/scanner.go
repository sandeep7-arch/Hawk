package scanner

import "os"

func ScanDir(path string)([]string , error){
	entries , err := os.ReadDir(path)
	if err != nil {
		return nil , err
	}
	var files []string
	for _, val := range entries {
		files = append(files , val.Name())
	}
	return files , nil
}
