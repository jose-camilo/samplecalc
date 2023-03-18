package operations

import (
	"math/big"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	firstNumber := big.NewInt(int64(384928340928))
	secondNumber := big.NewInt(int64(3948503845))
	for n := 0; n < b.N; n++ {
		Add(firstNumber, secondNumber)
	}
}

func BenchmarkSub(b *testing.B) {
	firstNumber := big.NewInt(int64(384928340928))
	secondNumber := big.NewInt(int64(3948503845))
	for n := 0; n < b.N; n++ {
		Sub(firstNumber, secondNumber)
	}
}
