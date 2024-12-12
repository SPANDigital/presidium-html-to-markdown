package cmd

import (
	"htmltomarkdown/config"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfg config.Config

func init() {
	cobra.OnInitialize(initConfig)
}

var rootCmd = &cobra.Command{
	Use: "html2md",
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	log.SetLevel(log.InfoLevel)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.html2md")

	viper.SetDefault("assetDir", "assets")
	viper.SetDefault("contentDir", "content")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return
		}
		log.Fatalf("fatal error config file: %s", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("fatal error config file: %s", err)
	}
}
