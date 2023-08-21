package group

import (
	"wechat-bot/dao/gpt"

	"github.com/google/wire"
	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(
	NewGroup,
)

type Group struct {
	log *zap.SugaredLogger
	gpt *gpt.GPT
}

func NewGroup(log *zap.SugaredLogger, gpt *gpt.GPT) *Group {
	return &Group{log: log, gpt: gpt}
}
