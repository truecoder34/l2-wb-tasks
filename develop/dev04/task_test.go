package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain_sortLettersInWord(t *testing.T) {

	testTable := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "OK",
			input:  "ватерполистка",
			output: "аавеиклопрстт",
		},
		{
			name:   "Empty",
			input:  "",
			output: "",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			res := sortLettersInWord(testCase.input)
			assert.Equal(t, testCase.output, res)
		})
	}
}

func TestMain_sortArrayOfStrings(t *testing.T) {

	testTable := []struct {
		name   string
		input  []string
		output []string
	}{
		{
			name:   "OK",
			input:  []string{"пистолет", "арбат", "якорь"},
			output: []string{"арбат", "пистолет", "якорь"},
		},
		{
			name:   "Empty",
			input:  []string{""},
			output: []string{""},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			res := sortArrayOfStrings(testCase.input)
			assert.Equal(t, testCase.output, res)
		})
	}
}

func TestMain_removeDuplicates(t *testing.T) {
	testTable := []struct {
		name   string
		input  []string
		output []string
	}{
		{
			name:   "OK",
			input:  []string{"якорь", "кеша", "якорь", "дом"},
			output: []string{"якорь", "кеша", "дом"},
		},
		{
			name:   "Empty",
			input:  []string{""},
			output: []string{""},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			res := removeDuplicates(testCase.input)
			assert.Equal(t, testCase.output, res)
		})
	}

}

func TestMain_searchAnagrams(t *testing.T) {

	outputTrue := make(map[string][]string)
	outputTrue["листок"] = append(outputTrue["листок"], "листок")
	outputTrue["листок"] = append(outputTrue["листок"], "столик")

	testTable := []struct {
		name   string
		input  []string
		output map[string][]string
	}{
		{
			name:   "OK",
			input:  []string{"листок", "столик", "столик", "слоник", "дом"},
			output: outputTrue,
		},
		{
			name:   "Empty",
			input:  []string{""},
			output: make(map[string][]string),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			res := searchAnagrams(testCase.input)
			assert.Equal(t, testCase.output, res)
		})
	}
}
