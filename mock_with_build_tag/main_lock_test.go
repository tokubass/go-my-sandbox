// +build use_lock

package main_test

import (
	main "my/go-my-sandbox/mock_with_build_tag"
	"my/go-my-sandbox/mock_with_build_tag/client"
	"testing"
)

func TestFoo(t *testing.T) {
	tests := []struct {
		name  string
		value int
	}{
		{name: "test 1"}, {name: "test 2"}, {name: "test 3"}, {name: "test 4"},
		{name: "test 5"}, {name: "test 6"}, {name: "test 7"}, {name: "test 8"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			//Set_Client_MockTagでmutex.Lock()していない場合テストが失敗する
			t.Parallel()
			expected := tc.name
			resetMock := client.Set_Client_MockTag(func() string {
				return expected
			})
			defer resetMock()

			s := main.NewService()
			got := s.Run()
			if got != expected {
				t.Errorf("got %s, expected %s\n", got, expected)
			}

		})
	}
}
