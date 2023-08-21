package friend

import (
	"context"

	"github.com/eatmoreapple/openwechat"
	"github.com/pkg/errors"
)

func (f *Friend) Handle(ctx context.Context, msg *openwechat.Message) error {
	switch msg.MsgType {
	case openwechat.MsgTypeText:
		return f.replyText(ctx, msg)
	default:
		_, err := msg.ReplyText("不支持的消息")
		return errors.WithStack(err)
	}
}

func (f *Friend) replyText(ctx context.Context, msg *openwechat.Message) error {
	sender, err := msg.Sender()
	if err != nil {
		return err
	}
	f.log.Infof("receive message from %s: %s", sender.NickName, msg.Content)

	// 向GPT发起请求
	chatCompletion, err := f.gpt.CreateChatCompletion(ctx, msg.Content)
	if err != nil {
		msg.ReplyText("机器人神了，我一会发现了就去修。")
		return errors.WithStack(err)
	}

	// 回复用户
	_, err = msg.ReplyText(chatCompletion)
	return errors.WithStack(err)
}
