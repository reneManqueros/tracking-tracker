package cmd

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"strings"
	"trackingtracker/internal/models"
	"trackingtracker/internal/services/cainiao"
	"trackingtracker/internal/services/fedex"
	"trackingtracker/internal/services/usps"
)

var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Crawl all registered tracking codes",
	Long:  "Usually this is called via a cron job",
	Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir("./data")
		if err != nil {
			log.Fatal(err)
		}
		services := map[string]func(models.TrackingLog){
			"cainiao": cainiao.Crawl,
			"usps":    usps.Crawl,
			"fedex":   fedex.Crawl,
		}
		for _, file := range files {
			if err != nil {
				log.Println(err)
			}
			tl := models.TrackingLog{
				Code: strings.Replace(file.Name(), ".json", "", -1),
			}
			tl.Load()
			services[tl.Service](tl)
		}
	},
}

func init() {
	rootCmd.AddCommand(crawlCmd)
	crawlCmd.Flags()
}
