package main

import (
	txdefs "github.com/goledgerdev/selletiva-cc/chaincode/txdefs"

	tx "github.com/goledgerdev/cc-tools/transactions"
)

var txList = []tx.Transaction{
	txdefs.GetHeader,
	txdefs.EditarSolicitacao,
	txdefs.CriarSolicitacao,
	txdefs.CriarSaida,
}
