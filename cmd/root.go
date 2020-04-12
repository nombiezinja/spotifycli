package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "spotify",
		Short: "Spotify cli",
		Long:  `Placeholder`,
		Run:   func(cmd *cobra.Command, args []string) { fmt.Println("testy test test") },
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
