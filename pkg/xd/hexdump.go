package xd

import (
	"bin-analysis/pkg/xtools"
	"fmt"
	"io"
)

type HexDump struct {
	reader xtools.Reader
}

func New(reader io.Reader) HexDump {
	return HexDump{
		reader: xtools.New(reader),
	}
}

func (h *HexDump) Dump(w io.Writer) error {
	var err error
	for {
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		nextErr := h.reader.PrintLine(w)

		err = nextErr
		fmt.Fprintf(w, "\n")
	}
}
