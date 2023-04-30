package printer

import (
	"io"
	"os"
)

func Print(toPrint []string, outfile *os.File) {
	_, _ = outfile.Seek(0, io.SeekStart)
	_ = outfile.Truncate(0)
	for i := 0; i < len(toPrint); i++ {
		outfile.WriteString(toPrint[i] + "\n")
	}
	outfile.Close()
}
