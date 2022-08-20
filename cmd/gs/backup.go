package gs

import (
	"os"

	"github.com/go-spring/gs/internal"
	"github.com/spf13/cobra"
)

var backpupCmd = &cobra.Command{
	Use:     "backup",
	Aliases: []string{"bak"},
	Short:   "backup a project",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		rootDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		internal.Zip(rootDir)
	},
}

func init() {
	rootCmd.AddCommand(backpupCmd)
}
