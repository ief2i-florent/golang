package main

import (
	"fmt"

	"github.com/ief2i-florent/golang/internal/fssync"
	"github.com/ief2i-florent/golang/internal/multithread"
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
				for _, file := range fssync.Scan(source, source) {
					multithread.AddInQueue(
						multithread.MakeMission(
							source+file,
							target+file,
						),
					)
				}

				multithread.Run()
			} else {
				fmt.Println("Source folder not exist")
			}

		},
		DisableFlagsInUseLine: true,
	}

	root.Execute()
}
