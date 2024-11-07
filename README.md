介绍
在golang语言下，开发一个命令行工具是比较容易的，原因是有cobra这样一个sdk。
Cobra 在许多 Go 项目中使用，例如 Kubernetes、Hugo 和 GitHub CLI 、Vitess、Etcd等。此列表包含使用 Cobra 的更广泛项目列表。
Cobra提供了命令行工具常用的一些能力，或者说99%场景的能力
比如：
1. 支持构建子命令，支持构建嵌套的子命令
2. 符合POSIX规范的长短命令
3. 支持全局(所有子命令可用)、本地的(针对单个子命令可用)、级联的参数标识
4. 子命令拼写错误可以只能推荐正确的子命令
5. 子命令分组的帮助信息
6. 自动的help标识识别以及内容生成
7. 可以无缝集成viper包

相关术语
使用cat -A a.txt命令来讲解
Commands
指的是命令或者子命令，比如cat就是一个命令
Args
代表的是操作的对象，也就是a.txt
Flags
命令的标识，也就是参数，也就是-A


最佳实践
1. 安装
go get -u github.com/spf13/cobra@latest
  1. 多个子命令的文件组织
▾ appName/
  ▾ cmd/
      add.go  // 一个子命令一个文件
      your.go
      root.go  // root命令
      here.go
  main.go // 程序入口,调用root命令
  2. 嵌套子命令的文件组织
▾ appName/
  ▾ cmd/
      add.go  // 一个子命令一个文件
      your.go
      root.go  // root命令
      sub1/
          sub1.go
          sub2/
              sub2.go
              clean.go
              cache.go
      
  main.go // 程序入口,调用root命令
  3. 命令的关联关系是在上层命令的init()方法中使用AddCommand()方法引入下层命令的入口。
比如root.go是根命令，那么root.go的init()方法会执行
func init() {
  rootCmd.AddCommand(addCmd)
  rootCmd.AddCommand(yourCmd)
  rootCmd.AddCommand(sub1Cmd)
}

在sub1.go中
func init() {
  rootCmd.AddCommand(sub2Cmd)
}

在sub2.go中
func init() {
  rootCmd.AddCommand(cleanCmd)
  rootCmd.AddCommand(cacheCmd)
}
2. 只有一个主命令的场景
// main.go

package main

import (
    "app/cmd"
)

func main() {
    cmd.Execute()
}


// cmd/root.go
package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "app",
    Short: "Try and possibly fail at something",
    RunE: func(cmd *cobra.Command, args []string) error {
        if err := someFunc(); err != nil {
            return err
        }
        return nil
    },
}

func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func someFunc() error {
    fmt.Println("I'm trying to do something...")
    return nil
}

// 执行
 go build && ./app
I'm trying to do something...
3. 给主命令加一个参数
  1. 方法一 rootCmd.PersistentFlags().StringP，全局可用；参数使用cmd.Flags().GetString获取
package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "app",
    Short: "Try and possibly fail at something",
    RunE: func(cmd *cobra.Command, args []string) error {
        if err := someFunc(cmd, args); err != nil {
            return err
        }
        return nil
    },
}

func init() {
    // 结果存储在configPath, err := cmd.Flags().GetString("config")
    rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.cobra.yaml)")
}

func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func someFunc(cmd *cobra.Command, args []string) error {
    fmt.Println("I'm trying to do something...")
    configPath, err := cmd.Flags().GetString("config")
    if err != nil {
        return err
    }
    fmt.Println("config path is ", configPath)
    return nil
}


  2. 方法二 rootCmd.PersistentFlags().StringVarP，全局可用，使用变量存储
package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "app",
    Short: "Try and possibly fail at something",
    RunE: func(cmd *cobra.Command, args []string) error {
        if err := someFunc(cmd, args); err != nil {
            return err
        }
        return nil
    },
}
var cfgFile string

func init() {
    // 结果存储在configPath, err := cmd.Flags().GetString("config")
    rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.cobra.yaml)")
    rootCmd.PersistentFlags().StringVarP(&cfgFile, "config2", "d", "config2.yaml", "config2 file (default is $HOME/.config2.yaml)")
}

func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func someFunc(cmd *cobra.Command, args []string) error {
    fmt.Println("I'm trying to do something...")
    configPath, err := cmd.Flags().GetString("config")
    if err != nil {
        return err
    }
    fmt.Println("config path is ", configPath)
    fmt.Println("config2 path is ", cfgFile)
    return nil
}


4. 添加一个子命令
// 在cmd中添加一个info.go

package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
    Use:   "info",
    Short: "Show information about the application",
    Long:  `Show information about the application`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("This is a Cobra application")
    },
}

