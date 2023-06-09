/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"calcsizev1/FileWordCount"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "calcsizev1",
	Short: "Calculate the number of words inside directory files.",
	Long:  `Calculate the number of words inside directory files.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Errorf("Error in file path flag!")
			os.Exit(1)
		}
		totalWordsInDir := FileWordCount.CountTotalWordInDir(path)
		log.Infof("Total words in the directory %s: %d", path, totalWordsInDir)
	},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Version: "1.0.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.calcsizev1.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("path", "p", "C:\\Users\\Lenovo\\GolandProjects\\createfiles\\data", "The path of the directory to generate file on")
}
