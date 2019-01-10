// +build test

package client

import "my/go-my-sandbox/mock_with_build_tag/domain"

// 並列処理で書き換えられると困る
var MockTag = func() string { return "naive mock" }

type C struct {
	domain.C
	MockTag func() string
}

func NewClient() *C {
	return &C{MockTag: MockTag}
}

func (c *C) Tag() string {
	return c.MockTag()
}
