package main

import (
	"strings"
)

var priorityTable = map[string]int{"+": 1, "-": 1, "*": 2, "/": 2, "^": 3}

// GetPostfix - This function will be used to convert infix expression to postfix expression
// It accepts one argument a string which is infix. This string will be converted to postfix
// It returns the postfix conversion of the input string
func GetPostfix(str string) string {

	// topOperatorInStack := ""
	operatorStack := StackS{tos: -1}
	var popedEle *string
	postStr := []string{}
	// fmt.Print(str)
	// scan each an every character and form the postfix string
	for _, char := range str {
		// fmt.Println(operatorStack.arr, operatorStack.tos)
		// fmt.Print(string(char))
		// if the char is empty line then pass
		// if string(char) == "\n" {
		// 	continue
		// }
		// if it is a operand then add it to the postfix string
		// fmt.Println(string(char))
		if _, isOperator := priorityTable[string(char)]; !isOperator && string(char) != "(" && string(char) != ")" {
			postStr = append(postStr, string(char))
			// fmt.Println(postStr, string(char))
			// if scanned char is "(" push it to stack
		} else if string(char) == "(" {
			operatorStack.Push(string(char))

			// if scanned char is ")" pop all elements till "("
		} else if string(char) == ")" {
			popedEle = operatorStack.Pop()
			// fmt.Println(*popedEle)
			for *popedEle != "(" {
				postStr = append(postStr, *popedEle)
				popedEle = operatorStack.Pop()
			}
			// fmt.Println(string(char))
		} else {
			// fmt.Println(string(char))
			// If the program reaches this part means its a operator with priority

			// pop element till the priority of scanned element is greater than tos of element
			for !operatorStack.IsEmpty() && priorityTable[string(char)] <= priorityTable[operatorStack.arr[operatorStack.tos]] {
				popedEle = operatorStack.Pop()
				// fmt.Println("aa", *popedEle)
				postStr = append(postStr, *popedEle)
			}
			operatorStack.Push(string(char))
		}
		// fmt.Print(postStr)
	}
	// fmt.Println("a", operatorStack.arr, operatorStack.tos)

	for !operatorStack.IsEmpty() {
		popedEle = operatorStack.Pop()
		// fmt.Println("p", *popedEle)
		postStr = append(postStr, *popedEle)
		// fmt.Println("pa", postStr)
	}
	return strings.Join(postStr, "")
}
