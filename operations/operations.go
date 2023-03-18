package operations

import "math/big"

func Add(firstNumber, secondNumber *big.Int) *big.Int {
	return firstNumber.Add(firstNumber, secondNumber)
}

func Sub(firstNumber, secondNumber *big.Int) *big.Int {
	return firstNumber.Sub(firstNumber, secondNumber)
}
