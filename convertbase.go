package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	nbrDecimal := baseToDecimal(nbr, baseFrom)
	return decimalToBase(nbrDecimal, baseTo)
}

func baseToDecimal(nbr string, baseFrom string) int {
	nbrDecimal := 0
	for i := 0; i < len(nbr); i++ {
		nbrDecimal = nbrDecimal*len(baseFrom) + Index(baseFrom, string(rune(nbr[i])))
	}
	return nbrDecimal
}

func decimalToBase(nbr int, baseTo string) string {
	nbrBase := ""
	for nbr > 0 {
		nbrBase = string(rune(baseTo[nbr%len(baseTo)])) + nbrBase
		nbr /= len(baseTo)
	}
	return nbrBase
}
