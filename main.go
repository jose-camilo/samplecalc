package main

import (
	"fmt"
	"github.com/jose-camilo/simplecalc/operations"
	"math/big"
	"os"
)

const (
	AddOperation = "add"
	SubOperation = "sub"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Operations should be formulated in this way: <operation> <firstNumber> <secondNumber>")
		return
	}

	var result *big.Int
	switch os.Args[1] {
	case AddOperation:
		result = operations.Add(scanStringToBig(os.Args[2]), scanStringToBig(os.Args[3]))
	case SubOperation:
		result = operations.Sub(scanStringToBig(os.Args[2]), scanStringToBig(os.Args[3]))
	default:
		fmt.Printf("\nWrong operation: the only options available are %q and %q", "add", "sub")
		return
	}

	text, err := result.MarshalText()
	if err != nil {
		fmt.Println("Error marshaling response")
		return
	}

	fmt.Printf("\nThe result is: %s\n", text)
}
