/*По данному трехзначному числу определите, все ли его цифры различны.

Формат входных данных
На вход подается одно натуральное трехзначное число.

Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".

Sample Input 1:
237
Sample Output 1:
YES
Sample Input 2:
117
Sample Output 2:
NO*/

package main

import (
    "fmt"
)

func main() {
    var (
        numb int16

        first_dig int16
        second_dig int16
        third_dig int16
    )

    fmt.Scan(&numb)

    first_dig = numb % 10
    second_dig = numb / 10 % 10
    third_dig = numb / 100

    switch {
    case first_dig == second_dig || first_dig == third_dig || second_dig == third_dig:
        fmt.Println("NO")
    default:
        fmt.Println("YES")
    }
}