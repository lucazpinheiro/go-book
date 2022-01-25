package main

import (
	"fmt"
	"strings"
)

var table = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func convertStringToIntSlice(rawSample string) []int {
	var mappedSample []int
	for _, char := range strings.Split(rawSample, "") {
		digit := table[char]
		mappedSample = append(mappedSample, digit)
	}
	return mappedSample
}

func verifierDigitIsValid(sample []int, digit, multiplier int) bool {
	sum := 0
	expected := 0
	for _, number := range sample {
		sum += number * multiplier
		multiplier--
	}
	rest := sum % 11
	if rest > 1 {
		expected = 11 - rest
	}
	return expected == digit
}

func main() {
	cpf := "44368564960"
	// cpf := "07328681063" // valido
	// cpf := "11122244467" // não valido
	convertedSample := convertStringToIntSlice(cpf)
	verifierDigits := convertedSample[9:]
	firstVerifierDigit := verifierDigitIsValid(convertedSample[:9], verifierDigits[0], 10)
	secondVerifierDigit := verifierDigitIsValid(convertedSample[:10], verifierDigits[1], 11)

	if firstVerifierDigit && secondVerifierDigit {
		fmt.Println("cpf é valido")
	} else {
		fmt.Println("cpf não é valido")
	}
}
