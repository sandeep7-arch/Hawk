package runner

import ("os/exec")

func RunGitLeaks(path string , arg1 string) ([]byte , error) {
	cmd := exec.Command(
		"gitleaks",
		arg1,
		"-f",
		"json",
		path,
		)
	return cmd.Output()
}
