package keeper

import (
	"github.com/zarazan/apollo/x/apollo/types"
)

var _ types.QueryServer = Keeper{}
