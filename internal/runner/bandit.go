package runner

import "os/exec"

func RunBandit(path string)([]byte , error) {
	cmd := exec.Command(
		"Bandit",
		"-r",
		path,
		"-f",
		"json",
		)

	return cmd.Output()
}
