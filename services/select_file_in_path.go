package services

import (
	"errors"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"nort/bash"
	"nort/config"
	"nort/utils"
	"strings"
)

func SelectFileInPath(path string) (string, error) {
	configs := config.GetConfigFromViper()
	_ = configs
	homeDirectory, err := homedir.Dir()

	err, out, errout := utils.ExecGetOutputArray(bash.ListFilesInPath(path))
	if err != nil {
		return "", err
	}
	if errout != "" {
		return "", errors.New(errout)
	}

	if len(out) < 1 {
		// directory empty
		return "", nil
	}

	if selected, err := fuzzyfinder.Find(out,
		func(i int) string {
			return out[i]
		}, fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			relativeFilePath := configs.Nort.Path + strings.TrimPrefix(out[i], "./")
			fileFullPath := homeDirectory + "/" + strings.TrimPrefix(relativeFilePath, "~/")
			data, err := ioutil.ReadFile(fileFullPath)
			if err != nil {
				return ""
			}
			return string(data)
		}));
	err != nil {
		// nothing selected, nothing to return
			return "", nil
	} else {
		relativePath := out[selected]
		fullPath := configs.Nort.Path + strings.TrimPrefix(relativePath, "./")
		// return selected option
		return fullPath, nil
	}
}
