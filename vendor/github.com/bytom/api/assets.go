package api

import (
	"context"
	"strings"

	"corepool/core/asset"
	"corepool/core/crypto/ed25519/chainkd"
	chainjson "corepool/core/encoding/json"

	log "github.com/sirupsen/logrus"
)

// POST /create-asset
func (a *API) createAsset(ctx context.Context, ins struct {
	Alias           string                 `json:"alias"`
	RootXPubs       []chainkd.XPub         `json:"root_xpubs"`
	Quorum          int                    `json:"quorum"`
	Definition      map[string]interface{} `json:"definition"`
	IssuanceProgram chainjson.HexBytes     `json:"issuance_program"`
}) Response {
	ass, err := a.wallet.AssetReg.Define(
		ins.RootXPubs,
		ins.Quorum,
		ins.Definition,
		strings.ToUpper(strings.TrimSpace(ins.Alias)),
		ins.IssuanceProgram,
	)
	if err != nil {
		return NewErrorResponse(err)
	}

	annotatedAsset, err := asset.Annotated(ass)
	if err != nil {
		return NewErrorResponse(err)
	}

	log.WithField("asset ID", annotatedAsset.ID.String()).Info("Created asset")

	return NewSuccessResponse(annotatedAsset)
}

// POST /update-asset-alias
func (a *API) updateAssetAlias(updateAlias struct {
	ID       string `json:"id"`
	NewAlias string `json:"alias"`
}) Response {
	if err := a.wallet.AssetReg.UpdateAssetAlias(updateAlias.ID, updateAlias.NewAlias); err != nil {
		return NewErrorResponse(err)
	}

	return NewSuccessResponse(nil)
}
