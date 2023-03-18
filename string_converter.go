package main

import (
	"fmt"
	"log"
	"math/big"
)

func scanStringToBig(addend string) *big.Int {
	i := new(big.Int)
	_, err := fmt.Sscan(addend, i)
	if err != nil {
		log.Println("error scanning value:", err)
	}

	return i
}
