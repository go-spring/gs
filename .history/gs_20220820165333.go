package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/go-spring/gs/internal"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

// help 展示命令行用法
const help = `command:
  gs pull spring-*/starter-* [branch]
  gs repair spring-*/starter-* [branch]
  gs push spring-*/starter-*
  gs remove spring-*/starter-*
  gs release tag`

type Command struct {
	backup bool // 是否需要备份
	fn     func(rootDir string)
}

// commands 命令与处理函数的映射
var commands = map[string]Command{
	"pull":    {backup: true, fn: pull},    // 拉取单个远程项目
	"repair":  {backup: false, fn: repair}, // 拉取单个远程项目
	"push":    {backup: true, fn: push},    // 推送单个远程项目
	"remove":  {backup: true, fn: remove},  // 移除单个远程项目
	"release": {backup: true, fn: release}, // 发布所有远程项目
	"backup":  {backup: true, fn: nil},     // 备份本地项目文件
}

func test() {
	fmt.Println(help)
	defer func() { fmt.Println() }()

	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			os.Exit(-1)
		}
	}()

	command := arg(1)
	cmd, ok := commands[command]
	if !ok {
		panic("error command " + command)
	}

	// 获取工作目录
	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// 加载 project.xml 配置文件
	projectFile := path.Join(rootDir, "project.xml")
	err = projectXml.Read(projectFile)
	if err != nil {
		panic(err)
	}

	count := len(projectXml.Projects)
	defer func() {
		if count != len(projectXml.Projects) {
			// 保存 project.xml 配置文件
			err = projectXml.Save(projectFile)
			if err != nil {
				panic(err)
			}
		}
	}()

	fmt.Print(os.Args, " 输入 Yes 执行该命令: ")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	if strings.TrimSpace(input) != "Yes" {
		os.Exit(-1)
	}

	// 备份本地文件
	if cmd.backup {
		internal.Zip(rootDir)
	}

	// 执行命令
	if cmd.fn != nil {
		cmd.fn(rootDir)
	}
}