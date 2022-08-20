package gs

import (
	"os"
	"path"

	"github.com/go-spring/gs/internal"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove spring-*/starter-*",
	Aliases: []string{"pl"},
	Short:   "pull remote code",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		rootDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		backup, err := cmd.Flags().GetBool("backup")
		if err != nil {
			panic("illegal backup value,must")
		}
		if backup {
			internal.Zip(rootDir)
		}

		remove(rootDir, projectName)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

// remove 移除远程项目
func remove(rootDir string, projectName string) {

	_, dir, project := validProject(projectName)
	internal.Remove(rootDir, project)

	projectDir := path.Join(rootDir, dir)
	_ = os.RemoveAll(projectDir)

	if _, err := os.Stat(projectDir); !os.IsNotExist(err) {
		panic(err)
	}

	projectXml.Remove(project)
	internal.Remotes(rootDir)
}
