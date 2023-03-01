package memo

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "memo",
	Version: version,
	Short:   "memo is a quick note-taking CLI",
	Long: `memo is a CLI written in Go.
				  kindness matters!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kindness matters!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
