package parser

import (
	"errors"
	"flag"
	"os"
)

func ParseSources() (*os.File, *os.File, error) {
	sources := make([]string, 0)
	sources = append(sources, flag.Args()...)
	var err error = nil
	if len(sources) > 2 {
		err = errors.New("Wrong flag input. It should be like: \n" +
			"uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
	}
	if err != nil {
		return nil, nil, err
	}
	inpfile, outfile := os.Stdin, os.Stdout
	if len(sources) > 0 {
		inpfile, err = os.Open(sources[0])
		if err != nil {
			return nil, nil, err
		}
	}
	if len(sources) > 1 {
		outfile, err = os.OpenFile(sources[1], os.O_WRONLY, 0666)
		if err != nil {
			return inpfile, nil, err
		}
	}
	return inpfile, outfile, nil
}
