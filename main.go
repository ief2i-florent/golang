package main

import (
	"fmt"

	"github.com/ief2i-florent/golang/internal/fssync"
	"github.com/spf13/cobra"
)

func main() {
	var root = &cobra.Command{
		Use:   "app [source] [target]",
		Short: "Permet de copier des fichiers",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			source := args[0]
			target := args[1]

			if fssync.FolderExist(source) {
				fssync.CreateFolderIfNotExist(target)
				for _, file := range fssync.Scan(source, source) {
					fssync.Copy(source+file, target+file)
				}
			} else {
				fmt.Println("Source folder not exist")
			}

		},
		DisableFlagsInUseLine: true,
	}

	root.Execute()
}
