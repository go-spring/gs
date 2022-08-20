package gs

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-spring/gs/internal"
	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:     "release tag",
	Aliases: []string{"rs"},
	Short:   "release tag",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tag := args[0]
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
		release(rootDir, tag)
	},
}

func init() {
	rootCmd.AddCommand(releaseCmd)
}

// release 发布所有远程项目
func release(rootDir string, tag string) {
	// tag := arg(2)
	err := filepath.Walk(rootDir, func(walkFile string, _ os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if path.Base(walkFile) != "go.mod" {
			return nil
		}

		fmt.Println(walkFile)
		fileData, e0 := ioutil.ReadFile(walkFile)
		if e0 != nil {
			return nil
		}

		outBuf := bytes.NewBuffer(nil)
		r := bufio.NewReader(strings.NewReader(string(fileData)))
		for {
			line, isPrefix, e1 := r.ReadLine()
			if len(line) > 0 && e1 != nil {
				panic(e1)
			}
			if isPrefix {
				panic(errors.New("ReadLine returned prefix"))
			}
			if e1 != nil {
				if e1 != io.EOF {
					panic(err)
				}
				break
			}
			s := strings.TrimSpace(string(line))
			if strings.HasPrefix(s, "github.com/go-spring/spring-") ||
				strings.HasPrefix(s, "github.com/go-spring/starter-") {
				index := strings.LastIndexByte(s, ' ')
				if index <= 0 {
					panic(errors.New(s))
				}
				b := append(line[:index+2], []byte(tag)...)
				outBuf.Write(b)
			} else {
				outBuf.Write(line)
			}
			outBuf.WriteString("\n")
		}

		fmt.Println(outBuf.String())
		return ioutil.WriteFile(walkFile, outBuf.Bytes(), os.ModePerm)
	})

	if err != nil {
		panic(err)
	}

	// 提交代码更新
	internal.Commit(rootDir, "publish "+tag)

	// 遍历所有项目，推送远程更新
	for _, project := range projectXml.Projects {
		_, dir, _ := validProject(project.Name)
		internal.Push(rootDir, project.Name, dir, project.Branch)
	}

	// 创建临时目录
	now := time.Now().Format("20060102150405")
	buildDir := path.Join(rootDir, "..", "go-spring-build-"+now)
	err = os.MkdirAll(buildDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// 遍历所有项目，推送远程标签
	for _, project := range projectXml.Projects {
		projectDir := internal.Clone(buildDir, project.Name, project.Url)
		internal.Release(projectDir, tag)
	}
}
