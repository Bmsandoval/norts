package bash

import "fmt"

func ListFilesInPath(path string) string { return fmt.Sprintf("cd %s && find . -type f", path) }
func ListFoldersInPath(path string) string { return fmt.Sprintf("cd %s && find . -type d", path) }
func ListFilesAndFoldersInPath(path string) string { return fmt.Sprintf("cd %s && find .", path) }

