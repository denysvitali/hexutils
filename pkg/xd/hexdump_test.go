package xd_test

import (
	"bin-analysis/pkg/xd"
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestReader_HexDump(t *testing.T) {
	for _, v := range []string{
		"example-1/1.bin",
		"example-1/2.bin",
		"example-1/3.bin",
		"example-1/4.bin",
	} {
		fmt.Printf("---\n")
		f, err := os.Open(fmt.Sprintf("./resources/%s", v))
		if err != nil {
			t.Fatalf("unable to open file: %v", err)
		}
		r := xd.New(f)

		b := bytes.NewBuffer(nil)
		err = r.Dump(b)
		if err != nil {
			t.Fatalf("unable to produce hexdump: %v", err)
		}

		fmt.Printf("%s\n", b.String())
	}
}
