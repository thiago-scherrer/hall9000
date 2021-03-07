package volume

import (
	"log"
	"os"
	"os/exec"
)

func Start(v string) {
	cmd := exec.Command("amixer", "sset", "Master "+v+"%")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
