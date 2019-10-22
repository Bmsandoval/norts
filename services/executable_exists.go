package services

import (
	"log"
	"os/exec"
)

func ExecutableExists(executable string) bool {
	what, err := exec.LookPath(executable)
	if err != nil {
		log.Println(what)
		return false
	}
	log.Println(what)
	return true
}
