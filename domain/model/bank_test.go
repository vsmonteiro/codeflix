package model_test

import (
	"testing"

	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/vsmonteiro/codepix/domain/model"
)

func TestModel_NewBank(t *testing.T) {

	code := "001"
	name := "Banco do Brasil"

	bank, err := model.NewBank(code, name)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.MustParse(bank.ID))
	require.Equal(t, bank.Code, code)
	require.Equal(t, bank.Name, name)

	_, err = model.NewBank("", "")
	require.NotNil(t, err)
}
