/*
	Петя торопился в школу и неправильно написал программу, которая сначала находит квадраты двух чисел, а затем их суммирует. Исправьте его программу.

	Sample Input:
	2
	2
	Sample Output:
	8
*/

package main

import (
	"fmt"
)

func main() {
    var (
		numb_1 int
		numb_2 int
		res int
	)

	fmt.Scan(&numb_1)
	fmt.Scan(&numb_2)
    
    res = numb_1 * numb_1 + numb_2 * numb_2
    
    fmt.Println(res)
}