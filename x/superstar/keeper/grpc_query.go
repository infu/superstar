package keeper

import (
	"github.com/infu/superstar/x/superstar/types"
)

var _ types.QueryServer = Keeper{}
