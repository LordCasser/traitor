package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"traitor/pkg/exploits"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available methods.",
	Run: func(cmd *cobra.Command, args []string) {
		allExploits := exploits.Get(exploits.SpeedAny)
		for _, exploit := range allExploits {
			fmt.Println(exploit.Name)
		}
	},
}
