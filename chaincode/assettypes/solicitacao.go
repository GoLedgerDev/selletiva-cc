package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// SampleBook is the demo subAsset for the template
var Solicitacao = assets.AssetType{
	Tag:         "solicitacao",
	Label:       "Solicitacao",
	Description: "Solicitacao",

	Props: []assets.AssetProp{
		{
			Tag:      "id",
			Label:    "ID",
			Required: true,
			IsKey:    true,
			DataType: "number",
		},
		{
			Tag:      "categoria",
			Label:    "Categoria",
			Required: true,
			DataType: "number",
		},
		{
			Tag:      "tipo",
			Label:    "Tipo de Resídua",
			DataType: "number",
		},
		{
			Tag:      "quantidade",
			Label:    "Quantidade",
			DataType: "number",
		},
		{
			Tag:      "transportadora",
			Label:    "Transportadora",
			DataType: "number",
		},
		{
			Tag:      "observacao",
			Label:    "Observação",
			DataType: "string",
		},
		{
			Tag:      "dta_solicita",
			Label:    "Data da Solicitação",
			DataType: "datetime",
		},
		{
			Tag:      "peso_real",
			Label:    "Peso real",
			DataType: "number",
		},
		{
			Tag:      "dta_recebimento",
			Label:    "Data Recebimento",
			DataType: "datetime",
		},
		{
			Tag:      "status",
			Label:    "Status",
			DataType: "string",
		},
		{
			Tag:      "distancia_app",
			Label:    "Distância",
			DataType: "number",
		},
		{
			Tag:      "dias_evidencia_app",
			Label:    "Dias da Evidencia",
			DataType: "number",
		},
		{
			Tag:      "campo_saida",
			Label:    "Campo Saida",
			DataType: "string",
		},
		{
			Tag:      "campo_controle",
			Label:    "Campo Controle",
			DataType: "boolean",
		},
	},
}
