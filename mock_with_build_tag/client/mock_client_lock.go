// +build use_lock

package client

import (
	"my/go-my-sandbox/mock_with_build_tag/domain"
	"reflect"
	"sync"
)

type C struct {
	domain.C
	MockTag func() string
}

// mockがdomainのinterfaceから乖離していないかチェック
func init() {
	c := &C{}
	if reflect.TypeOf(c.Tag) != reflect.TypeOf(c.MockTag) {
		panic("mock interface should be the same as domain interface")
	}
}

var defaultMockTag = func() string { return "use_lock" }
var mockTag = defaultMockTag

func NewClient() *C {
	return &C{MockTag: mockTag}
}

func (c *C) Tag() string {
	return c.MockTag()
}

// 並列処理で書き換えられると困るのでlock
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
