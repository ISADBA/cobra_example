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
