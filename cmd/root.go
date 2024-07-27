package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli-installer",
	Short: "CLI Installer for various technologies",
	Long:  `CLI Installer helps you install, update, and check versions of various technologies like Golang, Laravel, Java, Node.js, and React.js.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Note: We don't need to add subcommands here because we handle that in main.go
}
