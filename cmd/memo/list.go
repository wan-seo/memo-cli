package memo

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wan-seo/memo-cli/pkg/memo"
)

var tag string
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List notes",
	Run: func(cmd *cobra.Command, args []string) {
		entries := memo.ReadMemoEntries()
		if tag != "" {
			entries = memo.FilterMemoEntries(entries, tag)
		}
		for idx, entry := range entries {
			fmt.Println(idx, entry)
		}
	},
}

func init() {
	listCmd.Flags().StringVarP(&tag, "filter-by", "t", "", "Filter by tag")
	rootCmd.AddCommand(listCmd)
}
