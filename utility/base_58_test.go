package utility_test

import (
	"testing"

	"github.com/CityOfZion/neo-go-sdk/utility"
	"github.com/stretchr/testify/assert"
)

func TestBase58(t *testing.T) {
	t.Run(".Decode()", func(t *testing.T) {
		testCases := []struct {
			description string
			expected    string
			value       string
		}{
			{
				description: "Empty",
				expected:    "",
				value:       "",
			},
			{
				description: "SingleDecimal",
				expected:    "0",
				value:       "1",
			},
			{
				description: "SingleCharacter",
				expected:    "32",
				value:       "Z",
			},
			{
				description: "FullStringOne",
				expected:    "79228162514264337593543950336",
				value:       "5qCHTcgbQwpvYZQ9d",
			},
			{
				description: "FullStringTwo",
				expected:    "10398420938103833",
				value:       "2QCTtMag6t",
			},
		}

		base58 := utility.NewBase58()

		for _, testCase := range testCases {
			t.Run(testCase.description, func(t *testing.T) {
				result, err := base58.Decode(testCase.value)
				assert.Nil(t, err)
				assert.Equal(t, testCase.expected, string(result))
			})
		}
	})
}
