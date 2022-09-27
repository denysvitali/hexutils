package xtools

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/fatih/color"
	"io"
)

func (r *Reader) PrintAddr(w io.Writer) {
	r.PrintSpecificAddr(w, uint32(r.addr))
}

func (r *Reader) PrintSpecificAddr(w io.Writer, addr uint32) {
	addrBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(addrBytes, addr)

	addressString := fmt.Sprintf("%4s:  ", hex.EncodeToString(addrBytes))
	if r.styleOptions.BoldAddress {
		color.New(color.Bold).Fprint(w, addressString)
	} else {
		fmt.Fprint(w, addressString)
	}
}
