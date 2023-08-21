package group

import (
	"context"
	"strings"

	"github.com/eatmoreapple/openwechat"
	"github.com/pkg/errors"
)

func (g *Group) Handle(ctx context.Context, msg *openwechat.Message) error {
	// 非@不处理
	if !msg.IsAt() {
		return nil
	}

	switch msg.MsgType {
	case openwechat.MsgTypeText:
		return g.replyText(ctx, msg)
	default:
		_, err := msg.ReplyText("不支持的消息")
		return errors.WithStack(err)
	}
}

func (g *Group) replyText(ctx context.Context, msg *openwechat.Message) error {
	// 获取@我的用户
	groupSender, err := msg.SenderInGroup()
	if err != nil {
		return errors.WithStack(err)
	}
	// 回复信息前面加前缀
	atText := "@" + groupSender.NickName

	// 替换掉@文本，然后向GPT发起请求
	replaceText := "@" + groupSender.NickName
	requestText := strings.TrimSpace(strings.ReplaceAll(msg.Content, replaceText, ""))
	reply, err := g.gpt.CreateChatCompletion(ctx, requestText)
	if err != nil {
		msg.ReplyText(atText + "机器人神了，我一会发现了就去修。")
		return errors.WithStack(err)
	}

	// 回复@我的用户
	replyText := atText + reply
	_, err = msg.ReplyText(replyText)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
