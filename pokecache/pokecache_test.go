package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for index, nextCase := range cases {
		t.Run(fmt.Sprintf("Test case %v", index), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(nextCase.key, nextCase.val)

			val, ok := cache.Get(nextCase.key)

			if !ok {
				t.Errorf("expected to find key")
				return
			}

			if string(val) != string(nextCase.val) {
				t.Errorf("expected original value to match cache value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 10*time.Second

	key := "https://example.com"
	val := []byte("testdata")

	cache := NewCache(baseTime)
	cache.Add(key, val)

	_, ok := cache.Get(key)

	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get(key)

	if ok {
		t.Errorf("expected to not find key")
	}
}
