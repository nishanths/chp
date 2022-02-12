package chp

import (
	"fmt"
	"os"
	"sort"
	"testing"
)

func TestMain(m *testing.M) {
	// rand.Seed(time.Now().Unix())
	os.Exit(m.Run())
}

func nchan(n int) ([]chan int, []int) {
	var cs []chan int
	var sent []int

	for i := 0; i < n; i++ {
		c := make(chan int)
		cs = append(cs, c)
		sent = append(sent, i)
		go func(c chan int, v int) {
			c <- v
			close(c)
		}(c, i)
	}

	return cs, sent
}

func TestFanIn(t *testing.T) {
	cs, sent := nchan(10000)

	combined := FanIn(nil, cs...)
	var got []int
	for v := range combined {
		got = append(got, v)
	}

	sort.Ints(got)
	if reason, ok := equalslice(sent, got); !ok {
		t.Error(reason)
	}
}

func TestFirst(t *testing.T) {
	cs, sent := nchan(10)
	got := First(cs...)

	if !contains(sent, got) {
		t.Errorf("unexpected value: %d", got)
	}
}

func contains[T comparable](haystack []T, needle T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func equalslice[T comparable](want, got []T) (reason string, ok bool) {
	if len(want) != len(got) {
		return fmt.Sprintf("unequal slice lengths: want %d, got %d", len(want), len(got)), false
	}
	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			return fmt.Sprintf("wrong slice element at index %d: want %v, got %v", i, want[i], got[i]), false
		}
	}
	return "", true
}
