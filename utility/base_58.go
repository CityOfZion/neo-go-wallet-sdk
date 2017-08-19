package utility

import (
	"fmt"
	"math/big"
)

type (
	// Base58 is a encode/decode utility for base58 strings.
	Base58 struct {
		alphabet  [58]byte
		decodeMap [256]int64
	}
)

const (
	defaultAlphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

// NewBase58 creates a new Base58 struct, using the default alphabet.
func NewBase58() Base58 {
	base58 := Base58{}

	copy(base58.alphabet[:], []byte(defaultAlphabet))

	for i := range base58.decodeMap {
		base58.decodeMap[i] = -1
	}

	for i, element := range base58.alphabet {
		base58.decodeMap[element] = int64(i)
	}

	return base58
}

// Decode decodes the base58 encoded bytes.
func (b Base58) Decode(s string) ([]byte, error) {
	source := []byte(s)

	if len(source) == 0 {
		return []byte{}, nil
	}

	var zeros []byte

	for i, element := range source {
		if element == b.alphabet[0] && i < len(source)-1 {
			zeros = append(zeros, '0')
		} else {
			break
		}
	}

	radix := big.NewInt(58)
	n := new(big.Int)
	var i int64

	for _, element := range source {
		if i = b.decodeMap[element]; i < 0 {
			return nil, fmt.Errorf(
				"invalid character '%c' in decoding a base58 string \"%s\"", element, source,
			)
		}

		n.Add(
			n.Mul(n, radix),
			big.NewInt(i),
		)
	}

	return n.Append(zeros, 10), nil
}
