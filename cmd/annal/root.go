package main

import (
	"fmt"
	"os"

	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var cfgFile string
var verbose bool

var rootCmd = &cobra.Command{
	Use:   "genVideo",
	Short: "generate video for AHAOAHA",
	Long:  `generate video for AHAOAHA.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

}

// initEnv reads in ENV variables.
func initEnv() {
	log.AddHook(filename.NewHook())
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, TimestampFormat: "2006-01-02 15:04:05"})
	if verbose {
		log.SetLevel(log.DebugLevel)
	}

}
