package runner

import "os/exec"

func RunBandit(path string)([]byte , error) {
	cmd := exec.Command(
		"bandit",
		"-r",
		path,
		"-f",
		"json",
		)

	return cmd.Output()
}
