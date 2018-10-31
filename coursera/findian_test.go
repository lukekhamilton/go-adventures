package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

var testDataResults = map[string]string{
	"ian":              "Found!",
	"Ian":              "Found!",
	"iuiygaygn":        "Found!",
	"I d skd a efju N": "Found!",
	"ihhhhhn":          "Not Found!",
	"ina":              "Not Found!",
	"xian":             "Not Found!",
}

func TestMain(t *testing.T) {
	// fmt.Printf("%+v\n", testDataResults)

	for input, expectedResult := range testDataResults {
		fmt.Printf("Test Data Input:'%+v' expectedResult:'%+v'\n", input, expectedResult)

		result := ProcessINPUT(bufio.NewReader(strings.NewReader(input)))

		if result != expectedResult {
			t.Errorf("Result: %v not expected for input: %v", result, input)
		}

	}
}
