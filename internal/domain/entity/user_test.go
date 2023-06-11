package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidatePassewordSuccess(t *testing.T) {
	password := "@Teste123oa"

	user := User{
		Password: password,
	}

	validate := NewValidator()
	err := validate.Struct(user)

	require.NoError(t, err)
}

func TestValidatePassewordSuccesError(t *testing.T) {
	password := "12345678"

	user := User{
		Password: password,
	}

	validate := NewValidator()
	err := validate.Struct(user)

	require.Error(t, err)
}
