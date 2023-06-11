package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPasswordHashed(t *testing.T) {
	password := GenerateRandomString(8)

	hashedPassword, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
}
