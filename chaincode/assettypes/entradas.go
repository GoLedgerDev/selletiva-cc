package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// SampleBook is the demo subAsset for the template
var Entradas = assets.AssetType{
	Tag:         "entradas",
	Label:       "Entradas",
	Description: "Entradas",

	Props: []assets.AssetProp{
		{
			Tag:      "id",
			Label:    "ID",
			Required: true,
			IsKey:    true,
			DataType: "number",
		},
		{
			Tag:      "cod_entrada",
			Label:    "Código Entrada",
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
			Tag:      "data_registro",
			Label:    "Data do registro de Entrada",
			DataType: "datetime",
		},
		{
			Tag:      "ent_sai",
			Label:    "Controle de Entradas",
			DataType: "boolean",
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
