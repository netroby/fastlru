package fastlru

import (
	"fmt"
	"testing"
)

func Test_lru(t *testing.T) {
	lc := New(10)
	lc.Add("hello", "Hello world")
	val, ok := lc.Get("hello")

	if ok != true {
		fmt.Println(val)
		t.Fatalf("%s: cache hit = %v, want %v", "hello", ok, !ok)
	} else {
		fmt.Println(val)
	}
}
