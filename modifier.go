package piscine

import (
	"strings"
	"strconv"
)

// Converting 2 Decimal 
func Hex2dec(hexVal string) string {
	base := 1
	decimal := 0

	for i := len(hexVal) - 1; i >= 0; i-- {
		if hexVal[i] >= '0' && hexVal[i] <= '9' {
			decimal += (int(hexVal[i]) - 48) * base
			base = base * 16
		}
		if hexVal[i] >= 'A' && hexVal[i] <= 'F' {
			decimal += (int(hexVal[i]) - 55) * base
			base = base * 16
		}
	}
	return strconv.Itoa(decimal)
}

func Bin2dec(BinVal string) string {
	base := 1
	decimal := 0
	for i := len(BinVal) - 1; i >= 0; i-- {
		if BinVal[i] >= '0' && BinVal[i] <= '9' {
			decimal += (int(BinVal[i]) - 48) * base
			base = base * 2
		}
	}
	return strconv.Itoa(decimal)
}

// Editing punctuations 
func Punc(data string) string {
    nData := ""
    for _,ch := range data {
        if ch != '.' && ch != ',' && ch != '!' && ch != '?' && ch != ':' && ch != ';' {
            nData += string(ch)
        } else  {
            nData +=  string(ch) + " "
        } 
    }
    sub := strings.Fields(nData)
	editedData := ""
	for i := 0; i < len(sub); i++ {
		editedData += sub[i]
		// to check the individual punc marks 
		if i < len(sub)-1 {
			if sub[i+1] != "." && sub[i+1] != "," && sub[i+1] != "!" && sub[i+1] != "?" && sub[i+1] != ":" && sub[i+1] != ";" {
				editedData += " "
			}
		}
	}
	return editedData
}
