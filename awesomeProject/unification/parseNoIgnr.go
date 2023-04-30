package unification

func parseNoIgnr(args []string, flagi bool, printType string) []string {
	var uniqStrings []stringWithRepCnt
	for i, arg := range args {
		var isFound = false
		for j, uniqString := range uniqStrings {
			if CompareStrings(arg, uniqString.Str, flagi) {
				uniqStrings[j].Cnt += 1
				isFound = true
				break
			}
		}
		if !isFound {
			uniqStrings = append(uniqStrings, stringWithRepCnt{args[i], 1})
		}
	}
	return chooseToPrint(uniqStrings, printType)
}
