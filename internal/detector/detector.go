package detector

import ("os" ; "path/filepath")

func DetectPy(path string) (bool) {
	found := false
	err := filepath.Walk(path , func(currPath string , info os.FileInfo , err error)(error){
		if err!=nil{
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(currPath)==".py"{
			found = true;

		}
		return nil

	})
	if err!=nil{
	return false
	}
	return found
}
