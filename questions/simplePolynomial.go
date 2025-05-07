package questions

import (
    "fmt"
    "math/rand"
	"strings"
)

// SimplePolynomial is a type of Question where you solve
// for the derivative of a polynomial equation.
type SimplePolynomial struct {
	degree int
	question string
	solution string
}

func (sp SimplePolynomial) GetQuestionString() string {
	return sp.question
}

func (sp SimplePolynomial) GetSolutionString() string {
	return sp.solution
}

// Creates a randomly generated SimplePolynomial question.
func GenerateSimplePolynomial(degree int) SimplePolynomial {
	res := SimplePolynomial{degree: degree}
	res.makeQuestion()
	return res
}

func FormatEquation(poly []string) (string) {
    var question string = strings.ReplaceAll(strings.Join(poly[:], " + "), " + -", " - ")
    if question == "" {
        question = "0"
    }

    return question
}

func (sp SimplePolynomial) coef_to_string(coef, pow int) (string) {
	term := ""
    if coef == 0 {
        return term
    }

    num := fmt.Sprintf("%d", coef)

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

func (sp *SimplePolynomial) makeQuestion() {
	degree := sp.degree

	coefficients := make([]int, degree + 1)
	powers := make([]int, degree + 1)

	var poly, deri []string

	for i:= 0; i <= degree; i++ {
		powers[i] = degree - i
		coefficients[i] = rand.Intn(10)

		term := sp.coef_to_string(coefficients[i], powers[i])
		if term != "" {
			poly = append(poly, term)
		}

		term = sp.coef_to_string(coefficients[i] * powers[i], powers[i] - 1)
		if term != "" {
			deri = append(deri, term)
		}
	}

	var question string = FormatEquation(poly)
	var solution string = FormatEquation(deri)

	sp.question = question
	sp.solution = solution
} 
