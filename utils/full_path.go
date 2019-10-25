package utils

import (
	"github.com/mitchellh/go-homedir"
	"norts/config"
	"strings"
)

func FullPathToFile(relativePath string) (string, error) {
	configs := config.GetConfigFromViper()
	homeDirectory, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	remainingPath := configs.Norts.Path
	// compatibility layer. Don't force users to include trailing slash
	trimmedPath := strings.TrimSuffix(remainingPath, "/")

	fileFullPath := trimmedPath + "/" + strings.TrimPrefix(relativePath, "./")

	fileFullyQualifiedPath := homeDirectory + "/" + strings.TrimPrefix(fileFullPath, "~/")

	return fileFullyQualifiedPath, nil
}
