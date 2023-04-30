package unification

import (
	"awesomeProject/structs"
	"strings"
)

func parseWithIgnr(toIgnr []string, flagSet structs.FlagSet) []string {
	ignrdStrings := make([]string, 0)
	var uniqStrings []stringWithRepCnt
	for i, nowWatch := range toIgnr {
		if flagSet.FCnt > 0 {
			nowWatch = parseFieldsIgnr(nowWatch, flagSet.FCnt)
		}
		if flagSet.CCnt > 0 {
			nowWatch = parseCharIgnr(nowWatch, flagSet.CCnt)
		}
		var isFound = false
		for j, ignrdString := range ignrdStrings {
			if CompareStrings(nowWatch, ignrdString, flagSet.Flagi) {
				uniqStrings[j].Cnt++
				isFound = true
				break
			}
		}
		if !isFound {
			uniqStrings = append(uniqStrings, stringWithRepCnt{toIgnr[i], 1})
			ignrdStrings = append(ignrdStrings, nowWatch)
		}
	}
	return chooseToPrint(uniqStrings, flagSet.PrintType)
}
func parseFieldsIgnr(strToParse string, fCnt int) string {
	if strings.Count(strToParse, " ") == 0 {
		return strToParse
	}
	if fCnt > strings.Count(strToParse, " ") {
		return ""
	}
	return strings.Join(strings.Fields(strToParse)[fCnt:], " ")
}
func parseCharIgnr(strToParse string, cCnt int) string {
	if cCnt >= len(strToParse) {
		return ""
	}
	return strToParse[cCnt:]
}
