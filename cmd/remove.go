package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Delete tracking code from crawls",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			filename := fmt.Sprintf(`./data/%s.json`, args[0])
			err := os.Remove(filename)
			if err != nil {
				log.Println(err)
			} else {
				log.Println("Removed:", args[0])
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags()
}
