package tests

import (
	"testing"

	"github.com/peidrao/instago/utils"
	"github.com/stretchr/testify/require"
)

func TestGenerateToken(t *testing.T) {

	t.Run("Should create a token", func(t *testing.T) {
		token, err := utils.GenerateToken("test")

		require.NoError(t, err)
		require.NotEmpty(t, token)
	})

	t.Run("Should extract bearer token", func(t *testing.T) {
		token, err := utils.GenerateToken("test")

		require.NoError(t, err)
		require.NotEmpty(t, token)

		headerValue := "Bearer " + token

		require.Contains(t, headerValue, "Bearer")

		tokenResult := utils.ExtractBearerToken(headerValue)

		require.NotEmpty(t, tokenResult)
		require.Equal(t, tokenResult, token)

	})
}
