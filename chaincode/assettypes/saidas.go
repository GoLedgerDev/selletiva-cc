package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// SampleBook is the demo subAsset for the template
var Saidas = assets.AssetType{
	Tag:         "saidas",
	Label:       "Saídas",
	Description: "Saídas",

	Props: []assets.AssetProp{
		{
			Tag:      "id",
			Label:    "ID",
			Required: true,
			IsKey:    true,
			DataType: "number",
		},
		{
			Tag:      "cod_saida",
			Label:    "Código Saída",
			DataType: "number",
		},
		{
			Tag:      "corporativo",
			Label:    "Corporativo",
			DataType: "number",
		},
		{
			Tag:      "entidade",
			Label:    "Entidade Gerenciadora",
			DataType: "number",
		},
		{
			Tag:      "categoria",
			Label:    "Categoria",
			DataType: "number",
		},
		{
			Tag:      "tipo",
			Label:    "Tipo de Resíduo",
			DataType: "number",
		},
		{
			Tag:      "quantidade",
			Label:    "Quantidade",
			DataType: "number",
		},
		{
			Tag:      "unidade",
			Label:    "Unidade",
			DataType: "number",
		},
		{
			Tag:      "gerador",
			Label:    "Gerador",
			DataType: "number",
		},
		{
			Tag:      "transportadora",
			Label:    "Transportadora",
			DataType: "number",
		},
		{
			Tag:      "destinatario",
			Label:    "Destinatario",
			DataType: "number",
		},
		{
			Tag:      "acondicionamento",
			Label:    "Acondicionamento",
			DataType: "number",
		},
		{
			Tag:      "mtr",
			Label:    "MTR",
			DataType: "string",
		},
		{
			Tag:      "origem",
			Label:    "Origem",
			DataType: "number",
		},
		{
			Tag:      "lacre",
			Label:    "Lacre",
			DataType: "string",
		},
		{
			Tag:      "lote",
			Label:    "Lote",
			DataType: "string",
		},
		{
			Tag:      "veiculo",
			Label:    "Veiculo",
			DataType: "number",
		},
		{
			Tag:      "motorista",
			Label:    "Motorista",
			DataType: "number",
		},
		{
			Tag:      "data_registro",
			Label:    "Data do registro de Saida",
			DataType: "datetime",
		},
		{
			Tag:      "dias_evidencia_app",
			Label:    "Dias da Evidencia do APP",
			DataType: "number",
		},
		{
			Tag:      "distancia_app",
			Label:    "Distancia APP",
			DataType: "number",
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
