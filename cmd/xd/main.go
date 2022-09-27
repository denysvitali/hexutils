package main

import (
	"bin-analysis/pkg/xd"
	"github.com/alexflint/go-arg"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"os"
)

var args struct {
	InputFile   string `arg:"positional,required"`
	ColorAlways *bool  `arg:"--color"`
	Debug       *bool  `arg:"-D,--debug"`
}

var logger = logrus.New()

func main() {
	arg.MustParse(&args)

	if args.Debug != nil && *args.Debug {
		logger.SetLevel(logrus.DebugLevel)
	}

	if args.ColorAlways != nil && *args.ColorAlways {
		color.NoColor = false
	}

	f, err := os.Open(args.InputFile)
	if err != nil {
		logger.Fatalf("unable to open file: %v", err)
	}

	r := xd.New(f)
	err = r.Dump(os.Stdout)
	if err != nil {
		logger.Fatalf("unable to create hexdump: %v", err)
	}
}
