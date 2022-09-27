package xtools

import (
	"fmt"
	"io"
)

func (r *Reader) PrintLine(w io.Writer) error {
	lineLength := r.lineSize
	lineBytes := make([]byte, lineLength)
	read, err := r.r.Read(lineBytes)

	if err != nil {
		return err
	}

	r.PrintAddr(w)
	r.addr += lineLength

	// Print Hex
	for i := 0; i < read; i++ {
		if i%2 == 0 && i != 0 {
			fmt.Fprint(w, " ")
		}
		if r.styleOptions.AsciiHighlight && isAscii(lineBytes[i]) {
			r.styleOptions.AsciiHighlightColor.Fprintf(w, "%02x", lineBytes[i])
		} else {
			fmt.Fprintf(w, "%02x", lineBytes[i])
		}
	}

	if uint(read) < lineLength {
		for i := uint(read); i < lineLength; i++ {
			if i%2 == 0 {
				fmt.Fprint(w, " ")
			}
			fmt.Fprint(w, "..")
		}
	}

	fmt.Fprintf(w, "\t")

	// Print Text
	for i := 0; i < read; i++ {
		theByte := lineBytes[i]
		if isAscii(theByte) {
			if r.styleOptions.AsciiHighlight {
				r.styleOptions.AsciiHighlightColor.Fprintf(w, "%s", string(theByte))
			} else {
				fmt.Fprintf(w, "%s", string(theByte))
			}
		} else {
			fmt.Fprintf(w, ".")
		}
	}

	if uint(read) < lineLength {
		for i := uint(read); i < lineLength; i++ {
			fmt.Fprint(w, "•")
		}
	}
	return nil
}

func (r *Reader) PrintLineDiff(addr uint32, w io.Writer, line1 []byte, line2 []byte) (int, error) {
	lineLength := len(line1)
	if lineLength < len(line2) {
		lineLength = len(line2)
	}

	r.PrintSpecificAddr(w, addr)

	// Print Hex (L1)
	r.printLineHex(w, lineLength, line1, line2)
	fmt.Fprintf(w, "\t")
	r.printTextDiff(w, lineLength, line1, line2)

	fmt.Fprintf(w, "\t")
	r.printLineHex(w, lineLength, line2, line1)
	fmt.Fprintf(w, "\t")
	r.printTextDiff(w, lineLength, line2, line1)

	return lineLength, nil
}

func (r *Reader) printTextDiff(w io.Writer, lineLength int, line1 []byte, line2 []byte) {
	for i := 0; i < len(line1); i++ {
		theByte := line1[i]
		var resultString string
		if isAscii(theByte) {
			resultString = fmt.Sprintf("%s", string(theByte))
		} else {
			resultString = fmt.Sprintf(".")
		}

		if len(line2) <= i || theByte != line2[i] {
			// Different
			r.styleOptions.DifferenceColor.Fprintf(w, "%s", resultString)
		} else {
			fmt.Fprintf(w, "%s", resultString)
		}
	}

	if len(line1) < lineLength {
		for i := len(line1); i < lineLength; i++ {
			fmt.Fprint(w, "•")
		}
	}
}

func (r *Reader) printLineHex(w io.Writer, lineLength int, line1 []byte, line2 []byte) {
	for i := 0; i < lineLength; i++ {
		if i%2 == 0 && i != 0 || i >= len(line1) {
			fmt.Fprint(w, " ")
		}
		if i >= len(line1) {
			continue
		}

		if i >= len(line2) || line1[i] != line2[i] {
			// different
			r.styleOptions.DifferenceColor.Fprintf(w, "%02x", line1[i])
		} else {
			fmt.Fprintf(w, "%02x", line1[i])
		}
	}

	if len(line1) < lineLength {
		for i := len(line1); i < lineLength; i++ {
			if i%2 == 0 {
				fmt.Fprint(w, " ")
			}
			fmt.Fprint(w, "..")
		}
	}
}
