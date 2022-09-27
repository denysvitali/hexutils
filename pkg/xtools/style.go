package xtools

import "github.com/fatih/color"

type StyleOptions struct {
	BoldAddress         bool
	AsciiHighlight      bool
	AsciiHighlightColor *color.Color

	DifferenceColor *color.Color
}

var DefaultStyle = StyleOptions{
	BoldAddress:         true,
	AsciiHighlight:      true,
	AsciiHighlightColor: color.New(color.BgYellow, color.FgBlack),
	DifferenceColor:     color.New(color.BgRed, color.FgBlack),
}
