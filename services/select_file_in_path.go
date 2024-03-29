package services

import (
	"errors"
	"github.com/ktr0731/go-fuzzyfinder"
	"io/ioutil"
	"norts/bash"
	"norts/utils"
)

func SelectFileInPath(path string) (string, error) {
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
			if i == -1 { return "" }

			// Given a relative path, get full path
			fileFullPath, err := utils.FullPathToFile(out[i])
			if err != nil { return "" }

			// Read contents of file
			data, err := ioutil.ReadFile(fileFullPath)
			if err != nil { return "" }

			return string(data)
		}));
	err != nil {
		// nothing selected, nothing to return
			return "", nil
	} else {
		// Given a relative path, get full path
		fileFullPath, err := utils.FullPathToFile(out[selected])
		if err != nil { return "", err }
		// return selected option
		return fileFullPath, nil
	}
}
