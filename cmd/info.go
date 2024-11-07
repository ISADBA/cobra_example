package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use: "info",
	// 至少提供两个参数
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
