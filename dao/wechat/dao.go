package wechat

import "github.com/eatmoreapple/openwechat"

func NewBot() *openwechat.Bot {
	//bot := openwechat.DefaultBot(openwechat.Normal) // 网页模式，登录不上可以切换到桌面模式
	return openwechat.DefaultBot(openwechat.Desktop) // 桌面模式
}
