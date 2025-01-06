package main

import (
	"flag"
	"log"
	"strings"
)

// isBalanced returns whether the given expression
// has balanced brackets.
const openBrackets string = "({["
const closeBrackets string = ")}]"

func isBalanced(expr string) bool {
	var stack = make([]byte, 0)
	for i := 0; i < len(expr); i++ {
		c := expr[i]
		if strings.Contains(openBrackets, string(c)) {
			stack = append(stack, c)
		} else if idx := strings.Index(closeBrackets, string(c)); idx >= 0 {
			if len(stack) > 0 && openBrackets[idx] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				log.Printf("Find a closing bracket without a corresponding open bracket %v at position %v", string(c), i)
				break
			}
		}
	}
	return len(stack) == 0
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
