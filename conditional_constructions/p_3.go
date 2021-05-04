/*Дано неотрицательное целое число. Найдите и выведите первую цифру числа. 

Формат входных данных
На вход дается натуральное число, не превосходящее 10000.

Формат выходных данных
Выведите одно целое число - первую цифру заданного числа.

Sample Input:
1234
Sample Output:
1*/

package main

import (
    "fmt"
)

// Циклы на данный момент еще не пройдены

func main() {
    var (
        numb int16
    )

    fmt.Scan(&numb)

    switch {
    case numb / 10 == 0:
        fmt.Println(numb)
    case numb / 100 == 0:
        fmt.Println(numb / 10)
    case numb / 1000 == 0:
        fmt.Println(numb / 100)
    case numb / 10000 == 0:
        fmt.Println(numb / 1000)
    default:
        fmt.Println("1")
    }
}