package gs

import (
	"os"

	"github.com/go-spring/gs/internal"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:     "pull spring-*/starter-* [branch]",
	Aliases: []string{"pl"},
	Short:   "pull remote code",
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

		backup, err := cmd.Flags().GetBool("backup")
		if err != nil {
			panic("illegal backup value,must")
		}
		if backup {
			internal.Zip(rootDir)
		}
		pull(rootDir, projectName, branch)
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}

// pull 拉取远程项目
func pull(rootDir string, projectName string, branch string) {
	_, dir, project := validProject(projectName)
	internal.SafeStash(rootDir, func() {
		remotes := internal.Remotes(rootDir)
		if internal.ContainsString(remotes, project) < 0 {
			add := false
			defer func() {
				if !add {
					remove(rootDir, projectName)
				}
			}()
			repository := internal.Add(rootDir, project, dir, branch)
			projectXml.Add(internal.Project{
				Name:   project,
				Dir:    dir,
				Url:    repository,
				Branch: branch,
			})
			add = true
		}
		internal.Sync(rootDir, project, dir, branch)
	})
}
