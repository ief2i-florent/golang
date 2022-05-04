package fssync

import (
	"fmt"
	"os"
	"path"
)

func Copy(source string, target string) error {

	fi, err := os.Open(source)
	defer fi.Close()
	if err != nil {
		return err
	}

	parentFolder := path.Dir(target)
	err = os.MkdirAll(parentFolder, 0755)
	if err != nil {
		return err
	}

	fo, err := os.Create(target)
	defer fo.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	buffer := make([]byte, 1024)
	for {
		n, err := fi.Read(buffer)
		if n == 0 {
			break
		} else if err != nil {
			return err
		}
		if _, err := fo.Write(buffer[:n]); err != nil {
			break
		}
	}

	return nil
}
