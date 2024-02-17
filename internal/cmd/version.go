package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of golang-aio",
	Long:  `All software has versions. This is golang-aio's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("golang-aio v0.0.0-dev")
	},
}
