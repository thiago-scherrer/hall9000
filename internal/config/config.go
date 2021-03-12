package config

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func GetKey() string {
	tgkey := os.Getenv("TGKEY")
	if len(tgkey) <= 0 {
		log.Fatal("Telegram key empty")
	}

	return tgkey
}

func GetDebug() bool {
	debug := os.Getenv("DEBUG")
	return debug == "true"
}

func GetTimeOut() int {
	timeout := os.Getenv("TIMEOUT")
	if len(timeout) <= 0 {
		return 60
	}

	i, _ := strconv.Atoi(timeout)
	return i
}

func GetClimaKey() string {
	k := os.Getenv("CLIMAKEY")
	return k
}

func GetTvIp() string {
	k := os.Getenv("TVIP")
	return k
}

func GetControl() bool {
	control := os.Getenv("CONTROL")

	if control == "true" {
		os.Setenv("CONTROL", "false")

		out, err := exec.Command("killall", "-9", "gtts-cli").Output()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(out)

		out2, err := exec.Command("killall", "-9", "play").Output()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(out2)
		return true
	}

	return false
}
