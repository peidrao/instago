package tests

import (
	"testing"

	"github.com/peidrao/instago/utils"
	"github.com/stretchr/testify/require"
)

func TestRandomStrings(t *testing.T) {

	t.Run("Should create a string of 5 characters", func(t *testing.T) {
		randomString := utils.GenerateRandomString(5)

		require.NotEmpty(t, randomString)
		require.Equal(t, len(randomString), 5)
	})

	t.Run("Should create a string of 10 characters", func(t *testing.T) {
		randomString := utils.GenerateRandomString(10)

		require.NotEmpty(t, randomString)
		require.Equal(t, len(randomString), 10)
	})

	t.Run("Should create a string of 50 characters", func(t *testing.T) {
		randomString := utils.GenerateRandomString(50)

		require.NotEmpty(t, randomString)
		require.Equal(t, len(randomString), 50)
	})

}
