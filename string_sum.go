package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.TrimSpace(input)
	if input == "" {
		err = fmt.Errorf("Bad string %w", errorEmptyInput)
		return "", err
	}
	input = strings.Replace(input, " ", "", -1)

	isFirstNegative := strings.HasPrefix(input, "-")

	if isFirstNegative {
		input = input[1:]
	}

	opsPositive := strings.Split(input, "+")
	opsNegative := strings.Split(input, "-")

	if len(opsNegative) != 2 && len(opsPositive) != 2 {
		err = fmt.Errorf("Bad string %w", errorNotTwoOperands)
		return "", err
	}

	var opFirstStr, opSecStr string
	if len(opsNegative) == 2 {
		opFirstStr = opsNegative[0]
		opSecStr = "-" + opsNegative[1]
	 } else {
		opFirstStr = opsPositive[0]
		opSecStr = opsPositive[1]
	 }

	 if (isFirstNegative) {
		opFirstStr = "-" + opFirstStr 
	 }

	var opFirst, opSec int 

	opFirst, err = strconv.Atoi(opFirstStr)
	if err != nil {
		err = fmt.Errorf("First operand is not right %w", err)
		return "", err
	}

	opSec, err = strconv.Atoi(opSecStr)
	if err != nil {
		err = fmt.Errorf("Second operand is not right %w", err)
		return "", err
	}

	sum := opFirst + opSec
	output = strconv.Itoa(sum)

	return output, nil
}

// func main() {
// 	str := "-3-5"
// 	fmt.Println(str)
// 	o, err := StringSum(str)
// 	if err == nil {
// 		fmt.Println(o)
// 	}
// 	fmt.Println(err)
// }