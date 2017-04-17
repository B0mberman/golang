package main
import (
	"fmt"
)
func main() {
	var linesNumber, elemsNumber, number, minx, maxx, miny, maxy, y, x, k int

	fmt.Println("Количество строк?")
	fmt.Scanf("%d", &linesNumber)
	fmt.Println("Количество элементов в строке?")
	fmt.Scanf("%d", &elemsNumber)

	maxx = elemsNumber-1
	maxy = linesNumber-1
	k = -1

	list := make([]int, linesNumber*elemsNumber)
	result := make([][]int, linesNumber)

	for i := range result {
		result[i] = make([]int, elemsNumber)
	}

	for i := 0; i < linesNumber*elemsNumber; i++ {
		fmt.Println("Введите число")
		fmt.Scanf("%d", &number)
		list[i] = number
	}

	for maxx != x && maxy != y {

		for x = minx; x <= maxx; x++ {
			k++
			result[y][x] = list[k]

		}
		miny++
		x--

		for y = miny; y <= maxy; y++ {
			k++
			result[y][x] = list[k]
		}
		y--
		maxx--

		for x = maxx; x >= minx; x-- {
			k++
			result[y][x] = list[k]
		}
		x++
		maxy--

		for y = maxy; y >= miny; y-- {
			k++
			result[y][x] = list[k]
		}
		y++
		minx++
	}

	for i := 0; i < linesNumber; i++ {
		for j := 0; j < elemsNumber; j++ {
			fmt.Print(result[i][j])
		}
		fmt.Println()
	}

}