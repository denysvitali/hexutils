package main

import (
	"bin-analysis/pkg/xdiff"
	"github.com/alexflint/go-arg"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"os"
)

var args struct {
	Input1      string `arg:"positional,required"`
	Input2      string `arg:"positional,required"`
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

	f1, err := os.Open(args.Input1)
	if err != nil {
		logger.Fatalf("unable to open file: %v", err)
	}

	f2, err := os.Open(args.Input2)
	if err != nil {
		logger.Fatalf("unable to open file: %v", err)
	}

	r := xdiff.New(f1, f2)
	err = r.Dump(os.Stdout)
	if err != nil {
		logger.Fatalf("unable to create hexdump: %v", err)
	}
}
