/*
	Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа). 

	Формат входных данных
	На вход дается натуральное число, не превосходящее 10000.

	Формат выходных данных
	Выведите одно целое число - число десятков.

	Sample Input:
	2010
	Sample Output:
	1
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

	res = numb / 10 % 10

	fmt.Println(res)
}