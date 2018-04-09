package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Run setups and runs the server
func Run() {
	var err error

	// Debug mode
	if !viper.GetBool("server.debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	// Router initialization
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Writer.Header().Add("Content-Disposition", `attachment; filename="output.zip"`)
		c.Status(http.StatusOK)
		if err := zipto(viper.GetString("output"), c.Writer); err != nil {
			logrus.WithError(err).Error("Something went wrong")
		}
	})

	// Run
	logrus.WithFields(logrus.Fields{
		"host": viper.GetString("server.host"),
		"port": viper.GetInt("server.port"),
	}).Info("Starting server")

	if err = r.Run(fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port"))); err != nil {
		logrus.WithError(err).Fatal("Couldn't start server")
	}
}
