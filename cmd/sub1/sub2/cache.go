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
