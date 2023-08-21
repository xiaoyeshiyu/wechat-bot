//go:build wireinject
// +build wireinject

package main

import (
	"wechat-bot/biz"
	"wechat-bot/dao"
	"wechat-bot/service"

	"github.com/google/wire"
	"go.uber.org/zap"
)

func initServer(log *zap.SugaredLogger) (*Server, error) {
	panic(wire.Build(
		dao.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		NewServer,
	))

	return &Server{}, nil
}
