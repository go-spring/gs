package gs

import (
	"os"

	"github.com/go-spring/gs/internal"
	"github.com/spf13/cobra"
)

var repairCmd = &cobra.Command{
	Use:     "repair spring-*/starter-* [branch]",
	Aliases: []string{"rp"},
	Short:   "repair project dir",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		branch := args[1]
		if branch == "" {
			branch = "main"
		}
		rootDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		repair(rootDir, projectName, branch)
	},
}

func init() {
	rootCmd.AddCommand(repairCmd)
}

// repair 修复远程项目的链接
func repair(rootDir string, projectName string, branch string) {
	_, dir, project := validProject(projectName)
	internal.Add(rootDir, project, dir, branch)
}
