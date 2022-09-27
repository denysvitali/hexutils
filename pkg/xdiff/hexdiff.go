package xdiff

import (
	"bin-analysis/pkg/xtools"
	"fmt"
	"io"
)

type HexDiff struct {
	r1 xtools.Reader
	r2 xtools.Reader
}

func New(r1 io.Reader, r2 io.Reader) HexDiff {
	return HexDiff{
		r1: xtools.New(r1),
		r2: xtools.New(r2),
	}
}

func (h *HexDiff) Dump(w io.Writer) error {
	addr := uint32(0)
	var line1, line2 []byte
	var read1, read2 int
	var err1, err2 error

	for {
		if err1 == io.EOF {
			return nil
		}
		if err2 == io.EOF {
			return nil
		}

		if err1 != nil {
			return err1
		}

		if err2 != nil {
			return err2
		}

		line1, read1, err1 = h.r1.ReadLine()
		line2, read2, err2 = h.r2.ReadLine()

		if read1 == 0 && read2 == 0 {
			break
		}

		var length int
		length, err1 = h.r1.PrintLineDiff(addr, w, line1, line2)
		addr += uint32(length)
		fmt.Fprintf(w, "\n")
	}
	return nil
}
