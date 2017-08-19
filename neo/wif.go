package neo

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
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

	fmt.Printf("Decoded (%d):\t%s\n", len(decoded), hex.EncodeToString(decoded))

	// Remove last 4 bytes from decoded and encode to hex
	shortened := hex.EncodeToString(decoded[:len(decoded)-4])
	fmt.Printf("Shortened (%d):\t%s\n", len(shortened), shortened)

	// SHA-256 the shortened hex string, result is bytes
	firstHasher := sha1.New()
	firstHasher.Write([]byte(shortened))
	firstSHA256 := firstHasher.Sum(nil)
	fmt.Printf("1st 256 (%d):\t%s\n", len(firstSHA256), hex.EncodeToString(firstSHA256))

	// SHA-256 it again, result is bytes
	secondHasher := sha1.New()
	secondHasher.Write(firstSHA256)
	secondSHA256 := secondHasher.Sum(nil)
	fmt.Printf("2nd 256 (%d):\t%s\n", len(secondSHA256), hex.EncodeToString(secondSHA256))

	shaResult := []byte(hex.EncodeToString(secondSHA256))

	// Take first 4 bytes of 2nd SHA-256 result
	firstFourBytes := shaResult[:4]
	fmt.Printf("1st 4 (%d):\t%s\n", len(firstFourBytes), hex.EncodeToString(firstFourBytes))

	// Take last 4 bytes of original decoded value (part that wasn't included in shortened value)
	lastFourBytes := decoded[len(decoded)-4 : len(decoded)]
	fmt.Printf("Last 4 (%d):\t%s\n", len(lastFourBytes), hex.EncodeToString(lastFourBytes))

	fmt.Println(firstFourBytes)
	fmt.Println(lastFourBytes)

	for i, x := range firstFourBytes {
		if x != lastFourBytes[i] {
			return "", errors.New("foo")
		}
	}

	return "WIP", nil
}
