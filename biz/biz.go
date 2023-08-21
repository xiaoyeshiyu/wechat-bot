package biz

import (
	"wechat-bot/biz/friend"
	"wechat-bot/biz/group"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	friend.ProviderSet,
	group.ProviderSet,
)
