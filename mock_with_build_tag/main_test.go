package main

import (
	"my/go-my-sandbox/mock_with_build_tag/client"
	"testing"
)

func TestFoo(t *testing.T) {
	expected := "TestFoo"

	c := client.NewClient()
	c.MockTag = func() string {
		return expected
	}

	got := c.Tag()
	if got != expected {
		t.Errorf("got %s, expected %s\n", got, expected)
	}
}
