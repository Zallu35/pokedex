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

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}
func TestReapLoop(t *testing.T) {
	testCache := NewCache(2 * time.Second)
	keyOne := "A bunch of stuff"
	keyTwo := "A bunch of stuff that happened next"
	valueOne := []byte("squishmallows")
	valueTwo := []byte("Dogman Books are goofy")

	testCache.Add(keyOne, valueOne)
	if _, ok := testCache.Get(keyOne); !ok {
		t.Fatalf("should find key immediately after adding")
	}

	time.Sleep(3 * time.Second)
	testCache.Add(keyTwo, valueTwo)

	if _, ok := testCache.Get(keyOne); ok {
		t.Fatalf("should not find keyOne after interval has passed")
	}
	if _, ok := testCache.Get(keyTwo); !ok {
		t.Fatalf("keyTwo added after interval, should be found")
	}
}
