package services

import (
	"errors"
	"github.com/ktr0731/go-fuzzyfinder"
	"norts/bash"
	"norts/utils"
	"path/filepath"
	"strings"
)

func SelectFolderInPath(path string) (string, error) {
	err, out, errout := utils.ExecGetOutputArray(bash.ListFoldersInPath(path))
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

			// Given a relative path, get full path
			fileFullPath, err := utils.FullPathToFile(out[i])
			if err != nil { return "" }

			// Get contents of folder
			folderContents, err := filepath.Glob(fileFullPath + "/*")
			if err != nil { return "" }

			for key, folderPath := range folderContents {
				folderContents[key] = strings.TrimPrefix(folderPath, fileFullPath)
			}

			return strings.Join(folderContents[:], "\n")
		}));
	err != nil {
		// nothing selected, nothing to return
			return "", nil
	} else {
		// Given a relative path, get full path
		folderFullPath, err := utils.FullPathToFile(out[selected])
		if err != nil { return "", err }
		// return selected option
		return folderFullPath, nil
	}
}
