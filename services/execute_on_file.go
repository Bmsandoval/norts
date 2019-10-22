package services

import (
	"nort/utils"
)

func ExecuteOnFile(executable string, filePath string) error {
	err := utils.Exec(executable + " " + filePath)
	return err
}