// 修改root.go引入info.go里面的命令
func init() {
    // 结果存储在configPath, err := cmd.Flags().GetString("config")
    rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.cobra.yaml)")
    rootCmd.PersistentFlags().StringVarP(&cfgFile, "config2", "d", "config2.yaml", "config2 file (default is $HOME/.config2.yaml)")
    rootCmd.AddCommand(infoCmd)
}
5. 给子上面子命令添加四个参数,并设定一些使用规则
  1. 参数paramA必须存在，参数paramB和参数paramC只能存在一个，参数paramD可以指定，也可以从环境变量自动获取，至少显示传入一个参数
// 修改info.go
package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
    Use: "info",
    // 至少提供一个参数
    Args:  cobra.MinimumNArgs(1),
    Short: "Show information about the application",
    Long:  `Show information about the application`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("This is a Cobra application")
    },
}

func init() {
    // 获取环境变量的值
    envVarValue := os.Getenv("PARAM_D")

    // 如果环境变量没有设置，则使用空字符串作为默认值
    if envVarValue == "" {
        envVarValue = ""
    }
    infoCmd.Flags().StringP("paramA", "", "", "A parameter")
    infoCmd.Flags().StringP("paramB", "", "", "B parameter")
    infoCmd.Flags().StringP("paramC", "", "", "C parameter")
    infoCmd.Flags().StringP("paramD", "", envVarValue, "D parameter")
    // paramA标记为必填参数
    infoCmd.MarkFlagRequired("paramA")
    // paramB和paramC标记为互斥
    infoCmd.MarkFlagsMutuallyExclusive("paramB", "paramC")
}


6. 添加一个两级联子命令，第二级子命令有一个参数
.
├── app
├── cmd
│   ├── info.go
│   ├── root.go
│   └── sub1
│       ├── sub1.go
│       └── sub2
│           ├── cache.go
│           └── sub2.go
├── go.mod
├── go.sum
└── main.go


// root.go
package cmd

import (
    "app/cmd/sub1"
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "app",
    Short: "Try and possibly fail at something",
    RunE: func(cmd *cobra.Command, args []string) error {
        if err := someFunc(cmd, args); err != nil {
            return err
        }
        return nil
    },
}
var cfgFile string

func init() {
    // 结果存储在configPath, err := cmd.Flags().GetString("config")
    rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.cobra.yaml)")
    rootCmd.PersistentFlags().StringVarP(&cfgFile, "config2", "d", "config2.yaml", "config2 file (default is $HOME/.config2.yaml)")
    rootCmd.AddCommand(infoCmd)
    rootCmd.AddCommand(sub1.Sub1Cmd)
}

func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func someFunc(cmd *cobra.Command, args []string) error {
    fmt.Println("I'm trying to do something...")
    configPath, err := cmd.Flags().GetString("config")
    if err != nil {
        return err
    }
    fmt.Println("config path is ", configPath)
    fmt.Println("config2 path is ", cfgFile)
    return nil
}

// sub1.go
package sub1

import (
    "app/cmd/sub1/sub2"

    "github.com/spf13/cobra"
)

var Sub1Cmd = &cobra.Command{
    Use:   "sub1",
    Short: "This is a sub command of the root command",
    Long: `This is a sub command of the root command.
It has its own set of flags and arguments.
`,
}

func init() {
    Sub1Cmd.AddCommand(sub2.Sub2Cmd)
}



// sub2.go
package sub2

import (
    "github.com/spf13/cobra"
)

var Sub2Cmd = &cobra.Command{
    Use:   "sub2",
    Short: "sub2 command",
    Long:  `sub2 command`,
}

func init() {
    Sub2Cmd.AddCommand(cacheCmd)
}



// cache.go
package sub2

import (
    "fmt"

    "github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
    Use:   "cache",
    Short: "cache related commands",
    Long:  `cache related commands`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("cache called")
    },
}

func init() {
    cacheCmd.Flags().StringP("cache-path", "p", "", "cache path")
}



cobra-cli工具使用
参考文档：https://github.com/spf13/cobra-cli/blob/main/README.md
1. 安装cobra-cli
go install github.com/spf13/cobra-cli@latest
2. 项目初始化
cd $HOME/code 
mkdir myapp
cd myapp
go mod init myapp
3. 命令行工具代码初始化
cd $HOME/code/myapp
cobra-cli init
go run main.go
4. 添加其他子命令和带参数的命令
cobra-cli add serve
cobra-cli add config
cobra-cli add create -p 'configCmd'

  ▾ app/
    ▾ cmd/
        config.go
        create.go
        serve.go
        root.go
      main.go
5. 可以通过编写~/.cobra.yaml配置文件定制命令行工具
author: Steve Francia <spf@spf13.com>
year: 2020
license:
  header: This file is part of CLI application foo.
  text: |
    {{ .copyright }}

    This is my license. There are many like it, but this one is mine.
    My license is my best friend. It is my life. I must master it as I must
    master my life.
6. 

其他
1. cobra.Command{}对象支持的参数和方法都在 https://github.com/spf13/cobra/blob/main/command.go文件中定义了，想要了解更多深入的用法，可以查看项目源码。