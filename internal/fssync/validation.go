package fssync

import (
	"os"
)

func FolderExist(path string) bool {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateFolderIfNotExist(path string) {
	if !FolderExist(path) {
		os.MkdirAll(path, 0755)
	}
}
