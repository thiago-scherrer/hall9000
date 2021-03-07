package volume

import (
	"log"
	"os"
	"os/exec"
)

func Start(v string) {
	var cmd *exec.Cmd

	if v == "+" {
		cmd = exec.Command("amixer", "-q", "-D", "pulse", "sset", "Master", "25%+")
	} else {
		cmd = exec.Command("amixer", "-q", "-D", "pulse", "sset", "Master", "25%-")

	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
