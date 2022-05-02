package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var root = &cobra.Command{
		Use:   "app [source] [destination]",
		Short: "Permet de copier des fichiers",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			source := args[0]
			destination := args[1]

			fmt.Println(source, destination)

		},
		DisableFlagsInUseLine: true,
	}

	root.Execute()
}
