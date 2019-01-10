// +build test

package main_test

import (
	main "my/go-my-sandbox/mock_with_build_tag"
	"my/go-my-sandbox/mock_with_build_tag/client"
	"testing"
)

func TestFoo(t *testing.T) {
	expected := "TestFoo"

	client.MockTag = func() string {
		return expected
	}

	s := main.NewService()
	got := s.Run()

	if got != expected {
		t.Errorf("got %s, expected %s\n", got, expected)
	}
}
