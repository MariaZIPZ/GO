package handlers

import (
	"errors"
	"fmt"
	"lab7/packages/templates"
	"net/http"
	"strconv"
	"strings"
)

func Slice(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, templates.PageStart, templates.BodyStart, templates.SliceTemplate)

	var inputStr string
	var slice []float64
	var err error

	if r.Method == http.MethodGet {
		inputStr = r.FormValue("slice")

		if inputStr == "" {
			return
		}

		fmt.Fprint(w, "<p>Calculated using a Get request:</p>")
	}

	if r.Method == http.MethodPost {
		err = r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, templates.ErrorTag, err.Error())
			return
		}

		post := r.PostForm
		inputStr = post.Get("slice")

		if inputStr == "" {
			return
		}

		fmt.Fprint(w, "<p>Calculated using a Post request:</p>")
	}

	slice, err = readSlice(inputStr)

	if err != nil {
		fmt.Fprintf(w, templates.ErrorTag, err.Error())
		return
	}

	multiplication := calculateMultiplicationEvenItems(slice)
	sum := calculateSumBetweenFirstZeroItemAndLastZeroItem(slice)

	responseStr := "<p></p><p><big>Results: </big></p>"
	responseStr += fmt.Sprintf("<p><i>Multiplication of elements with even indices</i>: <b>%.2f</b></p>", multiplication)
	responseStr += fmt.Sprintf("<p><i>Sum of elements between the first and last zero elements</i>: <b>%.2f</b></p>", sum)
	fmt.Fprint(w, responseStr)

	fmt.Fprint(w, templates.BodyEnd)
	fmt.Fprint(w, templates.PageEnd)
}

func readSlice(sliceStr string) ([]float64, error) {
	if sliceStr == "" {
		return nil, errors.New("empty input string")
	}

	elementsStr := strings.Split(sliceStr, " ")
	var result []float64

	for _, element := range elementsStr {
		number, err := strconv.ParseFloat(element, 64)

		if err != nil {
			return nil, fmt.Errorf("unable to convert '%s' to float64: %w", element, err)
		}

		result = append(result, number)
	}

	return result, nil
}

func calculateMultiplicationEvenItems(slice []float64) float64 {
	if len(slice) <= 1 {
		return 0
	}

	result := 1.0
	for i := 0; i < len(slice); i += 2 {
		result *= slice[i]
	}

	return result
}

func calculateSumBetweenFirstZeroItemAndLastZeroItem(slice []float64) float64 {
	countZeroItems, firstZeroItemId, lastZeroItemId := findIndexesZeroItems(slice)

	if countZeroItems <= 1 {
		return 0
	}

	result := 0.0
	for i := firstZeroItemId + 1; i < lastZeroItemId; i++ {
		result += slice[i]
	}

	return result
}

func findIndexesZeroItems(slice []float64) (int, int, int) {
	firstZeroItemId, lastZeroItemId := -1, -1

	for i, val := range slice {
		if val == 0 {
			if firstZeroItemId == -1 {
				firstZeroItemId = i
			} else {
				lastZeroItemId = i
			}
		}
	}

	if lastZeroItemId == -1 {
		if firstZeroItemId == -1 {
			return 0, firstZeroItemId, lastZeroItemId
		}
		return 1, firstZeroItemId, lastZeroItemId
	}

	return 2, firstZeroItemId, lastZeroItemId
}
