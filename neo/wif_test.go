package neo_test

import (
	"testing"

	"github.com/CityOfZion/neo-go-sdk/neo"
	"github.com/stretchr/testify/assert"
)

func TestWIF(t *testing.T) {
	t.Run(".ToPrivateKey()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			wif := neo.NewWIF("5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ")

			foo, err := wif.ToPrivateKey()
			assert.Nil(t, err)
			assert.Equal(t, "WIP", foo)
		})
	})
}
