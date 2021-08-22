package keeper

import (
	"github.com/orionis-labs/orionis/x/orionis/types"
)

var _ types.QueryServer = Keeper{}
