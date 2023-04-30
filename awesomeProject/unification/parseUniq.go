package unification

import "awesomeProject/structs"

func ParsingUnification(flagSet structs.FlagSet, args []string) []string {
	if flagSet.FCnt > 0 || flagSet.CCnt > 0 {
		return parseWithIgnr(args, flagSet)
	} else {
		return parseNoIgnr(args, flagSet.Flagi, flagSet.PrintType)
	}
}
