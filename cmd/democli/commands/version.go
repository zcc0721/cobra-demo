package commands

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	version = "1.0.0"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hugo Static Site Generator v%s %s/%s\n", version, runtime.GOOS, runtime.GOARCH)
	},
}
