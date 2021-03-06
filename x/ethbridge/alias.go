// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/peggy/x/ethbridge/querier
// ALIASGEN: github.com/cosmos/peggy/x/ethbridge/types
package ethbridge

import (
	"github.com/cosmos/peggy/x/ethbridge/keeper"
	"github.com/cosmos/peggy/x/ethbridge/querier"
	"github.com/cosmos/peggy/x/ethbridge/types"
)

const (
	QueryEthProphecy       = querier.QueryEthProphecy
	DefaultCodespace       = types.DefaultCodespace
	CodeInvalidEthNonce    = types.CodeInvalidEthNonce
	CodeInvalidEthAddress  = types.CodeInvalidEthAddress
	CodeErrJSONMarshalling = types.CodeErrJSONMarshalling
	ModuleName             = types.ModuleName
	StoreKey               = types.StoreKey
	QuerierRoute           = types.QuerierRoute
	RouterKey              = types.RouterKey
)

var (
	// functions aliases
	NewKeeper                         = keeper.NewKeeper
	NewQuerier                        = querier.NewQuerier
	NewEthBridgeClaim                 = types.NewEthBridgeClaim
	NewOracleClaimContent             = types.NewOracleClaimContent
	CreateOracleClaimFromEthClaim     = types.CreateOracleClaimFromEthClaim
	CreateEthClaimFromOracleString    = types.CreateEthClaimFromOracleString
	CreateOracleClaimFromOracleString = types.CreateOracleClaimFromOracleString
	RegisterCodec                     = types.RegisterCodec
	ErrInvalidEthNonce                = types.ErrInvalidEthNonce
	ErrInvalidEthAddress              = types.ErrInvalidEthAddress
	ErrJSONMarshalling                = types.ErrJSONMarshalling
	NewEthereumAddress                = types.NewEthereumAddress
	NewMsgCreateEthBridgeClaim        = types.NewMsgCreateEthBridgeClaim
	MapOracleClaimsToEthBridgeClaims  = types.MapOracleClaimsToEthBridgeClaims
	NewQueryEthProphecyParams         = types.NewQueryEthProphecyParams
	NewQueryEthProphecyResponse       = types.NewQueryEthProphecyResponse

	CreateTestEthMsg                   = types.CreateTestEthMsg
	CreateTestEthClaim                 = types.CreateTestEthClaim
	CreateTestQueryEthProphecyResponse = types.CreateTestQueryEthProphecyResponse
)

type (
	Keeper                   = keeper.Keeper
	EthBridgeClaim           = types.EthBridgeClaim
	OracleClaimContent       = types.OracleClaimContent
	CodeType                 = types.CodeType
	EthereumAddress          = types.EthereumAddress
	MsgCreateEthBridgeClaim  = types.MsgCreateEthBridgeClaim
	MsgBurn                  = types.MsgBurn
	MsgLock                  = types.MsgLock
	QueryEthProphecyParams   = types.QueryEthProphecyParams
	QueryEthProphecyResponse = types.QueryEthProphecyResponse
)
