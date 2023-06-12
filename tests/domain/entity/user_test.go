package tests

import (
	"testing"

	"github.com/peidrao/instago/internal/domain/entity"
	"github.com/stretchr/testify/require"
)

func TestValidatePassewordSuccess(t *testing.T) {
	password := "@Teste123oa"

	user := entity.User{
		Password: password,
	}

	validate := entity.NewValidator()
	err := validate.Struct(user)

	require.NoError(t, err)
}

func TestValidatePassewordSuccesError(t *testing.T) {
	password := "12345678"

	user := entity.User{
		Password: password,
	}

	validate := entity.NewValidator()
	err := validate.Struct(user)

	require.Error(t, err)
}
