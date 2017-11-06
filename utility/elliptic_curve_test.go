package utility_test

import (
	"math/big"
	"testing"

	"github.com/CityOfZion/neo-go-sdk/utility"
	"github.com/stretchr/testify/assert"
)

func TestEllipticCurveAdd(t *testing.T) {
	t.Run(".Add()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			testCases := []struct {
				description string
				p           utility.EllipticCurvePoint
				q           utility.EllipticCurvePoint
				out         *utility.EllipticCurvePoint
			}{
				{
					description: "InfinitePoints",
					p:           utility.EllipticCurvePoint{X: nil, Y: nil},
					q:           utility.EllipticCurvePoint{X: nil, Y: nil},
					out:         &utility.EllipticCurvePoint{X: nil, Y: nil},
				},
				{
					description: "InfinitePoint1",
					p:           utility.EllipticCurvePoint{X: nil, Y: nil},
					q:           utility.EllipticCurvePoint{X: big.NewInt(1), Y: big.NewInt(1)},
					out:         &utility.EllipticCurvePoint{X: big.NewInt(1), Y: big.NewInt(1)},
				},
				{
					description: "InfinitePoint2",
					p:           utility.EllipticCurvePoint{X: big.NewInt(1), Y: big.NewInt(1)},
					q:           utility.EllipticCurvePoint{X: nil, Y: nil},
					out:         &utility.EllipticCurvePoint{X: big.NewInt(1), Y: big.NewInt(1)},
				},
				{
					description: "PQEquivX",
					p:           utility.EllipticCurvePoint{X: big.NewInt(5), Y: big.NewInt(2)},
					q:           utility.EllipticCurvePoint{X: big.NewInt(5), Y: big.NewInt(4)},
					out:         &utility.EllipticCurvePoint{X: nil, Y: nil},
				},
				{
					description: "PQEquivXYNeg",
					p:           utility.EllipticCurvePoint{X: big.NewInt(7), Y: big.NewInt(31)},
					q:           utility.EllipticCurvePoint{X: big.NewInt(7), Y: big.NewInt(31)},
					out:         &utility.EllipticCurvePoint{X: big.NewInt(2), Y: big.NewInt(1)},
				},
				{
					description: "PQDiffX",
					p:           utility.EllipticCurvePoint{X: big.NewInt(7), Y: big.NewInt(31)},
					q:           utility.EllipticCurvePoint{X: big.NewInt(9), Y: big.NewInt(31)},
					out:         &utility.EllipticCurvePoint{X: big.NewInt(2), Y: big.NewInt(2)},
				},
			}

			ec := utility.NewEllipticCurve()
			ec.P = big.NewInt(3)
			ec.A = big.NewInt(14)

			for _, testCase := range testCases {
				result, err := ec.Add(testCase.p, testCase.q)

				assert.NoError(t, err)
				assert.Equal(t, testCase.out, result)
			}
		})

		t.Run("SadCase", func(t *testing.T) {
			testCases := []struct {
				description string
				p           utility.EllipticCurvePoint
				q           utility.EllipticCurvePoint
			}{
				{
					description: "InvalidPointAddition",
					p:           utility.EllipticCurvePoint{X: big.NewInt(7), Y: big.NewInt(6)},
					q:           utility.EllipticCurvePoint{X: big.NewInt(7), Y: big.NewInt(4)},
				},
			}

			ec := utility.NewEllipticCurve()
			ec.P = big.NewInt(3)

			for _, testCase := range testCases {
				result, err := ec.Add(testCase.p, testCase.q)

				assert.Error(t, err)
				assert.Nil(t, result)
			}
		})
	})
}

func TestEllipticCurveIsOnCurve(t *testing.T) {
	t.Run(".IsOnCurve()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			testCases := []struct {
				description string
				p           utility.EllipticCurvePoint
				out         bool
			}{
				{
					description: "Infinite",
					p:           utility.EllipticCurvePoint{X: nil, Y: nil},
					out:         false,
				},
				{
					description: "FiniteOnPoint",
					p:           utility.EllipticCurvePoint{X: big.NewInt(15), Y: big.NewInt(30)},
					out:         true,
				},
				{
					description: "FiniteNotOnPoint",
					p:           utility.EllipticCurvePoint{X: big.NewInt(15), Y: big.NewInt(47)},
					out:         false,
				},
			}

			ec := utility.NewEllipticCurve()
			ec.P = big.NewInt(3)
			ec.A = big.NewInt(15)
			ec.B = big.NewInt(30)

			for _, testCase := range testCases {
				result := ec.IsOnCurve(testCase.p)

				assert.Equal(t, testCase.out, result)
			}
		})
	})
}

func TestEllipticCurveScalarMult(t *testing.T) {
	t.Run(".ScalarMult()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			testCases := []struct {
				description string
				k           *big.Int
				p           utility.EllipticCurvePoint
				out         *utility.EllipticCurvePoint
			}{
				{
					description: "ValidPointAddition",
					k:           big.NewInt(7),
					p:           utility.EllipticCurvePoint{X: big.NewInt(13), Y: big.NewInt(21)},
					out:         &utility.EllipticCurvePoint{X: big.NewInt(13), Y: big.NewInt(21)},
				},
			}

			ec := utility.NewEllipticCurve()
			ec.P = big.NewInt(3)
			ec.N = big.NewInt(41)
			ec.A = big.NewInt(1)

			for _, testCase := range testCases {
				result, err := ec.ScalarMult(testCase.k, testCase.p)

				assert.NoError(t, err)
				assert.Equal(t, testCase.out, result)
			}
		})
	})
}
