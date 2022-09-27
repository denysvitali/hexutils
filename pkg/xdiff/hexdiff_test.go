package xdiff_test

import (
	"bin-analysis/pkg/xdiff"
	"os"
	"testing"
)

func mustOpen(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return f
}

func TestHexDiff_Dump(t *testing.T) {
	h := xdiff.New(
		mustOpen("./resources/example-1/1.bin"),
		mustOpen("./resources/example-1/2.bin"),
	)
	err := h.Dump(os.Stdout)
	if err != nil {
		t.Fatalf("err=%v", err)
	}
}
