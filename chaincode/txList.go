package main

import (
	txdefs "github.com/goledgerdev/template-cc/chaincode/txdefs"

	tx "github.com/goledgerdev/cc-tools/transactions"
)

var txList = []tx.Transaction{
	txdefs.GetHeader,
	txdefs.EditarSolicitacao,
	txdefs.CriarSolicitacao,
	txdefs.CriarSaida,
	txdefs.EditarSaida,
	txdefs.CriarEntrada,
	txdefs.EditarEntrada,
}
