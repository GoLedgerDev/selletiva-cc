package txdefs

import (
	"encoding/json"
	"time"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	tx "github.com/goledgerdev/cc-tools/transactions"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var EditarEntrada = tx.Transaction{
	Tag:         "editaEntrada",
	Label:       "Editar Entrada",
	Description: "Editar Entrada",
	Method:      "POST",
	//Callers:     []string{`org\dMSP`},

	Args: []tx.Argument{
		{
			Tag:         "entrada",
			Label:       "Entrada",
			Description: "Entrada",
			DataType:    "entradas",
			Required:    true,
		},
		{
			Tag:         "cod_entrada",
			Label:       "Código Entrada",
			Description: "Código Entrada",
			DataType:    "number",
		},
		{
			Tag:         "corporativo",
			Label:       "Corporativo",
			Description: "Corporativo",
			DataType:    "number",
		},
		{
			Tag:         "entidade",
			Label:       "Entidade Gerenciadora",
			Description: "Entidade Gerenciadora",
			DataType:    "number",
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
			Tag:         "unidade",
			Label:       "Unidade",
			Description: "Unidade",
			DataType:    "number",
		},
		{
			Tag:         "gerador",
			Label:       "Gerador",
			Description: "Gerador",
			DataType:    "number",
		},
		{
			Tag:         "acondicionamento",
			Label:       "Acondicionamento",
			Description: "Acondicionamento",
			DataType:    "number",
		},
		{
			Tag:         "mtr",
			Label:       "MTR",
			Description: "MTR",
			DataType:    "string",
		},
		{
			Tag:         "origem",
			Label:       "Origem",
			Description: "Origem",
			DataType:    "number",
		},
		{
			Tag:         "lacre",
			Label:       "Lacre",
			Description: "Lacre",
			DataType:    "string",
		},
		{
			Tag:         "lote",
			Label:       "Lote",
			Description: "Lote",
			DataType:    "string",
		},
		{
			Tag:         "data_registro",
			Label:       "Data de Registro de Entrada",
			Description: "Data de Registro de Entrada",
			DataType:    "datetime",
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
			Tag:      "categoria_nome",
			Label:    "Nome da Categoria",
			DataType: "string",
		},
		{
			Tag:      "tipo_nome",
			Label:    "Nome do Tipo",
			DataType: "string",
		},
		{
			Tag:      "status",
			Label:    "Status",
			DataType: "string",
		},
		{
			Tag:         "regra1_cliente_blockchain",
			Label:       "Regra Data Registro de Entrada Blockchain",
			Description: "Se esta regra será utilizada ou não para: Data Registro",
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
		entradaKey, ok := req["entrada"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "missing parameter entrada")
		}
		cod_entrada, hasCodEntrada := req["cod_entrada"].(float64)
		corporativo, hasCorporativo := req["corporativo"].(float64)
		entidade, hasEntidade := req["entidade"].(float64)
		categoria, hasCategoria := req["categoria"].(float64)
		tipo, hasTipo := req["tipo"].(float64)
		quantidade, hasQuantidade := req["quantidade"].(float64)
		unidade, hasUnidade := req["unidade"].(float64)
		gerador, hasGerador := req["gerador"].(float64)
		transportadora, hasTransportadora := req["transportadora"].(float64)
		acondicionamento, hasACondicionamento := req["acondicionamento"].(float64)
		mtr, hasMtr := req["mtr"].(string)
		origem, hasOrigem := req["origem"].(float64)
		lacre, hasLacre := req["lacre"].(string)
		lote, hasLote := req["lote"].(string)
		veiculo, hasVeiculo := req["veiculo"].(float64)
		motorista, hasMotorista := req["motorista"].(float64)
		data_registro, hasDataRegistro := req["data_registro"].(time.Time)
		distancia_app, hasDistanciaApp := req["distancia_app"].(float64)
		dias_evidencia_app, hasDiasEvidenciaApp := req["dias_evidencia_app"].(float64)
		categoria_nome, hasCategoriaNome := req["categoria_nome"].(string)
		tipo_nome, hasTipoNome := req["tipo_nome"].(string)
		status, hasStatus := req["status"].(string)

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

		entradaAsset, err := entradaKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get entrada from the ledger")
		}

		entrada := (map[string]interface{})(*entradaAsset)
		if entrada["@assetType"].(string) != "entradas" {
			return nil, errors.WrapError(err, "failed to get entrada from the ledger")
		}

		if regra1_cliente_blockchain {
			if !hasDataAtual {
				return nil, errors.WrapError(nil, "Verificação de regra 1 exige o campo data_atual")
			}
			if data_registro.Before(data_atual) {
				campo_saida += "Regra1:data_registro"
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
			entrada["categoria"] = categoria
		}
		if hasTipo {
			entrada["tipo"] = tipo
		}
		if hasQuantidade {
			entrada["quantidade"] = quantidade
		}
		if hasTransportadora {
			entrada["transportadora"] = transportadora
		}
		if hasDataRegistro {
			entrada["data_registro"] = data_registro.Format(time.RFC3339)
		}

		if hasDistanciaApp {
			entrada["distancia_app"] = distancia_app
		}
		if hasDiasEvidenciaApp {
			entrada["dias_evidencia_app"] = dias_evidencia_app
		}

		if hasCodEntrada {
			entrada["cod_entrada"] = cod_entrada
		}

		if hasCorporativo {
			entrada["corporativo"] = corporativo
		}
		if hasEntidade {
			entrada["entidade"] = entidade
		}
		if hasUnidade {
			entrada["unidade"] = unidade
		}

		if hasGerador {
			entrada["gerador"] = gerador
		}

		if hasACondicionamento {
			entrada["acondicionamento"] = acondicionamento
		}

		if hasMtr {
			entrada["mtr"] = mtr
		}

		if hasOrigem {
			entrada["origem"] = origem
		}

		if hasLacre {
			entrada["lacre"] = lacre
		}

		if hasLote {
			entrada["lote"] = lote
		}

		if hasVeiculo {
			entrada["veiculo"] = veiculo
		}

		if hasMotorista {
			entrada["motorista"] = motorista
		}

		if hasCategoriaNome {
			entrada["categoria_nome"] = categoria_nome
		}

		if hasTipoNome {
			entrada["tipo_nome"] = tipo_nome
		}

		if hasStatus {
			entrada["status"] = status
		}

		entrada["campo_controle"] = campo_controle
		entrada["campo_saida"] = campo_saida

		entradaJSON, nerr := json.Marshal(entrada)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		entradaEdit, err := assets.NewAsset(entrada)

		_, err = entradaEdit.Put(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		// Marshal asset back to JSON format
		entradaJSON, nerr = json.Marshal(entradaEdit)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return entradaJSON, nil
	},
}
