package commands

import (
	"os"

	"github.com/spf13/cobra"
)

// DemocliCmd is Democli's root command.
// Every other command attached to DemocliCmd is a child command to it.
var DemocliCmd = &cobra.Command{
	Use:   "democli",
	Short: "Democli is a commond line client for cobra demo",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
			os.Exit(0)
		}
	},
}

// Execute adds all child commands to the root command DemocliCmd and sets flags appropriately.
func Execute() {
	// AddCommands add sub commands.
	AddCommands()

	if err := DemocliCmd.Execute(); err != nil {
		os.Exit(0)
	}
}

// AddCommands adds child commands to the root command DemocliCmd.
func AddCommands() {
	DemocliCmd.AddCommand(versionCmd)
}
