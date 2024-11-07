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
