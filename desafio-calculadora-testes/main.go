package main

import "fmt"


func main(){
    soma := sum(1, 1)
    subtracao := subtract(10, 1)
    multiplica := multiply(2, 5)
    divisao := divide(10, 2)
    fmt.Println(soma, subtracao, multiplica, divisao)
}


func sum(i ...int) int {
    total := 0
    for _, j := range i {
        total += j
    }
    return total
}

func subtract(total int, i ...int) int {
    for _, j := range i {
        total -= j
    }
    return total
}

func multiply(i ...int) int {
    total := 1
    for _, j := range i {
        total *= j
    }
    return total
}

func divide(total int, i ...int) int {
    for _, j := range i {
        total /= j
    }
    return total
}
