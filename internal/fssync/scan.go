package fssync

import (
	"os"
	"strings"
)

func Scan(path string, base string) []string {
	list := make([]string, 0)

	dirEntries, _ := os.ReadDir(path)
	for _, entry := range dirEntries {
		if entry.IsDir() {
			list = append(list, Scan(path+"/"+entry.Name(), base)...)
		} else {
			relative := strings.Replace(path, base, "", 1)
			list = append(list, relative+"/"+entry.Name())
		}
	}

	return list
}
