package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"htmltomarkdown/collector"
	"htmltomarkdown/converter"
	"htmltomarkdown/util"
	"net/url"
	"os"
	"path/filepath"
)

func init() {
	flags := convertCmd.Flags()
	flags.StringVar(&config.Html.Selector, "select", "body", "the part of the page to select and convert")
	flags.StringSliceVar(&config.Html.HeaderTags, "headers", []string{"h1", "h2"}, "article header tags")
	rootCmd.AddCommand(convertCmd)
}

var convertCmd = &cobra.Command{
	Use:   "convert [source] [dest]",
	Short: "convert",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(2), validateNPaths(2)),
	RunE:  run,
}

func run(_ *cobra.Command, args []string) error {
	var src, dst = args[0], args[1]
	if util.IsURL(src) {
		srcPath, err := collect(src)
		if err != nil {
			return err
		}
		defer os.RemoveAll(srcPath)

		parsedUrl, err := url.Parse(src)
		if err != nil {
			return err
		}

		src = filepath.Join(srcPath, parsedUrl.Path)
	}
	return converter.NewConverter(config).Convert(src, dst)
}

func collect(src string) (string, error) {
	tmp, err := os.MkdirTemp(".", "_clone")
	if err != nil {
		return tmp, err
	}

	if err := collector.Collect(src, tmp); err != nil {
		return tmp, err
	}

	return tmp, nil
}

func validateNPaths(n int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("requires at least 1 file or folder path")
		}

		for _, path := range args[:n] {
			if util.IsURL(path) {
				return nil
			}

			if _, err := os.Stat(path); os.IsNotExist(err) {
				return fmt.Errorf(`provided path "%s" does not exist`, path)
			}
		}
		return nil
	}
}
