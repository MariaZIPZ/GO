package handlers

import (
	"fmt"
	"lab7/packages/templates"
	"math"
	"net/http"
	"strconv"
)

func QuadraticEquation(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, templates.PageStart, templates.BodyStart, templates.QuadraticEquationTemplate)

	var a, b, c float64
	var err error

	if r.Method == http.MethodPost {
		err = r.ParseForm()

		if err != nil {
			fmt.Fprintf(w, templates.ErrorTag, err.Error())
			return
		}

		post := r.PostForm

		aStr := post.Get("a")
		bStr := post.Get("b")
		cStr := post.Get("c")

		if aStr == "" && bStr == "" && cStr == "" {
			return
		}

		a, err = strconv.ParseFloat(aStr, 64)
		b, err = strconv.ParseFloat(bStr, 64)
		c, err = strconv.ParseFloat(cStr, 64)
		if err != nil {
			fmt.Fprintf(w, templates.ErrorTag, err.Error())
			return
		}

		fmt.Fprint(w, "<p>Calculated using a Post request:</p>")
	}

	if r.Method == http.MethodGet {
		aString := r.FormValue("a")
		bString := r.FormValue("b")
		cString := r.FormValue("c")

		if aString == "" && bString == "" && cString == "" {
			return
		}

		a, err = strconv.ParseFloat(aString, 64)
		b, err = strconv.ParseFloat(bString, 64)
		c, err = strconv.ParseFloat(cString, 64)

		if err != nil {
			fmt.Fprintf(w, templates.ErrorTag, err.Error())
			return
		}

		fmt.Fprint(w, "<p>Calculated using a Get request:</p>")
	}

	resultStr := makeResultString(a, b, c)
	fmt.Fprintf(w, resultStr)

	fmt.Fprint(w, templates.BodyEnd)
	fmt.Fprint(w, templates.PageEnd)
}

func makeResultString(a, b, c float64) string {
	var result string

	roots, x1, x2 := calculate(a, b, c)
	if roots == 2 {
		result = fmt.Sprintf("<p><b>Result:</b> <i>x1</i> = <b>%.3f</b> <i>x2</i> = <b>%.3f</b></p>", x1, x2)
	} else if roots == 1 {
		result = fmt.Sprintf("<p><b>Result:</b> <i>x</i> = <b>%.3f</b>", x1)
	} else {
		result = "<p><b>Result:</b> No solutions (D < 0)</p>"
	}

	return result
}

func calculate(a, b, c float64) (int, float64, float64) {
	discriminant := b*b - 4*a*c

	var roots int
	var firstRoot, secondRoot float64

	if discriminant > 0 {
		roots = 2

		firstRoot = (-b + math.Sqrt(discriminant)) / (2 * a)
		secondRoot = (-b - math.Sqrt(discriminant)) / (2 * a)
	} else if discriminant == 0 {
		roots = 1

		firstRoot = -b / (2 * a)
	}

	return roots, firstRoot, secondRoot
}
