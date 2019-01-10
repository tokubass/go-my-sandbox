package main

import (
	"fmt"
	"my/go-my-sandbox/mock_with_build_tag/client"
	"my/go-my-sandbox/mock_with_build_tag/domain"
)

type Service struct {
	client domain.C
}

func NewService() *Service {
	//clientが引数から取得できない状況設定
	return &Service{client: client.NewClient()}
}

func (s *Service) Run() string {
	return s.client.Tag()
}

func main() {
	s := NewService()
	fmt.Println(s.Run())
}
