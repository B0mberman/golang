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

	for i := 0; i < linesNumber; i++ {
		for j := 0; j < linesNumber; j++ {
			fmt.Println("Введите число")
			fmt.Scanf("%d", &number)
			list[i][j] = number
		}
	}

	for i := 0; i < linesNumber; i++ {
		sum = sum + list[i][i]
	}

	fmt.Println(sum)

}