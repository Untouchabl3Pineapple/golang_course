/*
	Дано натуральное число, выведите его последнюю цифру.

	Формат входных данных
	На вход дается натуральное число N, не превосходящее 10000.

	Формат выходных данных
	Выведите одно целое число - ответ на задачу.

	Sample Input:
	123
	Sample Output:
	3
*/

package main

import (
	"fmt"
)

func main() {
	var (
		numb uint
		res uint
	)

	fmt.Scan(&numb)

	res = numb % 10

	fmt.Println(res)
}