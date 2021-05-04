// Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.



// Формат входных данных

// На вход подается номер билета - одно шестизначное  число.

// Формат выходных данных
// Выведите "YES", если билет счастливый, в противном случае - "NO".

// Sample Input:
// 613244
// Sample Output:
// YES

package main

import (
    "fmt"
)

// Циклы на данный момент еще не пройдены

func main() {
    var (
        numb        int32

        first_dig   int32
        second_dig  int32
        third_dig   int32
        fourth_dig  int32
        fifth_dig   int32
        sixth_dig   int32
    )

    fmt.Scan(&numb)

    first_dig  =    numb % 10
    second_dig =    numb / 10 % 10
    third_dig  =    numb / 100 % 10
    fourth_dig =    numb / 1000 % 10
    fifth_dig  =    numb / 10000 % 10
    sixth_dig  =    numb / 100000

    switch {
    case first_dig + second_dig + third_dig == fourth_dig + fifth_dig + sixth_dig:
        fmt.Println("YES")
    default:
        fmt.Println("NO")
    }
}