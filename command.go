package dot


func bash(script string) (string, error) {
	return exec.Command("/bin/bash", script).Output()
}
