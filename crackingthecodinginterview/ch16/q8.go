package main

import (
	"fmt"
	"strings"
)

var m map[int]string = map[int]string{
	1:       "one",
	2:       "two",
	3:       "three",
	4:       "four",
	5:       "five",
	6:       "six",
	7:       "seven",
	8:       "eight",
	9:       "nine",
	10:      "ten",
	11:      "eleven",
	12:      "twelve",
	13:      "thriteen",
	14:      "fourteen",
	15:      "fifteen",
	16:      "sixteen",
	17:      "seventeen",
	18:      "eighteen",
	19:      "nineteen",
	20:      "twenty",
	30:      "thirty",
	40:      "forty",
	50:      "fifty",
	60:      "sixty",
	70:      "seventy",
	80:      "eighty",
	90:      "ninty",
	100:     "hundred",
	1000:    "thousand",
	1000000: "million",
}

const (
	TEN      = 10
	HUNDRED  = 100
	THOUSAND = 1000
	MILLION  = 1000000
)

func EnglishInt(n int) string {
	if v, ok := m[n]; ok {
		return v
	}
	r := " "

	if n/MILLION != 0 {
		return EnglishInt(n/MILLION) + " " + m[MILLION] + " " + EnglishInt(n%MILLION)
	} else if n/THOUSAND != 0 {
		return EnglishInt(n/THOUSAND) + " " + m[THOUSAND] + " " + EnglishInt(n%THOUSAND)
	} else if n/HUNDRED != 0 {
		return EnglishInt(n/HUNDRED) + " " + m[HUNDRED] + " " + EnglishInt(n%HUNDRED)
	} else if n/TEN != 0 {
		q := (n / TEN) * 10
		return EnglishInt(q) + " " + EnglishInt(n%TEN)
	}

	return strings.Trim(r," ")
}

func main() {
	fmt.Println(EnglishInt(45))
	fmt.Println(EnglishInt(99))
	fmt.Println(EnglishInt(199))
	fmt.Println(EnglishInt(987))
	fmt.Println(EnglishInt(1024))
	fmt.Println(EnglishInt(112407))
	fmt.Println(EnglishInt(1124077))
}
