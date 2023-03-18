package operations

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	for _, tc := range []struct {
		firstNumber  int
		secondNumber int
	}{
		{
			firstNumber:  23948729347934,
			secondNumber: 948209384027392,
		},
		{
			firstNumber:  92837498273,
			secondNumber: 29384092395,
		},
	} {
		firstNumber := big.NewInt(int64(tc.firstNumber))
		secondNumber := big.NewInt(int64(tc.secondNumber))

		result := firstNumber.Add(firstNumber, secondNumber)
		assert.Equal(t, result, Add(firstNumber, secondNumber), fmt.Sprintf("failed to add: %d and %d", tc.firstNumber, tc.secondNumber))
	}
}

func TestSub(t *testing.T) {
	for _, tc := range []struct {
		firstNumber  int
		secondNumber int
	}{
		{
			firstNumber:  23948729347934,
			secondNumber: 948209384027392,
		},
		{
			firstNumber:  92837498273,
			secondNumber: 29384092395,
		},
	} {
		firstNumber := big.NewInt(int64(tc.firstNumber))
		secondNumber := big.NewInt(int64(tc.secondNumber))

		result := firstNumber.Sub(firstNumber, secondNumber)
		assert.Equal(t, result, Sub(firstNumber, secondNumber), fmt.Sprintf("failed to substract: %d and %d", tc.firstNumber, tc.secondNumber))
	}
}
