package unification

import "strconv"

func chooseToPrint(toChoose []stringWithRepCnt, printType string) []string {
	var ans []string
	switch printType {
	case "": //without flags
		for _, pairToChoose := range toChoose {
			ans = append(ans, pairToChoose.Str)
		}
	case "c": //flag c
		for _, pairToChoose := range toChoose {
			ans = append(ans, strconv.Itoa(pairToChoose.Cnt)+" "+pairToChoose.Str)
		}
	case "d": //flag d
		for _, pairToChoose := range toChoose {
			if pairToChoose.Cnt > 1 {
				ans = append(ans, pairToChoose.Str)
			}
		}
	case "u": //flag u
		for _, pairToChoose := range toChoose {
			if pairToChoose.Cnt == 1 {
				ans = append(ans, pairToChoose.Str)
			}
		}
	}
	return ans
}
