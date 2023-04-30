package reader

import (
	"bufio"
	"os"
)

func Read(inpfile *os.File) []string {
	args := make([]string, 0)
	scanner := bufio.NewScanner(inpfile)
	for scanner.Scan() {
		args = append(args, scanner.Text())
	}
	inpfile.Close()
	return args
}
