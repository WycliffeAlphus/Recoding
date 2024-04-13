package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	Modify()
}

func Modify() {
	args := os.Args[1:]
	filename, _ := os.ReadFile(args[0])
	wordsInText := strings.Split(string(filename), " ")
	for i := 0; i < (len(wordsInText)); i++ {
		if wordsInText[i] == "(hex)" {
			wordsInText[i-1] = hextodecimal(wordsInText[i-1])
			wordsInText = removePattern(wordsInText, i)
			i--
		}
		if wordsInText[i] == "(bin)" {
			wordsInText[i-1] = binarytodecimal(wordsInText[i-1])
			wordsInText = removePattern(wordsInText, i)
			i--
		}
		if wordsInText[i] == "(low)" {
			wordsInText[i-1] = strings.ToLower(wordsInText[i-1])
			wordsInText = removePattern(wordsInText, i)
			i--
		}
		if wordsInText[i] == "(up)" {
			wordsInText[i-1] = strings.ToUpper(wordsInText[i-1])
			wordsInText = removePattern(wordsInText, i)
			i--
		}
		if wordsInText[i] == "(cap)" {
			wordsInText[i-1] = strings.Title(wordsInText[i-1])
			wordsInText = removePattern(wordsInText, i)
			i--
		}
		if wordsInText[i] == "(up," {
			digit := strings.Trim(wordsInText[i+1], wordsInText[i+1][1:])
			newdigit, _ := strconv.Atoi(digit)
			for j := 1; j <= newdigit; j++ {
				wordsInText[i-j] = strings.ToUpper(wordsInText[i-j])
			}
		}
		if wordsInText[i] == "(low," {
			digit := strings.Trim(wordsInText[i+1], wordsInText[i+1][1:])
			newdigit, _ := strconv.Atoi(digit)
			for j := 1; j <= newdigit; j++ {
				wordsInText[i-j] = strings.ToLower(wordsInText[i-j])
			}
		}
		if wordsInText[i] == "(cap," {
			digit := strings.Trim(wordsInText[i+1], wordsInText[i+1][1:])
			newdigit, _ := strconv.Atoi(digit)
			for j := 1; j <= newdigit; j++ {
				wordsInText[i-j] = strings.Title(wordsInText[i-j])
			}
		}
	}
	re := regexp.MustCompile(`\(\w+,\s\d+\)`) // removing pattern incase of number specification
	output := strings.Join(wordsInText, " ")
	Newoutput := re.ReplaceAllString(output, "")
	Newoutput = removeSpaceBefore(Newoutput)
	Newoutput = mergeSymbols(Newoutput)
	Newoutput = handlesinglequotes(Newoutput)
	Newoutput2 := vowelh(strings.Fields(Newoutput))
	Newoutput = strings.Join(Newoutput2, " ")
	os.WriteFile("result.txt", []byte(Newoutput), 0o644)
}

// hextodecimal
func hextodecimal(hex string) string {
	decimal, _ := strconv.ParseInt(hex, 16, 64)
	return fmt.Sprint(decimal)
}

// bin to hex
func binarytodecimal(bin string) string {
	decimal, _ := strconv.ParseInt(bin, 2, 64)
	return fmt.Sprint(decimal)
}

// removing the pattern
func removePattern(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

// remove space before punctuation
func removeSpaceBefore(s string) string {
	re := regexp.MustCompile(`(\s+)([.,!?:;])`)
	return re.ReplaceAllString(s, "$2 ")
}

// mergimg symbols
func mergeSymbols(s string) string {
	re := regexp.MustCompile(`([.,!?:;])(\s+)([.,!?:;])`)
	return re.ReplaceAllString(s, "$1$3")
}

// func for single quotes
func handlesinglequotes(s string) string {
	words := strings.Fields(s)
	output := ""
	isStart := true
	for _, w := range words {
		if w == "'" {
			if isStart {
				output += w
				isStart = false
			} else {
				output += w + " "
				isStart = true
			}
		} else {
			output += w + " "
		}
	}
	words = strings.Fields(output)
	output = ""
	for i, w := range words {
		if i == 0 {
			output += w
		} else if w == "'" {
			output += w
		} else {
			output += " " + w
		}
	}
	return output
}

// bool for vowel
func isVowelh(s string) bool {
	vowelsh := []uint8{'a', 'e', 'i', 'o', 'u', 'h'}
	for _, letter := range vowelsh {
		if s[0] == letter {
			return true
		}
	}
	return false
}

// handling "a" and "an"
func vowelh(s []string) []string {
	l := len(s)
	for i, word := range s {
		j := i + 1
		if strings.ToLower(word) == "a" && j < l && isVowelh(s[j]) {
			s[i] = word + "n"
		}
	}
	return s
}
