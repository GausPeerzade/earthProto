package keeper

import (
	"earthProto/x/earthproto/types"
)

var _ types.QueryServer = Keeper{}
