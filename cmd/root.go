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
