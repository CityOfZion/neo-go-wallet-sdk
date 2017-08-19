package neo

import (
	"fmt"

	"github.com/CityOfZion/neo-go-sdk/utility"
)

type (
	// WIF (wallet import format) is a struct that holds a string, which has a number of
	// utility functions on it.
	WIF struct {
		Value string
	}
)

// NewWIF creates a WIF struct.
func NewWIF(s string) WIF {
	return WIF{
		Value: s,
	}
}

// ToPrivateKey converts the WIF to private key format.
func (w WIF) ToPrivateKey() (string, error) {
	base58 := utility.NewBase58()

	decoded, err := base58.Decode(w.Value)
	if err != nil {
		return "", err
	}

	if len(decoded) != 38 {
		return "", fmt.Errorf(
			"Expected length of decoded WIF to be 38, got: %d", len(decoded),
		)
	}

	if decoded[0] != 0x80 {
		return "", fmt.Errorf(
			"Expected first byte of decoded WIF to be '0x80', got: %x", decoded[0],
		)
	}

	if decoded[33] != 0x01 {
		return "", fmt.Errorf(
			"Expected 34th byte of decoded WIF to be '0x01', got: %x", decoded[33],
		)
	}

	return "WIP", nil
}
