package cmd

import (
	"github.com/rsturla/golang-aio/internal/http"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "golang-aio",
	Short: "golang-aio is a full-stack GoLang application template",
	Long: `golang-aio is an all-in-one GoLang application that serves as a template for building
	modern full-stack applications using GoLang and NextJS.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return http.Serve(embedFS, cfg)
	},
}
