package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	errNotEnough       = errors.New("not enough elements for operation")
	errZeroDivision    = errors.New("division by zero")
	errSqrtForNegative = errors.New("square root of negative number")
	errInvalidRPN      = errors.New("the user input does not form a valid RPN expression")
	errUnknownOperator = errors.New("unknown token/operator")
)

func main() {
	result, err := RPNcalculator("5 1 2 + 4 * + 3 - +") // 14

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

// Reverse Polish Notation calculator
func RPNcalculator(input string) (float64, error) {
	// Stack for operands
	operands := make([]float64, 0)

	// Slice splitted input
	splittedInput := strings.Split(input, " ")

	// Loop to add operands to the stack and check for operations when encountered
	for i := 0; i < len(splittedInput); i++ {
		operand, err := strconv.ParseFloat(splittedInput[i], 64)
		// if the input is not and operant try to do the operation and check for errors
		if err != nil {
			opErr := operationCheck(splittedInput[i], &operands)
			if opErr != nil {
				return 0, opErr
			}
			continue
		}
		operands = append(operands, operand)
	}

	// If there are more operands than operations return an error
	if len(operands) > 1 {
		fmt.Println(operands)
		return 0, errInvalidRPN
	}
	return operands[0], nil
}

func operationCheck(operator string, operands *[]float64) error {
	// Check for the number of operants that the operation will affect
	operandsAffected := operandsAffectedCheck(operator)
	operandsLen := len(*operands)
	if operandsLen < operandsAffected {
		fmt.Println(*operands, "operation:", operator)
		return errNotEnough
	}

	num1 := (*operands)[operandsLen-operandsAffected]
	num2 := (*operands)[operandsLen-1]
	result, err := operation(num1, num2, operator)
	if err != nil {
		fmt.Println(operands, "operation:", operator)
		return err
	}
	*operands = append((*operands)[:operandsLen-operandsAffected], result)
	return nil
}

// Simple check for special operations that don't affect two operands
func operandsAffectedCheck(operand string) int {
	switch operand {
	case "sqrt":
		return 1
	default:
		return 2
	}
}

// Do the operation and return the result or fail and return the error
func operation(num1 float64, num2 float64, operation string) (float64, error) {
	var result float64
	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return 0, errZeroDivision
		}
		result = num1 / num2
	case "sqrt":
		if num1 < 0 {
			return 0, errSqrtForNegative
		}
		return math.Sqrt(num1), nil
	default:
		return 0, errUnknownOperator
	}

	return result, nil
}
