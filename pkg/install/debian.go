package install

import (
	"os"
	"os/exec"
	"strings"
)

func installDebian(packages []string) error {
	cmd := exec.Command("apt install " + strings.Join(packages, " "))
	cmd.Env = append(os.Environ())
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
