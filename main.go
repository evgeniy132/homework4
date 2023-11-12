package main

import (
	"fmt"
	"strings"
)

var text string
var textRedactor []string = []string{"seafood", "food", "c", "central"}

func Search(textRedactor []string, search string) []string {
	var Complete []string
	for _, v := range textRedactor {
		if strings.Contains(v, search) {
			Complete = append(Complete, v)
		}

	}

	return Complete
}

func main() {

	fmt.Println("enter the text which need to find:")
	fmt.Scanf("%s\n", &text)
	Found := Search(textRedactor, text)
	fmt.Println(Found)
}
