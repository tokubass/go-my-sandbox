// +build test

package client

type C struct {
}

func NewClient() *C {
	return &C{}
}

func (c *C) Tag() string {
	return "test"
}
