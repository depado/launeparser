package cmd

import (
	"github.com/Depado/launeparser/models"
	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NP is the main list of NewsPapers to scrape
var NP *models.NewsPapers

var rootCmd = &cobra.Command{
	Use:   "launeparser",
	Short: "Launeparser scrapes newspapers",
}

// Execute executes the commands
func Execute(b, v string) {
	Build = b
	Version = v
	rootCmd.AddCommand(version, startCmd, scrapeCmd)

	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal()
	}
}

func init() {
	cobra.OnInitialize(initialize)

	// Global flags
	rootCmd.PersistentFlags().String("log.level", "info", "one of debug, info, warn, error or fatal")
	rootCmd.PersistentFlags().String("log.format", "text", "one of text or json")
	rootCmd.PersistentFlags().Bool("log.line", false, "enable filename and line in logs")
	rootCmd.PersistentFlags().String("output", "out", "output directory")

	// Flag binding
	viper.BindPFlags(rootCmd.PersistentFlags())
}

func initialize() {
	// Environment variables
	viper.AutomaticEnv()

	// Configuration file
	viper.SetConfigName("conf")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/config/")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal("No configuration file found")
	}

	if err := viper.Unmarshal(&NP); err != nil {
		logrus.WithError(err).Fatal("Couldn't unmarshal newspaper")
	}

	lvl := viper.GetString("log.level")
	l, err := logrus.ParseLevel(lvl)
	if err != nil {
		logrus.WithField("level", lvl).Warn("Invalid log level, fallback to 'info'")
	} else {
		logrus.SetLevel(l)
	}
	switch viper.GetString("log.format") {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
	if viper.GetBool("log.line") {
		logrus.AddHook(filename.NewHook())
	}
}
