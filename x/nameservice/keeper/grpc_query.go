package keeper

import (
	"github.com/orionis-labs/orionis/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
