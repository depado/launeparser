package cmd

import (
	"github.com/Depado/launeparser/router"
	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server and scraping",
	Run: func(cmd *cobra.Command, args []string) {
		go func() {
			gocron.Every(1).Day().At("11:00").Do(NP.Scrape)
			gocron.Every(1).Day().At("23:00").Do(NP.Scrape)
			_, t := gocron.NextRun()
			logrus.WithField("next", t).Info("Registered tasks")
			<-gocron.Start()
		}()
		router.Run()
	},
}

func init() {
	startCmd.Flags().String("server.host", "127.0.0.1", "host on which the server should listen")
	startCmd.Flags().Int("server.port", 8080, "port on which the server should listen")
	startCmd.Flags().Bool("server.debug", false, "debug mode for the server")

	viper.BindPFlags(startCmd.Flags())
}
