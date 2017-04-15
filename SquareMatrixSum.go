package main
import (
	"fmt"
)
func main() {
	var linesNumber, number, sum int

	fmt.Println("Количество строк?")
	fmt.Scanf("%d", &linesNumber)

	list := make([][]int, linesNumber)

	for i := range list {
		list[i] = make([]int, linesNumber)
	}

	for i := range linesNumber {
		for j := range linesNumber {
			fmt.Println("Введите число")
			fmt.Scanf("%d", &number)
			list[i][j] = number
		}
	}

	for i := range linesNumber{
		sum = sum + list[i][i]
	}

	fmt.Println(sum)

}