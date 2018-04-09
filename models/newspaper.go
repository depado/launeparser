package models

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/jaytaylor/html2text"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// NewsPaper is a struct representing a newspaper
type NewsPaper struct {
	Name   string `mapstructure:"name"`
	URL    string `mapstructure:"url"`
	Output string
}

// NewsPapers is a struct holding a slice of newspapers
type NewsPapers struct {
	NewsPapers []*NewsPaper `mapstructure:"newspapers"`
}

// CreateDirectories creates the raw output directories
func (ns *NewsPapers) CreateDirectories() {
	var err error
	output := viper.GetString("output")
	if _, err = os.Stat(output); os.IsNotExist(err) {
		if err = os.Mkdir(output, os.ModePerm); err != nil {
			logrus.WithError(err).Fatal("Could not create output directory")
		}
	}
	for _, n := range ns.NewsPapers {
		n.Output = filepath.Join(output, n.Name)
		if _, err = os.Stat(n.Output); os.IsNotExist(err) {
			if err = os.Mkdir(n.Output, os.ModePerm); err != nil {
				logrus.WithError(err).WithField("newspaper", n.Name).Fatal("Could not create output directory")
			}
		}
	}
}

// Scrape starts the routine to scrape all the newspapers in the slice
func (ns *NewsPapers) Scrape() {
	logrus.Info("Started scraping")
	ns.CreateDirectories()
	var wg sync.WaitGroup
	for _, n := range ns.NewsPapers {
		wg.Add(1)
		go func(n *NewsPaper) {
			defer wg.Done()
			n.Scrape()
		}(n)
	}
	wg.Wait()
	logrus.Info("Done scraping")
}

// Scrape scrapes a single newspaper
func (n *NewsPaper) Scrape() {
	var err error
	var resp *http.Response
	var fd *os.File
	var out string
	clog := logrus.WithField("newspaper", n.Name)

	c := &http.Client{
		Timeout: time.Second * 10,
	}
	if resp, err = c.Get(n.URL); err != nil {
		clog.WithError(err).Error("Couldn't scrape")
		return
	}
	defer resp.Body.Close()

	if fd, err = n.CreateDumpFile(); err != nil {
		clog.WithError(err).Error("Couldn't create dump file")
		return
	}
	out, err = html2text.FromReader(resp.Body)
	if _, err = fd.WriteString(out); err != nil {
		clog.WithError(err).Error("Couldn't copy ouput")
	}
}

// CreateDumpFile creates a file to write to with appropriate date format
func (n *NewsPaper) CreateDumpFile() (*os.File, error) {
	now := time.Now().Format("2006-01-02_15:04")
	out := filepath.Join(n.Output, fmt.Sprintf("%s.txt", now))
	return os.Create(out)
}
