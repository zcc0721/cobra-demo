package commands

import (
	"os"

	"github.com/bytom/util"
	"github.com/spf13/cobra"
)

// BytomcliCmd is Bytomcli's root command.
// Every other command attached to BytomcliCmd is a child command to it.
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

// Execute adds all child commands to the root command BytomcliCmd and sets flags appropriately.
func Execute() {
	// AddCommands add sub commands
	AddCommands()

	if _, err := DemocliCmd.ExecuteC(); err != nil {
		os.Exit(util.ErrLocalExe)
	}
}

// AddCommands adds child commands to the root command BytomcliCmd.
func AddCommands() {
	DemocliCmd.AddCommand(versionCmd)
}

// var demoCmd = &cobra.Command{
// 	Use:   "demo",
// 	Short: "demo is a commond line client for cobra demo",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		if len(args) < 1 {
// 			cmd.Usage()
// 			os.Exit(0)
// 		}
// 	},
// }
