package parser

import "fmt"

var depth = 0
var spacing = 4
var isEnabled = false

func untrace(subject string) {
	if isEnabled {
		depth--
		fmt.Printf("%*s", depth*spacing, " ")
		fmt.Printf("END %s\n", subject)
	}
}

func trace(subject string) string {
	if isEnabled {
		fmt.Printf("%*s", depth*spacing, " ")
		fmt.Printf("BEGIN %s\n", subject)
		depth++
	}

	return subject
}
