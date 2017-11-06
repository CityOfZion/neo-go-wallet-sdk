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
				x           *big.Int
				y           *big.Int
				p           *big.Int
				out         *big.Int
			}{
				{
					description: "Zero",
					x:           hexToInt("0000000000000000000000000000000000000000000000000000000000000000"),
					y:           hexToInt("0000000000000000000000000000000000000000000000000000000000000000"),
					p:           hexToInt("0000000000000000000000000000000000000000000000000000000000000001"),
					out:         hexToInt("0000000000000000000000000000000000000000000000000000000000000000"),
				},
				{
					description: "Nonzero",
					x:           hexToInt("50863ad64a87ae8a2fe83c1af1a8403cb53f53e486d8511dad8a04887e5b2352"),
					y:           hexToInt("2cd470243453a299fa9e77237716103abc11a1df38855ed6f2ee187e9c582ba6"),
					p:           hexToInt("5dbcd5dfea550eb4fd3b5333f533f086bb5267c776e2a1a9d8e84c16a6743d82"),
					out:         hexToInt("1f9dd51a9486426f2d4b600a738a5ff0b5fe8dfc487b0e4ac78fd0f0743f1176"),
				},
			}

			ma := utility.NewModularArithmetic()

			for _, testCase := range testCases {
				result := ma.Add(testCase.x, testCase.y, testCase.p)

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
				x           *big.Int
				y           *big.Int
				p           *big.Int
				out         *big.Int
			}{
				{
					description: "Zero",
					x:           hexToInt("0000000000000000000000000000000000000000000000000000000000000000"),
					y:           hexToInt("0000000000000000000000000000000000000000000000000000000000000000"),
					p:           hexToInt("0000000000000000000000000000000000000000000000000000000000000001"),
					out:         hexToInt("0000000000000000000000000000000000000000000000000000000000000000"),
				},
				{
					description: "Nonzero",
					x:           hexToInt("50863ad64a87ae8a2fe83c1af1a8403cb53f53e486d8511dad8a04887e5b2352"),
					y:           hexToInt("2cd470243453a299fa9e77237716103abc11a1df38855ed6f2ee187e9c582ba6"),
					p:           hexToInt("5dbcd5dfea550eb4fd3b5333f533f086bb5267c776e2a1a9d8e84c16a6743d82"),
					out:         hexToInt("02fafb8697960637249e54975cd64bc49ffb4a47a0fcdcbc375081bda8a75cf14"),
				},
			}

			ma := utility.NewModularArithmetic()

			for _, testCase := range testCases {
				result := ma.Exp(testCase.x, testCase.y, testCase.p)

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
				x           *big.Int
				p           *big.Int
				out         *big.Int
			}{
				{
					description: "Zero",
					x:           hexToInt("0000000000000000000000000000000000000000000000000000000000000000"),
					p:           hexToInt("0000000000000000000000000000000000000000000000000000000000000000"),
					out:         hexToInt("0000000000000000000000000000000000000000000000000000000000000000"),
				},
				{
					description: "Nonzero",
					x:           hexToInt("50863ad64a87ae8a2fe83c1af1a8403cb53f53e486d8511dad8a04887e5b2352"),
					p:           hexToInt("5dbcd5dfea550eb4fd3b5333f533f086bb5267c776e2a1a9d8e84c16a6743d82"),
					out:         hexToInt("2ddbccee5484a3a2d082b6abe1b6d2c6622bbb15f1497aa7f4cd1371230a4ae"),
				},
			}

			ma := utility.NewModularArithmetic()

			for _, testCase := range testCases {
				result := ma.Inverse(testCase.x, testCase.p)

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
				x           *big.Int
				y           *big.Int
				p           *big.Int
				out         *big.Int
			}{
				{
					description: "Zero",
					x:           hexToInt("0000000000000000000000000000000000000000000000000000000000000000"),
					y:           hexToInt("0000000000000000000000000000000000000000000000000000000000000000"),
					p:           hexToInt("0000000000000000000000000000000000000000000000000000000000000001"),
					out:         hexToInt("0000000000000000000000000000000000000000000000000000000000000000"),
				},
				{
					description: "Nonzero",
					x:           hexToInt("50863ad64a87ae8a2fe83c1af1a8403cb53f53e486d8511dad8a04887e5b2352"),
					y:           hexToInt("2cd470243453a299fa9e77237716103abc11a1df38855ed6f2ee187e9c582ba6"),
					p:           hexToInt("5dbcd5dfea550eb4fd3b5333f533f086bb5267c776e2a1a9d8e84c16a6743d82"),
					out:         hexToInt("b3ca04921b3d21cb028a1cf9d714b3985353e8ce23c195013aa91ed7eb3bde2"),
				},
			}

			ma := utility.NewModularArithmetic()

			for _, testCase := range testCases {
				result := ma.Mul(testCase.x, testCase.y, testCase.p)

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
				x           *big.Int
				p           *big.Int
				out         *big.Int
			}{
				{
					description: "Nonzero",
					x:           hexToInt("3e8"),
					p:           hexToInt("daf"),
					out:         hexToInt("9b8"),
				},
			}

			ma := utility.NewModularArithmetic()

			for _, testCase := range testCases {
				result := ma.Sqrt(testCase.x, testCase.p)

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
				x           *big.Int
				y           *big.Int
				p           *big.Int
				out         *big.Int
			}{
				{
					description: "Zero",
					x:           hexToInt("0000000000000000000000000000000000000000000000000000000000000000"),
					y:           hexToInt("0000000000000000000000000000000000000000000000000000000000000000"),
					p:           hexToInt("0000000000000000000000000000000000000000000000000000000000000001"),
					out:         hexToInt("0000000000000000000000000000000000000000000000000000000000000000"),
				},
				{
					description: "Nonzero",
					x:           hexToInt("50863ad64a87ae8a2fe83c1af1a8403cb53f53e486d8511dad8a04887e5b2352"),
					y:           hexToInt("2cd470243453a299fa9e77237716103abc11a1df38855ed6f2ee187e9c582ba6"),
					p:           hexToInt("5dbcd5dfea550eb4fd3b5333f533f086bb5267c776e2a1a9d8e84c16a6743d82"),
					out:         hexToInt("23b1cab216340bf03549c4f77a923001f92db2054e52f246ba9bec09e202f7ac"),
				},
			}

			ma := utility.NewModularArithmetic()

			for _, testCase := range testCases {
				result := ma.Sub(testCase.x, testCase.y, testCase.p)

				assert.Equal(t, testCase.out, result)
			}
		})
	})
}
