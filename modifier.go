package piscine

import (
	"strconv"
	"strings"
)

// converting numbers
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

// Editting the text
func ToUpperCase(str string) string {
	var newStr string // to store the edited version of str
	for _, ch := range str {
		if !(ch >= 'A' && ch <= 'Z') {
			newStr += string(ch - 32)
		} else {
			newStr += string(ch)
		}
	}
	return newStr
}

func ToLowerCase(str string) string {
	var newStr string // to store the edited version of str
	for _, ch := range str {
		if !(ch >= 'a' && ch <= 'z') {
			newStr += string(ch + 32)
		} else {
			newStr += string(ch)
		}
	}
	return newStr
}

func Capitalize(str string) string {
	flag := true
	arr := []rune(str)
	for index, ch := range arr {
		if !(ch >= 65 && ch <= 90) && !(ch >= 97 && ch <= 122) && !(ch >= '0' && ch <= '9') {
			flag = true
		} else if (ch >= 65 && ch <= 90) && flag == true {
			flag = false
		} else if (ch >= 65 && ch <= 90) && flag == false {
			arr[index] = arr[index] + 32
		} else if (ch >= 97 && ch <= 122) && flag == true {
			arr[index] = arr[index] - 32
			flag = false
		} else {
			flag = false
		}
	}
	return str
}

func Punctuation(data string) string {
	nData := ""
	for _, ch := range data {
		if ch != '.' && ch != ',' && ch != '!' && ch != '?' && ch != ':' && ch != ';' {
			nData += string(ch)
		} else {
			nData += string(ch) + " "
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

	// handling quotation
	/*
		count := 0
		for i := 0; i < len(editedData); i++ {
			for _, ch := range editedData {
				if ch == '\'' && count == 0 {
					count += 1
					nData[i+1] = string(ch) + nData[i+1]

				}updatedData
			}
		}*/

	return editedData
}

func Quotation(str string) string {
	results := ""
	inQuot := false
	for _, char := range str {
		switch char {
		case '\'':
			inQuot = true
		case '"':
			inQuot = false
		default:
			if !inQuot {
				results += string(char)
			}
		}
	}
	return results
}

func EditeA(str string) string {
	arrStr := strings.Fields(str)
	vowels := "aeiouhAEIOUH"

	for i, wrd := range arrStr {
		for _, ltr := range vowels {
			if wrd == "a" && rune(arrStr[i+1][0]) == ltr {
				arrStr[i] = "an"
			} else if wrd == "A" && rune(arrStr[i+1][0]) == ltr {
				arrStr[i] = "An"
			}
		}
	}

	result := ""
	for i := 0; i < len(arrStr); i++ {
		if i == len(arrStr)-1 {
			result += arrStr[i]
		} else {
			result += arrStr[i] + " "
		}
	}
	return result
}

// main func that recieves the data from input file and make the modifications !!
func Modifiy(data string) string {

	words := strings.Split(string(data), " ")
	for i, word := range words {
		if word == "(up)" {
			words[i-1] = ToUpperCase(words[i-1])
		} else if word == "(low)" {
			words[i-1] = ToLowerCase(words[i-1])
		} else if word == "(hex)" {
			words[i-1] = Hex2dec(words[i-1])
		} else if word == "(bin)" {
			words[i-1] = Bin2dec(words[i-1])
		} else if word == "(up," {
		/*	num := strconv.Atoi(string(words[i+1][0]))
			for i := 0; i < num; i++ {
				words[i-1] = ToUpperCase(words[i-1])
			}*/
		}
	}

	result := ""
	for i := 0; i < len(words); i++ {
		if i == len(words)-1 {
			result += words[i]
		} else {
			result += words[i] + " "
		}
	}
	result = RemoveCommands(result)
	result = EditeA(result)
	result = Punctuation(result)
	return result
}

func RemoveCommands(str string) string {
	results := ""
	inParentheses := false
	for _, char := range str {
		switch char {
		case '(':
			inParentheses = true
		case ')':
			inParentheses = false
		default:
			if !inParentheses {
				results += string(char)
			}
		}
	}
	return results
}
