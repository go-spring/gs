package gs

import (
	"os"

	"github.com/go-spring/gs/internal"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:     "push spring-*/starter-*",
	Aliases: []string{"ps"},
	Short:   "push local code to remote repo",
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
		push(rootDir, projectName)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}

// push 推送远程项目
func push(rootDir string, projectName string) {

	_, dir, project := validProject(projectName)
	internal.SafeStash(rootDir, func() {

		// 将修改提交到远程项目，不需要往回合并
		if p, ok := projectXml.Find(project); ok {
			internal.Push(rootDir, project, dir, p.Branch)
		}
	})
}
