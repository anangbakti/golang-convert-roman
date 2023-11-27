package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var mappingRomanNum = make(map[string]int)

func main() {
	// Mapping collection of Roman numeral key-value pairs to numbers.
	mappingRomanNum = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	file, err := os.Open("roman.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var readRoman string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		readRoman = scanner.Text()

		fmt.Println(readRoman)
		if !validRoman(readRoman) {
			continue
		}
		numRoman := convertToNumber(readRoman)
		fmt.Println(strconv.Itoa(numRoman))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func convertToNumber(roman string) int {
	var result int

	// The result value is taken from the last letter.
	lengthOfString := len(roman)
	lastElement := roman[len(roman)-1 : lengthOfString]
	result = mappingRomanNum[lastElement]

	// Read backward
	for i := len(roman) - 1; i > 0; i-- {
		//  If the last numeral <= the numeral before it, value result = result + the numeral before it.
		if mappingRomanNum[roman[i:i+1]] <= mappingRomanNum[roman[i-1:i]] {
			result += mappingRomanNum[roman[i-1:i]]
		} else {
			// Else, If the last numeral > the numeral before it, value result = result - the numeral before it.
			result -= mappingRomanNum[roman[i-1:i]]
		}
	}
	return result
}

func charInArray(char rune, charArray []rune) bool {
	for _, c := range charArray {
		if c == char {
			return true
		}
	}
	return false
}

func charInString(charArray []rune, str string) bool {
	for _, c := range str {
		if charInArray(c, charArray) {
			return true
		}
	}
	return false
}

func dlvOnlyOnce(roman string) bool {
	dlv := []rune{'D', 'L', 'V'}
	counts := make(map[rune]int)
	for _, char := range roman {
		if counts[char] == 1 && charInString(dlv, string(char)) {
			fmt.Printf("'%c' is a duplicate character.\n", char)
			fmt.Print("'D', 'L', or 'V' may only appear at most once.\n")
			return false
		}
		counts[char]++ // increment the count
	}
	return true
}

func mcxiNotAppearThreeTimesInARow(roman string) bool {
	mcxi := []rune{'M', 'C', 'X', 'I'}
	maxConsecutive := 3
	consecutiveCount := 1

	for i := 1; i < len(roman); i++ {
		if roman[i] == roman[i-1] && charInString(mcxi, string(roman[i])) {
			consecutiveCount++
			if consecutiveCount > maxConsecutive {
				fmt.Printf("'%c' Consecutive characters duplicated.\n", roman[i])
				fmt.Print("'M', 'C', 'X', or 'I' may appear no more than three times in a row. \n")
				return false
			}
		} else {
			consecutiveCount = 1
		}
	}

	return true
}

func invalidPairOfLettersValueSubtractedCantBeMore10xSubtractionLetter(roman string) bool {
	invalidPair := []string{"IL", "IC", "ID", "IM", "XD", "XM"}
	for _, pair := range invalidPair {
		if strings.Contains(roman, pair) {
			fmt.Printf("'%s' Invalid pair of letter.\n", pair)
			fmt.Print("When subtracting, the value of the letter being subtracted from cannot be more than \n")
			fmt.Print("10 times the value of letter being used for subtraction. \n")
			return false
		}
	}

	return true
}

func invalidPairOfLettersOnlyIXCCanBeUsedForSubstraction(roman string) bool {
	invalidPair := []string{"VX", "VL", "VC", "VD", "VM", "LC", "LD", "LM", "DM"}
	for _, pair := range invalidPair {
		if strings.Contains(roman, pair) {
			fmt.Printf("'%s' Invalid pair of letter.\n", pair)
			fmt.Print("Only I, X, and C can be used for subtraction (V, L, and D cannot). \n")
			return false
		}
	}

	return true
}

func NotIllegalSequence(roman string) bool {
	return doStrAFollowStrB(roman, "C", "CD") &&
		doStrAFollowStrB(roman, "C", "CM") &&
		doStrAFollowStrB(roman, "X", "XL") &&
		doStrAFollowStrB(roman, "X", "XC") &&
		doStrAFollowStrB(roman, "I", "IV") &&
		doStrAFollowStrB(roman, "I", "IX") &&
		doStrAFollowStrB(roman, "V", "IX") &&
		doStrAFollowStrB(roman, "L", "XC") &&
		doStrAFollowStrB(roman, "D", "CM")
}

func doStrAFollowStrB(roman string, strA string, strB string) bool {
	if strings.Contains(roman, strA+strB) {
		fmt.Printf("'%s' can not follow '%s'.\n", strA, strB)
		return false
	}
	return true
}

// Roman to number rules from : https://www.thevbprogrammer.com/Ch08/08-10-RomanNumerals.htm
func validRoman(roman string) bool {
	return dlvOnlyOnce(roman) && mcxiNotAppearThreeTimesInARow(roman) &&
		invalidPairOfLettersValueSubtractedCantBeMore10xSubtractionLetter(roman) &&
		invalidPairOfLettersOnlyIXCCanBeUsedForSubstraction(roman) &&
		NotIllegalSequence(roman)
}
