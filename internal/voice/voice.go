package voice

import (
	"log"
	"os"
	"os/exec"
	"regexp"
)

func Start(s string) {
	d := clear(s)
	cmda := exec.Command("gtts-cli", "-l", "pt", `--nocheck`, `-o`, "/tmp/.hall9000.mp3", `"`+d+`"`)
	cmda.Stdout = os.Stdout
	cmda.Stderr = os.Stderr
	err := cmda.Run()
	if err != nil {
		os.Setenv("CONTROL", "true")
		log.Println(err)
	}

	cmdb := exec.Command("play", "-t", "mp3", "/tmp/.hall9000.mp3")
	cmdb.Stdout = os.Stdout
	cmdb.Stderr = os.Stderr
	err = cmdb.Run()
	if err != nil {
		os.Setenv("CONTROL", "true")
		log.Println(err)
	}

}

func clear(s string) string {
	myRegex, _ := regexp.Compile(`(")`)
	return myRegex.ReplaceAllString(s, "")
}
