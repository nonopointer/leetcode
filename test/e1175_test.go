package test

import (
	"fmt"
	"testing"
)

func Test1175(t *testing.T) {
	fmt.Println(numPrimeArrangements(100))
}


const mod int = 1e9 + 7

func numPrimeArrangements(n int) int {
	Pcnt := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			Pcnt++
		}
	}
	return (getSum(Pcnt) * getSum(n-Pcnt)) % mod
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true

}

func getSum(n int) int {
	sum := 1
	for i := 1; i <= n; i++ {
		sum = sum * i % mod
	}
	return sum
}
