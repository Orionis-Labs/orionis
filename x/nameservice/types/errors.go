package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/nameservice module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
	// this line is used by starport scaffolding # ibc/errors
)

var (
	ErrNameDoesNotExist = sdkerrors.Register(ModuleName, 1, "name does not exist")
)
