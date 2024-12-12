package main

import (
	log "github.com/sirupsen/logrus"
	"htmltomarkdown/cmd"
)

func main() {
	log.SetLevel(log.DebugLevel)
	cmd.Execute()
}
