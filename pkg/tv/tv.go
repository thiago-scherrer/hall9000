package tv

import (
	"log"
	"os"
	"os/exec"

	"github.com/thiago-scherrer/hall9000/internal/config"
)

func Jornal() {
	var cmd *exec.Cmd
	tv := config.GetTvIp()

	cmd = exec.Command("amsungctl", "--host", tv, "--id", "42", "KEY_5")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
