package friend

import (
	"wechat-bot/dao/gpt"

	"github.com/google/wire"
	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(
	NewFriend,
)

type Friend struct {
	log *zap.SugaredLogger
	gpt *gpt.GPT
}

func NewFriend(log *zap.SugaredLogger, gpt *gpt.GPT) *Friend {
	return &Friend{log: log, gpt: gpt}
}
