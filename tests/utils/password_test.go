package tests

import (
	"testing"

	"github.com/peidrao/instago/utils"
	"github.com/stretchr/testify/require"
)

func TestPasswordHashed(t *testing.T) {
	password := utils.GenerateRandomString(8)

	hashedPassword, err := utils.HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
}
