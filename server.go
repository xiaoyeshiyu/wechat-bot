package main

import (
	"wechat-bot/service/wechat"
)

type Server struct {
	wechat *wechat.Wechat
}

func NewServer(wechat *wechat.Wechat) *Server {
	return &Server{wechat: wechat}
}

func (s *Server) Start() error {
	err := s.wechat.Start()
	if err != nil {
		return err
	}

	return s.wechat.Block()
}
