package install

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func getDistribution() string {
	cmd := exec.Command("uname")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	return strings.Trim(out.String(), "\n")
}

// Packages install on OS
func Packages(packages string) {
	distrib := getDistribution()

	switch distrib {
	case "Darwin":
		println("Try on unix")
	case "Ubuntu":
		println("installDebian()")
	default:
		println("Incompatible distribution")
	}
}
