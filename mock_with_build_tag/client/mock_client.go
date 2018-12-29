// +build test

package client

import "my/go-my-sandbox/mock_with_build_tag/domain"

type C struct {
	domain.C
	MockTag func() string
}

func NewClient() *C {
	return &C{
		MockTag: func() string {
			return "mock tag"
		},
	}
}

func (c *C) Tag() string {
	return c.MockTag()
}
