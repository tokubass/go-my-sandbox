// +build !test,!use_lock

package client

import "my/go-my-sandbox/mock_with_build_tag/domain"

type C struct {
}

func NewClient() domain.C {
	return &C{}
}

func (c *C) Tag() string {
	return "production"
}
