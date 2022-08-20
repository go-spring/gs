package gs

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/go-spring/gs/internal"
	"github.com/spf13/cobra"
)

// 配置文件
var projectXml internal.ProjectXml
var springProject = regexp.MustCompile("spring-.*")
var starterProject = regexp.MustCompile("starter-.*")

// validProject 项目名称是否有效，返回项目前缀、项目目录、项目名称
func validProject(project string) (prefix string, dir string, _ string) {
	if !springProject.MatchString(project) && !starterProject.MatchString(project) {
		panic("error project " + project)
	}
	prefix = strings.Split(project, "-")[0]
	return prefix, fmt.Sprintf("%s/%s", prefix, project), project
}

var rootCmd = &cobra.Command{
	Use:   "gs",
	Short: "gs - a simple CLI to transform and inspect strings",
	Long: `gs is a super fancy CLI (kidding)
   
One can use gs to add or modfiy go spring project from the terminal`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	log.SetFlags(log.Lshortfile)
	rootCmd.PersistentFlags().BoolP("backup", "b", true, "backup project code.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
