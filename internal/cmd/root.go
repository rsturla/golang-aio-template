package cmd

import (
	"embed"
	"fmt"
	"os"

	"github.com/rsturla/golang-aio/internal/config"
	"github.com/rsturla/golang-aio/internal/http"
	"github.com/rsturla/golang-aio/internal/utils"
	"github.com/spf13/cobra"
)

// EmbedFS must be passed in from the main package since go:embed cannot view
// objects in parent directories.
var EmbedFS embed.FS
var Config *config.Config

const configFileEnvPrefix = "GOLANG_AIO_"

var rootCmd = &cobra.Command{
	Use:   "golang-aio",
	Short: "golang-aio is a full-stack GoLang application template",
	Long: `golang-aio is an all-in-one GoLang application that serves as a template for building
	modern full-stack applications using GoLang and NextJS.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return http.Serve(EmbedFS, Config)
	},
}

func Execute(embed embed.FS) error {
	EmbedFS = embed

	if err := utils.SetupLogger(); err != nil {
		return err
	}

	configFileEnvName := fmt.Sprintf("%sCONFIG_FILE", configFileEnvPrefix)
	cfg, err := utils.SetupConfig(os.Getenv(configFileEnvName), configFileEnvPrefix)
	if err != nil {
		return err
	}

	Config = cfg

	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
