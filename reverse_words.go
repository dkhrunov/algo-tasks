package main

import "strings"

// fmt.Println(reverseWords("Hello, Pette, This is nice day, right? - Yeap!"))

func reverseWords(str string) string {
	var strBuilder strings.Builder
	stack := make([]rune, 0, 2)

	isLetter := func(r rune) bool {
		return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
	}

	for _, r := range str {
		if isLetter(r) {
			stack = append(stack, r)
		} else {
			for len(stack) != 0 {
				strBuilder.WriteRune(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			strBuilder.WriteRune(r)
		}
	}

	return strBuilder.String()
}
