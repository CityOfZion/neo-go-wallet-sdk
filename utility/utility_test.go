package utility_test

import (
	"math/big"
)

// hexToInt accepts a hex value represented as a string and returns an
// initialized big integer with the provided hex value ignoring any errors.
func hexToInt(h string) (i *big.Int) {
	i, _ = new(big.Int).SetString(h, 16)
	return
}
