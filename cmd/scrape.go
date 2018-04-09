package cmd

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Instantly scrape",
	Run: func(cmd *cobra.Command, args []string) {
		n := time.Now()
		NP.Scrape()
		logrus.WithField("took", time.Since(n)).Info("Done")
	},
}
