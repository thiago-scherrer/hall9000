package tv

import (
	"log"
	"os"
	"os/exec"

	"github.com/thiago-scherrer/hall9000/internal/config"
)

func Canal(s string) {
	closeAll()
	canalgo(s)
}

func Tvi() {
	var cmd *exec.Cmd
	tv := config.GetTvIp()

	i := 1
	for i <= 6 {
		cmd = exec.Command("samsungctl", "--host", tv, "--id", "42", "KEY_VOLUP")

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			log.Println(err)
		}

		i += 1
	}
}

func Tvd() {
	var cmd *exec.Cmd
	tv := config.GetTvIp()

	i := 1
	for i <= 6 {
		cmd = exec.Command("samsungctl", "--host", tv, "--id", "42", "KEY_VOLDOWN")

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			log.Println(err)
		}

		i += 1
	}
}

func closeAll() {
	var cmd *exec.Cmd
	tv := config.GetTvIp()

	cmd = exec.Command("samsungctl", "--host", tv, "--id", "42", "KEY_DTV")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}

func canalgo(s string) {
	var cmd *exec.Cmd
	ch := "KEY_" + s

	tv := config.GetTvIp()

	closeAll()

	cmd = exec.Command("samsungctl", "--host", tv, "--id", "42", ch)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
