package txdefs

import (
	"encoding/json"
	"time"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	tx "github.com/goledgerdev/cc-tools/transactions"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var CriarSaida = tx.Transaction{
	Tag:         "criarSaida",
	Label:       "Criar Saida",
	Description: "",
	Method:      "POST",
	//Callers:     []string{`org\dMSP`},

	Args: []tx.Argument{
		{
			Tag:         "id",
			Label:       "ID",
			Description: "id",
			DataType:    "number",
			Required:    true,
		},
		{
			Tag:         "cod_saida",
			Label:       "Código Saída",
			Description: "Código Saída",
			Required:    true,
			DataType:    "number",
		},
		{
			Tag:         "corporativo",
			Label:       "Corporativo",
			Description: "Corporativo",
			Required:    true,
			DataType:    "number",
		},
		{
			Tag:         "entidade",
			Label:       "Entidade Gerenciadora",
			Description: "Entidade Gerenciadora",
			Required:    true,
			DataType:    "number",
		},
		{
			Tag:         "categoria",
			Label:       "Categoria",
			Description: "Categoria",
			DataType:    "number",
			Required:    true,
		},
		{
			Tag:         "tipo",
			Label:       "Tipo",
			Description: "Tipo",
			DataType:    "number",
			Required:    true,
		},
		{
			Tag:         "quantidade",
			Label:       "Quantidade",
			Description: "Quantidade",
			DataType:    "number",
			Required:    true,
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
			Tag:         "destinatario",
			Label:       "Destinatario",
			Description: "Destinatario",
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
			Tag:         "veiculo",
			Label:       "Veiculo",
			Description: "Veiculo",
			DataType:    "number",
		},
		{
			Tag:         "motorista",
			Label:       "Motorista",
			Description: "Motorista",
			DataType:    "number",
		},
		{
			Tag:         "data_registro",
			Label:       "Data de Registro de Saída",
			Description: "Data de Registro de Saída",
			DataType:    "datetime",
			Required:    true,
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
			Label:       "Regra Data Registro de Saída Blockchain",
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
		id, ok := req["id"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "missing parameter id")
		}
		cod_saida, ok := req["cod_saida"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "missing parameter")
		}
		corporativo, ok := req["corporativo"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "missing parameter")
		}
		entidade, ok := req["entidade"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "missing parameter")
		}
		categoria, ok := req["categoria"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "missing parameter")
		}

		tipo, ok := req["tipo"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "missing  argument tipo")
		}

		quantidade, ok := req["quantidade"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "missing parameter quantidade")
		}
		unidade, ok := req["unidade"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "missing parameter unidade")
		}

		gerador, ok := req["gerador"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "missing parameter gerador")
		}

		transportadora, hasTransportadora := req["transportadora"].(float64)

		destinatario, ok := req["destinatario"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "missing parameter destinatario")
		}

		acondicionamento, ok := req["acondicionamento"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "missing parameter acondicionamento")
		}

		mtr, ok := req["mtr"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "missing parameter mtr")
		}

		origem, ok := req["origem"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "missing parameter origem")
		}

		lacre, ok := req["lacre"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "missing parameter lacre")
		}

		lote, ok := req["lote"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "missing parameter quantidade")
		}

		veiculo, ok := req["veiculo"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "missing parameter lote")
		}

		motorista, ok := req["motorista"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "missing parameter motorista")
		}

		data_registro, ok := req["data_registro"].(time.Time)
		if !ok {
			return nil, errors.WrapError(nil, "missing argument data de registro de saída")
		}

		categoria_nome, ok := req["categoria_nome"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "missing argument categoria_nome")
		}
		tipo_nome, ok := req["tipo_nome"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "missing argument tipo_nome")
		}
		status, ok := req["status"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "missing argument status")
		}

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

		saida := make(map[string]interface{})
		saida["@assetType"] = "saidas"
		saida["id"] = id
		saida["categoria"] = categoria
		saida["tipo"] = tipo
		saida["quantidade"] = quantidade
		if hasTransportadora {
			saida["transportadora"] = transportadora
		}
		saida["data_registro"] = data_registro.Format(time.RFC3339)

		if hasDistanciaApp {
			saida["distancia_app"] = distancia_app
		}
		if hasDiasEvidenciaApp {
			saida["dias_evidencia_app"] = dias_evidencia_app
		}

		saida["cod_saida"] = cod_saida
		saida["corporativo"] = corporativo
		saida["entidade"] = entidade
		saida["unidade"] = unidade
		saida["gerador"] = gerador
		saida["destinatario"] = destinatario
		saida["acondicionamento"] = acondicionamento
		saida["mtr"] = mtr
		saida["origem"] = origem
		saida["lacre"] = lacre
		saida["lote"] = lote
		saida["veiculo"] = veiculo
		saida["motorista"] = motorista
		saida["categoria_nome"] = categoria_nome
		saida["tipo_nome"] = tipo_nome
		saida["status"] = status

		saida["campo_controle"] = campo_controle
		saida["campo_saida"] = campo_saida

		saidaJSON, nerr := json.Marshal(saida)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		saidaAsset, err := assets.NewAsset(saida)

		_, err = saidaAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		// Marshal asset back to JSON format
		saidaJSON, nerr = json.Marshal(saidaAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return saidaJSON, nil
	},
}
