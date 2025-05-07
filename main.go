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
    if coef == 0 {
        return term
    }

    var num string = fmt.Sprintf("%v", coef)

    if (coef == 1 || coef == -1) && pow != 0 {
        num = strings.ReplaceAll(num, "1", "")
    }

    if pow == 1 {
        term = fmt.Sprintf("%vx", num)
    } else if pow == 0 {
        term = fmt.Sprintf("%v", num)
    } else {
        term = fmt.Sprintf("%vx^%v", num, pow)
    }

    return term
}

func get_question_string(poly []string) (string) {
    var question string = strings.ReplaceAll(strings.Join(poly[:], " + "), " + -", " - ")
    if question == "" {
        question = "0"
    }

    return question
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

    var question string = get_question_string(poly)

    fmt.Println("Find the derivative of:", question)

    reader := bufio.NewReader(os.Stdin)
    answer, _ := reader.ReadString('\n')
    answer = strings.TrimSuffix(answer, "\n")

    var solution string = get_question_string(deri)

    if answer == solution {
        fmt.Println("Correct!")
    } else {
        fmt.Println("Incorrect! The answer is:", solution)
    }
}
