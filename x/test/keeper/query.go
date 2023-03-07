package keeper

import (
	"github.com/coinhall/test/x/test/types"
)

var _ types.QueryServer = Keeper{}
