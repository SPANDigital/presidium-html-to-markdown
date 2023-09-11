package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"htmltomarkdown/pkg"
	"log"
)

var config pkg.Config

func init() {
	cobra.OnInitialize(initConfig)
}

var rootCmd = &cobra.Command{
	Use: "root",
	Run: func(cmd *cobra.Command, args []string) {
		return
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.html2md")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return
		}
		log.Fatalf("fatal error config file: %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("fatal error config file: %s", err)
	}
}
