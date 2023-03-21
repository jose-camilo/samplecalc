package operations

import "math/big"

// Add will do the addition of two big integer numbers
func Add(firstNumber, secondNumber *big.Int) *big.Int {
	return firstNumber.Add(firstNumber, secondNumber)
}

// Sub will subtract two big integer numbers
func Sub(firstNumber, secondNumber *big.Int) *big.Int {
	return firstNumber.Sub(firstNumber, secondNumber)
}
