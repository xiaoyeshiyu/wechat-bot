package dao

import (
	"wechat-bot/dao/gpt"
	"wechat-bot/dao/wechat"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	gpt.NewGPT,
	wechat.NewBot,
)
