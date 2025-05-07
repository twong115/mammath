package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
	q "github.com/twong115/mammath/questions"
)

func main() {
	var question q.Question = q.GenerateSimplePolynomial(3)

    fmt.Println("Find the derivative of:", question.GetQuestionString())

    reader := bufio.NewReader(os.Stdin)
    answer, _ := reader.ReadString('\n')
    answer = strings.TrimSuffix(answer, "\n")

    if answer == question.GetSolutionString() {
        fmt.Println("Correct!")
    } else {
        fmt.Println("Incorrect! The answer is:", question.GetSolutionString())
    }
}
