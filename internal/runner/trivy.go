package runner

import "os/exec"

func RunTrivy(path string) ([]byte , error) {
	cmd := exec.Command(
		"trivy",
		"fs",
		path,
		"--format",
		"sarif",
		)
	return cmd.Output()
}
