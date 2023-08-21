package wechat

import (
	"wechat-bot/biz/friend"
	"wechat-bot/biz/group"

	"github.com/eatmoreapple/openwechat"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(
	NewWechat,
)

type Wechat struct {
	bot    *openwechat.Bot
	friend *friend.Friend
	group  *group.Group
	log    *zap.SugaredLogger
}

func NewWechat(
	bot *openwechat.Bot,
	friend *friend.Friend,
	group *group.Group,
	log *zap.SugaredLogger,
) *Wechat {
	return &Wechat{
		bot:    bot,
		friend: friend,
		group:  group,
		log:    log,
	}
}
