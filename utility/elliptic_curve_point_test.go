package utility_test

import (
	"math/big"
	"testing"

	"github.com/CityOfZion/neo-go-sdk/utility"
	"github.com/stretchr/testify/assert"
)

func TestEllipticCurvePoint(t *testing.T) {
	t.Run(".Format()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			testCases := []struct {
				description string
				x           *big.Int
				y           *big.Int
				out         string
			}{
				{
					description: "NilPoints",
					x:           nil,
					y:           nil,
					out:         "(inf,inf)",
				},
				{
					description: "NonNilPoints",
					x:           big.NewInt(1),
					y:           big.NewInt(2),
					out:         "(01,02)",
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.description, func(t *testing.T) {
					point := utility.EllipticCurvePoint{
						X: testCase.x,
						Y: testCase.y,
					}
					result := point.Format()

					assert.Equal(t, testCase.out, result)
				})
			}
		})
	})
}
