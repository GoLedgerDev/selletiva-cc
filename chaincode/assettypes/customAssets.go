package assettypes

import (
	"github.com/goledgerdev/cc-tools/assets"
)

// CustomAssets contains all assets inserted via GoFabric's Template mode.
// For local development, this can be empty or could contain assets that
// are supposed to be defined via the Template mode.
var CustomAssets = []assets.AssetType{
	Solicitacao,
	Saidas,
	Entradas,
}
