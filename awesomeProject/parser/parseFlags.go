package parser

import (
	"awesomeProject/structs"
	"errors"
	"flag"
)

func ParseFlags() (structs.FlagSet, error) {
	var flagSet structs.FlagSet
	flagc := flag.Bool("c", false, "count the number of occurrences of a string in the input")
	flagd := flag.Bool("d", false, "output only those lines that are repeated in the input")
	flagu := flag.Bool("u", false, "output only those lines that are repeated in the input")
	flagi := flag.Bool("i", false, "ignore register")
	flag.IntVar(&flagSet.FCnt, "f", 0, "ignore fcnt fields during comparing")
	flag.IntVar(&flagSet.CCnt, "s", 0, "ignore ccnt characters during comparing")
	flag.Parse()
	flagSet.Flagi = *flagi
	flagSet.PrintType = parsePrintType(*flagc, *flagd, *flagu)
	errOfFlags := checkFlagCorr(flagSet)
	return flagSet, errOfFlags
}

func parsePrintType(flagCVal bool, flagDVal bool, flagUVal bool) string {
	res := ""
	if flagCVal {
		res += "c"
	}
	if flagDVal {
		res += "d"
	}
	if flagUVal {
		res += "u"
	}
	return res
}

func checkFlagCorr(flagSet structs.FlagSet) error {
	if len(flagSet.PrintType) > 1 || flagSet.FCnt < 0 || flagSet.CCnt < 0 {
		return errors.New("Wrong flag input. It should be like: \n" +
			"uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
	} else {
		return nil
	}
}
