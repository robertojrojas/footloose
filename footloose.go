package main

import (
	"os"
	"path"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Footloose is the default name of the footloose file.
const Footloose = "footloose.yaml"

var footloose = &cobra.Command{
	Use:           "footloose",
	Short:         "footloose - Container Machines",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func configFile(f string) string {
	fp := f
	if env := os.Getenv("FOOTLOOSE_CONFIG"); env != "" && f == Footloose {
		fp = env
	}
	return toAbs(fp)
}

func toAbs(p string) string {
	ap := p
	if !path.IsAbs(ap) {
		ap, err := filepath.Abs(ap)
		// if Abs reports and error just return the original path 'p'
		if err != nil {
			ap = p
		}
	}
	return ap
}

func main() {
	if err := footloose.Execute(); err != nil {
		log.Fatal(err)
	}
}
