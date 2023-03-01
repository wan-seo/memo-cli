package memo

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wan-seo/memo-cli/pkg/memo"
)

var readCmd = &cobra.Command{
	Use:     "read",
	Aliases: []string{"cat", "r"},
	Short:   "Read a note",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		read := memo.ReadMemo(args[0])
		fmt.Printf("%s\n", read)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}
