package main
import (
	"fmt"
)
func main() {
	var linesNumber int
	var elemsNumber int
	var number int
	var sum int

	fmt.Println("Количество строк?")
	fmt.Scanf("%d", &linesNumber)
	fmt.Println("Количество элементов в строке?")
	fmt.Scanf("%d", &elemsNumber)

	list := make([][]int, linesNumber)

	for i := range list {
		list[i] = make([]int, elemsNumber)
	}

	for i := 0; i < linesNumber; i++ {
		for j := 0; j < elemsNumber; j++ {
			fmt.Println("Введите число")
			fmt.Scanf("%d", &number)
			list[i][j] = number
		}
	}

	for i := 0; i < linesNumber; i++ {
		for j := 0; j < elemsNumber; j++ {
			if j == i {
				sum = sum + list[i][j]
			}
		}
	}


	fmt.Println(sum)

}