package cmd

import (
	"embed"
	"fmt"
	"os"

	"github.com/rsturla/golang-aio/internal/config"
	"github.com/rsturla/golang-aio/internal/utils"
)

// EmbedFS must be passed in from the main package since go:embed cannot view
// objects in parent directories.
var embedFS embed.FS
var cfg *config.Config

const configFileEnvPrefix = "GOLANG_AIO_"

func Execute(embed embed.FS) error {
	embedFS = embed

	if err := utils.SetupLogger(); err != nil {
		return err
	}

	configFileEnvName := fmt.Sprintf("%sCONFIG_FILE", configFileEnvPrefix)
	c, err := utils.SetupConfig(os.Getenv(configFileEnvName), configFileEnvPrefix)
	if err != nil {
		return err
	}
	cfg = c

	setupCommands()

	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func setupCommands() {
	rootCmd.AddCommand(versionCmd)
}
