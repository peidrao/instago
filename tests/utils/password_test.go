package tests

import (
	"testing"

	"github.com/peidrao/instago/utils"
	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := utils.GenerateRandomString(8)

	t.Run("Should create a hash from a string", func(t *testing.T) {

		hashedPassword, err := utils.HashPassword(password)
		require.NoError(t, err)
		require.NotEmpty(t, hashedPassword)
		require.NotEqual(t, hashedPassword, password)
	})

	t.Run("Should compare a string with hashed password", func(t *testing.T) {
		hashedPassword, _ := utils.HashPassword(password)

		err := utils.ComparePassword([]byte(hashedPassword), []byte(password))

		require.NoError(t, err)
	})
}
