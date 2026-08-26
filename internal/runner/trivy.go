package runner

import "os/exec"

func RunTrivyDependency(path string) ([]byte , error) {
	cmd := exec.Command(
		"trivy",
		"fs",
		path,
		"--format",
		"sarif",
		)
	return cmd.Output()
}

func RunTrivyDocker(path String) ([]byte , error) {
	cmd := exec.Command(
		"trivy",
		
		)
}
