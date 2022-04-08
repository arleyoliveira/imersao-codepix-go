package model_test

import (
	"github.com/arleyoliveira/imersao-codepix-go/domain/model"
	"testing"
)

func TestModel_NewPixKey(t *testing.T) {
	code := "380"
	name := "PicPay Serviços S.A"
	bank, _ := model.NewBank(code, name)

	accountNumber := "123456"
	ownerName := "Arley Oliveira"
	_, _ = model.NewAccount(bank, accountNumber, ownerName)
}
