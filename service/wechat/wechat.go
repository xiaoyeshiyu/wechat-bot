package wechat

import (
	"context"

	"wechat-bot/config"

	"github.com/eatmoreapple/openwechat"
	"github.com/pkg/errors"
)

func (w *Wechat) Start() error {
	// 注册消息处理函数
	w.bot.MessageHandler = w.Service
	// 注册登陆二维码回调
	w.bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 创建热存储容器对象
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	// 执行热登录
	err := w.bot.HotLogin(reloadStorage)
	if err != nil {
		if err = w.bot.Login(); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (w *Wechat) Service(msg *openwechat.Message) {
	var err error
	ctx := context.Background()
	switch {
	case msg.IsSendByGroup():
		// 处理群消息
		err = w.group.Handle(ctx, msg)
	case msg.IsFriendAdd():
		// 好友申请
		if config.GetAutoPass() {
			_, err = msg.Agree("你好我是基于chatGPT引擎开发的微信机器人，你可以向我提问任何问题。")
		}
	case msg.IsSendByFriend(), msg.IsSendBySelf():
		// 处理私聊
		err = w.friend.Handle(ctx, msg)
	}

	if err != nil {
		w.log.Errorf("%+v", err)
		return
	}
}

func (w *Wechat) Block() error {
	return w.bot.Block()
}
