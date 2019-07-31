package main

import (
	"fmt"
)

// func letterCombinations(digits string) []string {
// 	var result []string
// 	tmp := map[string][]string{
// 		"2": []string{"a", "b", "c"},
// 		"3": []string{"d", "e", "f"},
// 		"4": []string{"g", "h", "i"},
// 		"5": []string{"j", "k", "l"},
// 		"6": []string{"m", "n", "o"},
// 		"7": []string{"p", "q", "r", "s"},
// 		"8": []string{"t", "u", "v"},
// 		"9": []string{"w", "x", "y", "z"},
// 	}

// 	rs := strings.Split(digits, "")
// 	if len(rs) == 1 {
// 		result = tmp[digits]
// 	}
// 	if len(rs) > 1 {
// 		first := tmp[rs[0]]
// 		for _, x := range first {
// 			for _, y := range letterCombinations(strings.Join(rs[1:], "")) {
// 				result = append(result, fmt.Sprintf("%s%s", x, y))
// 			}
// 		}
// 	}

// 	return result
// }

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	} else if len(digits) == 1 {
		return lettersForDigit(digits[0])
	}
	letters := lettersForDigit(digits[0])
	// fmt.Println(letters)
	rest := letterCombinations(digits[1:])
	res := make([]string, len(letters)*len(rest))
	for i := 0; i < len(letters); i++ {
		for j := 0; j < len(rest); j++ {
			res[i*len(rest)+j] = letters[i] + rest[j]

		}
	}
	return res
}

func lettersForDigit(digit byte) []string {
	var res []string
	if digit <= '6' {
		first := (digit-'2')*3 + 'a'
		res = []string{string(first), string(first + 1), string(first + 2)}
	} else if digit == '7' {
		res = []string{"p", "q", "r", "s"}
	} else if digit == '8' {
		res = []string{"t", "u", "v"}
	} else if digit == '9' {
		res = []string{"w", "x", "y", "z"}
	} else {
		res = []string{" "}
	}
	return res
}

func main() {
	fmt.Println(letterCombinations("54235"))
	// a := []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(a[1:])
}
