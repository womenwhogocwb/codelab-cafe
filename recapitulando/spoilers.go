package main

import "fmt"

func main() {
	// 1
	meuInt := 27
	// 2
	fmt.Printf("1 e 2) valor: %d; tipo: %T\n", meuInt, meuInt)
	// 3
	meuOutroInt := 200
	// 4
	somaDosInts := meuInt + meuOutroInt
	fmt.Printf("3 e 4) soma dos inteiros: %d\n", somaDosInts)
	// 5
	ehPar := somaDosInts%2 == 0
	fmt.Println("5) a soma dos inteiros é par?", ehPar)
	// 6
	if ehPar {
		fmt.Println("6) >> a soma é par!")
	} else {
		fmt.Println("6) >> a soma é ímpar!")
	}
	// 7
	meuSlice := []int{5, 7}
	fmt.Printf("7) o valor do slice é: %+v\n", meuSlice)
	// 8
	somaDosElementos := meuSlice[0] + meuSlice[1]
	fmt.Printf("8) a soma dos dois elementos é: %d\n", somaDosElementos)
	// 9
	meuSlice = append(meuSlice, 13, 21, 27, 31)
	fmt.Printf("9) o novo valor do slice é: %+v\n", meuSlice)
	// 10
	soma := 0
	for _, elemento := range meuSlice {
		soma += elemento
	}
	fmt.Printf("10) a soma de todos os elementos do slice é: %d\n", soma)
	// 12
	fmt.Printf("11 e 12) a soma performada pela função é: %d\n", SomaSlices(meuSlice))
}

// 11
func SomaSlices(slice []int) int {
	soma := 0
	for _, elemento := range slice {
		soma += elemento
	}
	return soma
}