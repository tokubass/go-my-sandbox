// +build use_lock

package client

import (
	"my/go-my-sandbox/mock_with_build_tag/domain"
	"sync"
)

// 並列処理で書き換えられると困るのでlock
var defaultMockTag = func() string { return "use_lock" }
var mockTag = defaultMockTag
var mutex sync.Mutex

func Set_Client_MockTag(f func() string) func() error {
	mutex.Lock() //unlockせずにもう1回Lockを呼ぶとpanic発生

	mockTag = f
	return func() error {
		mockTag = defaultMockTag
		mutex.Unlock()
		return nil
	}
}

type C struct {
	domain.C
	MockTag func() string
}

func NewClient() *C {
	return &C{MockTag: mockTag}
}

func (c *C) Tag() string {
	return c.MockTag()
}
