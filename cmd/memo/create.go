package memo

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wan-seo/memo-cli/pkg/memo"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"cr"},
	Short:   "Create a note",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		created := memo.CreateMemo(args[0])
		fmt.Printf("Created %s\n", created)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
