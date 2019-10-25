package services

import (
	"log"
	"norts/utils"
)

func ExecuteOnFile(executable string, filePath string) error {
	log.Println(filePath)
	err := utils.ExecNotCapturingOutput(executable, []string{filePath})
	return err
}
