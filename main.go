package main

import (
    "fmt"
    "math/rand"
    "strings"
    "bufio"
    "os"
)

func coef_to_string(coef, pow int) (string) {
    var term string
    if coef != 0 {
        if pow == 1 {
            term = fmt.Sprintf("%vx", coef)
        } else if pow == 0 {
            term = fmt.Sprintf("%v", coef)
        } else {
            term = fmt.Sprintf("%vx^%v", coef, pow)
        }
    }

    return term
}

func main() {
    var degree int = 3

    coefficients := make([]int, degree + 1)
    powers := make([]int, degree + 1)

    var poly, deri []string

    // Get polynomial and derivative
    for i:= 0; i <= degree; i++ {
        powers[i] = degree - i
        coefficients[i] = rand.Intn(10)
        
        term := coef_to_string(coefficients[i], powers[i])
        if term != "" {
            poly = append(poly, term)
        }

        term = coef_to_string(coefficients[i] * powers[i], powers[i] - 1)
        if term != "" {
            deri = append(deri, term)
        }
    }

    fmt.Println("Find the derivative of:", strings.Join(poly[:], " + "))

    reader := bufio.NewReader(os.Stdin)
    answer, _ := reader.ReadString('\n')
    answer = strings.TrimSuffix(answer, "\n")

    var solution string = strings.Join(deri[:], " + ")

    if answer == solution {
        fmt.Println("Correct!")
    } else {
        fmt.Println("Incorrect! The answer is:", solution)
    }
}
