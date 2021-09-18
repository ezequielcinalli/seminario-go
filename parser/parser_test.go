package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEstructura(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TXJ6ABCDE", false, "", "", 0},
		{"NNF", false, "", "", 0},
	}

	for _, testData := range cases {
		r, err := GetEstructura(testData.Input)

		assert.Equal(t, err == nil, testData.Success, "err == nil")
		if err == nil {
			assert.Equal(t, r.Type, testData.Type, "Type")
			assert.Equal(t, r.Value, testData.Value, "Value")
			assert.Equal(t, r.Length, testData.Length, "Length")
		}
	}
}

// func TestGetEstructura(t *testing.T) {
// 	estructura, err := GetEstructura("TX03ABC")

// 	assert.Equal(t, estructura.Type, "TX")
// 	assert.Equal(t, estructura.Length, 3)
// 	assert.Equal(t, estructura.Value, "ABC")
// 	assert.Equal(t, err == nil, true)
// }
