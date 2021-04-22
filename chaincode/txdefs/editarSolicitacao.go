package txdefs

import (
	"encoding/json"
	"time"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	tx "github.com/goledgerdev/cc-tools/transactions"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var EditarSolicitacao = tx.Transaction{
	Tag:         "editarSolicitacao",
	Label:       "Editar Solicitacao",
	Description: "",
	Method:      "POST",
	//Callers:     []string{`org\dMSP`},

	Args: []tx.Argument{
		{
			Tag:         "solicitacao",
			Label:       "Solicitacao",
			Description: "Solicitação",
			DataType:    "solicitacao",
			Required:    true,
		},
		{
			Tag:         "categoria",
			Label:       "Categoria",
			Description: "Categoria",
			DataType:    "number",
		},
		{
			Tag:         "tipo",
			Label:       "Tipo",
			Description: "Tipo",
			DataType:    "number",
		},
		{
			Tag:         "quantidade",
			Label:       "Quantidade",
			Description: "Quantidade",
			DataType:    "number",
		},
		{
			Tag:         "transportadora",
			Label:       "Transportadora",
			Description: "Transportadora",
			DataType:    "number",
		},
		{
			Tag:         "observacao",
			Label:       "Observacao",
			Description: "Observacao",
			DataType:    "string",
		},
		{
			Tag:         "dta_solicita",
			Label:       "Data Solicitacao",
			Description: "Data Solicitacao",
			DataType:    "datetime",
		},
		{
			Tag:         "peso_real",
			Label:       "Peso Real",
			Description: "Peso Real",
			DataType:    "number",
		},
		{
			Tag:         "dta_recebimento",
			Label:       "Data Recebimento",
			Description: "Data Recebimento",
			DataType:    "datetime",
		},
		{
			Tag:         "status",
			Label:       "Status",
			Description: "Status",
			DataType:    "string",
		},
		{
			Tag:         "distancia_app",
			Label:       "Distancia",
			Description: "Distancia",
			DataType:    "number",
		},
		{
			Tag:         "dias_evidencia_app",
			Label:       "Dias da Evidencia",
			Description: "Dias da Evidencia",
			DataType:    "number",
		},
		{
			Tag:         "regra1_cliente_blockchain",
			Label:       "Regra Data Solicitação Blockchain",
			Description: "Se esta regra será utilizada ou não para: Data Solicitação",
			DataType:    "boolean",
			Required:    true,
		},
		{
			Tag:         "regra2_cliente_blockchain",
			Label:       "Regra Quantidade Blockchain",
			Description: "Se esta regra será utilizada ou não para: Quantidade",
			DataType:    "boolean",
			Required:    true,
		},
		{
			Tag:         "regra3_cliente_blockchain",
			Label:       "Regra Evidência App Blockchain",
			Description: "Se esta regra será utilizada ou não para: Evidência do App",
			DataType:    "boolean",
			Required:    true,
		},
		{
			Tag:         "regra4_cliente_blockchain",
			Label:       "Regra Distância App Blockchain",
			Description: "Se esta regra será utilizada ou não para: Distância App",
			DataType:    "boolean",
			Required:    true,
		},
		{
			Tag:         "regra5_cliente_blockchain",
			Label:       "Regra Dias App Blockchain",
			Description: "Se esta regra será utilizada ou não para: Dias App",
			DataType:    "boolean",
			Required:    true,
		},
		{
			Tag:         "regra6_cliente_blockchain",
			Label:       "Regra Sinir Blockchain",
			Description: "Se esta regra será utilizada ou não para: Sinir",
			DataType:    "boolean",
			Required:    true,
		},
		{
			Tag:         "data_atual",
			Label:       "Verifica Data Atual",
			Description: "Compara data de verificação com a data atual (regra 1)",
			DataType:    "datetime",
		},
		{
			Tag:         "quantidade_limite",
			Label:       "Verifica Limite de Quantidade",
			Description: "Verifica limite de quantidade",
			DataType:    "number",
		},
		{
			Tag:         "evidencia_app",
			Label:       "Verifica Evidencia App",
			Description: "Verifica se possui evidencia app",
			DataType:    "boolean",
		},
		{
			Tag:         "distancia_limite",
			Label:       "Verifica Diferença de Dias",
			Description: "Verifica diferença de dias",
			DataType:    "number",
		},
		{
			Tag:         "dias_evidencia_limite",
			Label:       "Verifica Diferença de Dias",
			Description: "Verifica diferença de dias",
			DataType:    "number",
		},
		{
			Tag:         "sinir",
			Label:       "Verifica Sinir",
			Description: "Verifica Sinir",
			DataType:    "boolean",
		},
	},
	Routine: func(stub shim.ChaincodeStubInterface, req map[string]interface{}) ([]byte, errors.ICCError) {
		solicitacaoKey, ok := req["solicitacao"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "missing parameter solicitacao")
		}
		categoria, hasCategoria := req["categoria"].(float64)
		tipo, hasTipo := req["tipo"].(float64)
		quantidade, hasQuantidade := req["quantidade"].(float64)
		transportadora, hasTransportadora := req["transportadora"].(float64)
		observacao, hasObservacao := req["observacao"].(string)
		dta_solicita, hasDataSolicitacao := req["dta_solicita"].(time.Time)
		peso_real, hasPesoReal := req["peso_real"].(float64)
		dta_recebimento, hasDataRecebimento := req["dta_recebimento"].(time.Time)
		status, hasStatus := req["status"].(string)
		distancia_app, hasDistanciaApp := req["distancia_app"].(float64)
		dias_evidencia_app, hasDiasEvidenciaApp := req["dias_evidencia_app"].(float64)

		regra1_cliente_blockchain, ok := req["regra1_cliente_blockchain"].(bool)
		if !ok {
			return nil, errors.WrapError(nil, "missing argument regra1_cliente_blockchain")
		}

		regra2_cliente_blockchain, ok := req["regra2_cliente_blockchain"].(bool)
		if !ok {
			return nil, errors.WrapError(nil, "missing argument regra2_cliente_blockchain")
		}

		regra3_cliente_blockchain, ok := req["regra3_cliente_blockchain"].(bool)
		if !ok {
			return nil, errors.WrapError(nil, "missing argument regra3_cliente_blockchain")
		}

		regra4_cliente_blockchain, ok := req["regra4_cliente_blockchain"].(bool)
		if !ok {
			return nil, errors.WrapError(nil, "missing argument regra4_cliente_blockchain")
		}

		regra5_cliente_blockchain, ok := req["regra5_cliente_blockchain"].(bool)
		if !ok {
			return nil, errors.WrapError(nil, "missing argument regra5_cliente_blockchain")
		}

		regra6_cliente_blockchain, ok := req["regra6_cliente_blockchain"].(bool)
		if !ok {
			return nil, errors.WrapError(nil, "missing argument regra6_cliente_blockchain")
		}

		data_atual, hasDataAtual := req["data_atual"].(time.Time)
		evidencia_app, hasEvidenciaApp := req["evidencia_app"].(bool)
		quantidade_limite, hasQuantidadeLimite := req["quantidade_limite"].(float64)
		distancia_limite, hasDistanciaLimite := req["distancia_limite"].(float64)
		dias_evidencia_limite, hasDiasEvidenciaLimite := req["dias_evidencia_limite"].(float64)
		sinir, hasSinir := req["sinir"].(bool)

		campo_saida := ""
		campo_controle := false

		solicitacaoAsset, err := solicitacaoKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get solicitacao from the ledger")
		}
		solicitacao := (map[string]interface{})(*solicitacaoAsset)
		if solicitacao["@assetType"].(string) != "solicitacao" {
			return nil, errors.WrapError(err, "failed to get solicitacao from the ledger")
		}

		if regra1_cliente_blockchain {
			if !hasDataAtual {
				return nil, errors.WrapError(nil, "Verificação de regra 1 exige o campo data_atual")
			}
			if dta_solicita.Before(data_atual) {
				campo_saida += "Regra1:dta_solicitacao "
				campo_controle = true
			}
		}

		if regra2_cliente_blockchain {
			if !hasQuantidadeLimite {
				return nil, errors.WrapError(nil, "Verificação de regra 2 exige o campo quantidadeLimite")
			}

			if quantidade > quantidade_limite {
				campo_saida += " Regra2:quantidade "
				campo_controle = true
			}
		}

		if regra3_cliente_blockchain {
			if !hasEvidenciaApp {
				return nil, errors.WrapError(nil, "Verificação de regra 3 exige o campo evidencia_app")
			}

			if !evidencia_app {
				campo_saida += " Regra3:evidencia "
				campo_controle = true
			}
		}

		if regra4_cliente_blockchain {
			if !hasDistanciaLimite {
				return nil, errors.WrapError(nil, "Verificação de regra 4 exige o campo distancia_limite")
			}

			if !hasDistanciaApp {
				return nil, errors.WrapError(nil, "Verificação de regra 4 exige o campo distancia_app")
			}

			if distancia_app > distancia_limite {
				campo_saida += " Regra4:distancia "
				campo_controle = true
			}
		}

		if regra5_cliente_blockchain {
			if !hasDiasEvidenciaLimite {
				return nil, errors.WrapError(nil, "Verificação de regra 5 exige o campo dias_evidencia_limite")
			}

			if !hasDiasEvidenciaApp {
				return nil, errors.WrapError(nil, "Verificação de regra 5 exige o campo dias_evidencia_app")
			}

			if dias_evidencia_app > dias_evidencia_limite {
				campo_saida += " Regra5:dias "
				campo_controle = true
			}
		}

		if regra6_cliente_blockchain {
			if !hasSinir {
				return nil, errors.WrapError(nil, "Verificação de regra 6 exige o campo sinir")
			}
			if sinir == false {
				campo_saida += " Regra6:sinir "
				campo_controle = true
			}
		}

		if hasCategoria {
			solicitacao["categoria"] = categoria
		}
		if hasTipo {
			solicitacao["tipo"] = tipo
		}
		if hasQuantidade {
			solicitacao["quantidade"] = quantidade
		}
		if hasTransportadora {
			solicitacao["transportadora"] = transportadora
		}
		if hasDataSolicitacao {
			solicitacao["dta_solicita"] = dta_solicita.Format(time.RFC3339)
		}
		if hasPesoReal {
			solicitacao["peso_real"] = peso_real
		}
		if hasObservacao {
			solicitacao["observacao"] = observacao
		}
		if hasDataRecebimento {
			solicitacao["dta_recebimento"] = dta_recebimento.Format(time.RFC3339)
		}
		if hasStatus {
			solicitacao["status"] = status
		}
		if hasDistanciaApp {
			solicitacao["distancia_app"] = distancia_app
		}
		if hasDiasEvidenciaApp {
			solicitacao["dias_evidencia_app"] = dias_evidencia_app
		}

		solicitacao["campo_controle"] = campo_controle
		solicitacao["campo_saida"] = campo_saida

		solicitacaoJSON, nerr := json.Marshal(solicitacao)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		solicitacaoEdit, err := assets.NewAsset(solicitacao)

		_, err = solicitacaoAsset.Put(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving vote on blockchain")
		}

		// Marshal asset back to JSON format
		solicitacaoJSON, nerr = json.Marshal(solicitacaoEdit)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return solicitacaoJSON, nil
	},
}
