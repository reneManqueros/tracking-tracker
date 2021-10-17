package cmd

import (
	"github.com/spf13/cobra"
	"trackingtracker/internal/models"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new tracking code to be crawled",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			tl := models.TrackingLog{
				Code:    args[1],
				Service: args[0],
			}
			tl.Save()
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags()
}
