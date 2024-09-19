package cmd

import (
	"errors"
	"fmt"
	"htmltomarkdown/collector"
	"htmltomarkdown/converter"
	"htmltomarkdown/util"
	"net/url"
	"os"
	"path/filepath"

	cp "github.com/otiai10/copy"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	flags := convertCmd.Flags()
	flags.StringVar(&cfg.Html.Selector, "select", "body", "the part of the page to select and convert")
	flags.StringSliceVar(&cfg.Html.HeaderTags, "headers", []string{"h1", "h2"}, "article header tags")
	flags.BoolVarP(&cfg.Debug, "debug", "d", false, "enable debug logging")
	rootCmd.AddCommand(convertCmd)
}

var convertCmd = &cobra.Command{
	Use:   "convert [source] [dest]",
	Short: "convert",
	Long:  "CLI tool for converting HTML to Presidium ready Markdown content",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(2), validateNPaths(2)),
	RunE:  run,
}

func run(_ *cobra.Command, args []string) error {
	if cfg.Debug {
		log.SetLevel(log.DebugLevel)
	}

	var src, dst = args[0], args[1]
	var baseUrl = src
	if util.IsURL(src) {
		clonePath, err := collect(src)
		if err != nil {
			return err
		}
		defer os.RemoveAll(clonePath)

		parsedUrl, err := url.Parse(src)
		if err != nil {
			return err
		}

		assetSrc := filepath.Join(clonePath, cfg.AssetDir)
		assetDst := filepath.Join(dst, cfg.AssetDir)
		if err := cp.Copy(assetSrc, assetDst); err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		}

		baseUrl = parsedUrl.Path
		src = filepath.Join(clonePath, baseUrl)
	}

	dst = filepath.Join(dst, cfg.ContentDir)
	return converter.NewConverter(baseUrl, cfg).Convert(src, dst)
}

func collect(src string) (string, error) {
	tmp, err := os.MkdirTemp(".", "_clone")
	if err != nil {
		return tmp, err
	}

	if err := collector.Collect(src, tmp, cfg); err != nil {
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
