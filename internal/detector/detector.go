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

var DependencyFiles = []string{
    "requirements.txt",
    "Pipfile.lock",
    "poetry.lock",
    "uv.lock",
    "package.json",
    "package-lock.json",
    "yarn.lock",
    "pnpm-lock.yaml",
    "go.mod",
    "go.sum",
    "Cargo.lock",
    "Gemfile.lock",
    "composer.lock",
    "pom.xml",
    "gradle.lockfile",
    "mix.lock",
    "pubspec.lock",
    "Podfile.lock",
    "Package.resolved",
    "Manifest.toml",
    "conan.lock",
}

var dockerFiles = []string{
    "Dockerfile",
    "docker-compose.yml",
    "docker-compose.yaml",
    "compose.yml",
    "compose.yaml",
}

func DetectTrivyDependencyScanner(path string) (bool) {
	found := false
	err := filepath.Walk(path , func(currPath string , info os.FileInfo , err error)(error){
	if err != nil {
		return err
	}
	if info.IsDir() {
	return nil
	}
	for _, file := range DependencyFiles {
			if info.Name() == file {
				found = true
				return nil
			}
		}
		return nil
		})
	if err!= nil{
		return false;
	}
	return found;
}

func DetectTrivyDockerScanner(path string) (bool) {
	found := false
		err := filepath.Walk(path , func(currPath string , info os.FileInfo , err error)(error){
		if err != nil {
			return err
		}
		if info.IsDir() {
		return nil
		}
		for _, file := range dockerFiles {
				if info.Name() == file {
					found = true
					return nil
				}
			}
			return nil
			})
		if err!= nil{
			return false;
		}
		return found;
}
