package utility_test

import (
	"math/big"
	"testing"

	"github.com/CityOfZion/neo-go-sdk/utility"
	"github.com/stretchr/testify/assert"
)

func TestModularArithmeticAdd(t *testing.T) {
	t.Run(".Add()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			testCases := []struct {
				description string
				X           *big.Int
				Y           *big.Int
				P           *big.Int
				out         *big.Int
			}{
				{
					description: "Zero",
					X:           big.NewInt(0),
					Y:           big.NewInt(0),
					P:           big.NewInt(1),
					out:         big.NewInt(0),
				},
				{
					description: "Nonzero",
					X:           big.NewInt(1000),
					Y:           big.NewInt(3500),
					P:           big.NewInt(34),
					out:         big.NewInt(12),
				},
			}

			ma := utility.NewModularArithmetic()

			for _, testCase := range testCases {
				result := ma.Add(testCase.X, testCase.Y, testCase.P)

				assert.Equal(t, testCase.out, result)
			}
		})
	})
}

func TestModularArithmeticExp(t *testing.T) {
	t.Run(".Exp()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			testCases := []struct {
				description string
				X           *big.Int
				Y           *big.Int
				P           *big.Int
				out         *big.Int
			}{
				{
					description: "Zero",
					X:           big.NewInt(0),
					Y:           big.NewInt(0),
					P:           big.NewInt(1),
					out:         big.NewInt(0),
				},
				{
					description: "Nonzero",
					X:           big.NewInt(1000),
					Y:           big.NewInt(3500),
					P:           big.NewInt(34),
					out:         big.NewInt(4),
				},
			}

			ma := utility.NewModularArithmetic()

			for _, testCase := range testCases {
				result := ma.Exp(testCase.X, testCase.Y, testCase.P)

				assert.Equal(t, testCase.out, result)
			}
		})
	})
}

func TestModularArithmeticInverse(t *testing.T) {
	t.Run(".Inverse()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			testCases := []struct {
				description string
				X           *big.Int
				P           *big.Int
				out         *big.Int
			}{
				{
					description: "Zero",
					X:           big.NewInt(0),
					P:           big.NewInt(0),
					out:         big.NewInt(0),
				},
				{
					description: "Nonzero",
					X:           big.NewInt(1000),
					P:           big.NewInt(3500),
					out:         big.NewInt(3497),
				},
			}

			ma := utility.NewModularArithmetic()

			for _, testCase := range testCases {
				result := ma.Inverse(testCase.X, testCase.P)

				assert.Equal(t, testCase.out, result)
			}
		})
	})
}

func TestModularArithmeticMul(t *testing.T) {
	t.Run(".Mul()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			testCases := []struct {
				description string
				X           *big.Int
				Y           *big.Int
				P           *big.Int
				out         *big.Int
			}{
				{
					description: "Zero",
					X:           big.NewInt(0),
					Y:           big.NewInt(0),
					P:           big.NewInt(1),
					out:         big.NewInt(0),
				},
				{
					description: "Nonzero",
					X:           big.NewInt(1000),
					Y:           big.NewInt(3500),
					P:           big.NewInt(34),
					out:         big.NewInt(6),
				},
			}

			ma := utility.NewModularArithmetic()

			for _, testCase := range testCases {
				result := ma.Mul(testCase.X, testCase.Y, testCase.P)

				assert.Equal(t, testCase.out, result)
			}
		})
	})
}

func TestModularArithmeticSqrt(t *testing.T) {
	t.Run(".Sqrt()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			testCases := []struct {
				description string
				X           *big.Int
				P           *big.Int
				out         *big.Int
			}{
				{
					description: "Nonzero",
					X:           big.NewInt(1000),
					P:           big.NewInt(3503),
					out:         big.NewInt(2488),
				},
			}

			ma := utility.NewModularArithmetic()

			for _, testCase := range testCases {
				result := ma.Sqrt(testCase.X, testCase.P)

				assert.Equal(t, testCase.out, result)
			}
		})
	})
}

func TestModularArithmeticSub(t *testing.T) {
	t.Run(".Sub()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			testCases := []struct {
				description string
				X           *big.Int
				Y           *big.Int
				P           *big.Int
				out         *big.Int
			}{
				{
					description: "Zero",
					X:           big.NewInt(0),
					Y:           big.NewInt(0),
					P:           big.NewInt(1),
					out:         big.NewInt(0),
				},
				{
					description: "Nonzero",
					X:           big.NewInt(1000),
					Y:           big.NewInt(3500),
					P:           big.NewInt(34),
					out:         big.NewInt(16),
				},
			}

			ma := utility.NewModularArithmetic()

			for _, testCase := range testCases {
				result := ma.Sub(testCase.X, testCase.Y, testCase.P)

				assert.Equal(t, testCase.out, result)
			}
		})
	})
}
