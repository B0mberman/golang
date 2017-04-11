package main

import "fmt"

func main() {
	var n int
	var x int
	fmt.Println("Количество имен?")
	fmt.Scanf("%d", &n)

	list := make([]string, 0)
	result := make([]string, 0)

	var name string

	for i:= 0; i < n; i++ {
		fmt.Println("Введите имя")
		fmt.Scanf("%s", &name)
		list = append(list, name)
	}

	result = append(result, list[0])

	for i:= 0; i < len(list); i++ {
		for j:= 0; j < len(result); j ++ {
			if result[j] != list[i]{
				x++
			}
			if x == len(result) {
				result = append(result, list[i])
				x = 0
			}
			if result[j] == list[i] {
				x = 0
				break
			}
		}

	}

	fmt.Println(result)
}