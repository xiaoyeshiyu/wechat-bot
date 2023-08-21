package service

import (
	"wechat-bot/service/wechat"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	wechat.ProviderSet,
)
