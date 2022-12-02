package main

import (
	"fmt"
	"os"

	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "annal",
	Short: "annal tools.",
	Long:  `annal tools for everything.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	initEnv()
}

// initEnv reads in ENV variables.
func initEnv() {
	log.AddHook(filename.NewHook())
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, TimestampFormat: "2006-01-02 15:04:05"})
	if verbose {
		log.SetLevel(log.DebugLevel)
	}
}
