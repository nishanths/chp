package chp

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().Unix())
	os.Exit(m.Run())
}

func TestMerge(t *testing.T) {
	cs, sent := nchan(1000, 5)

	merged := Merge(nil, cs...)
	var got []int
	for v := range merged {
		got = append(got, v)
	}

	sort.Ints(got)
	sort.Ints(sent)
	if err := equalslice(sent, got); err != nil {
		t.Error(err)
	}
}

func TestFirst(t *testing.T) {
	cs, sent := nchan(10, 1)
	got := First(cs...)

	if !contains(sent, got) {
		t.Errorf("unexpected value: %d", got)
	}
}

func TestCollect(t *testing.T) {
	cs, sent := nchan(1, 100)
	got := Collect(cs[0])

	sort.Ints(got)
	sort.Ints(sent)
	if err := equalslice(sent, got); err != nil {
		t.Error(err)
	}
}

// nchan creates n channels, and sends k value into each of the n channels.
// It closes each channel after sending the values.
// nchan returns the created channels and all of the sent values.
func nchan(n, k int) ([]chan int, []int) {
	var cs []chan int
	var sent []int

	for i := 0; i < n; i++ {
		// create the channel: i/n.
		c := make(chan int)
		cs = append(cs, c)

		// send k values on this channel.
		vs := randslice(k)
		sent = append(sent, vs...)
		go func(c chan int, vs []int) {
			for _, v := range vs {
				c <- v
			}
			close(c)
		}(c, vs)
	}

	return cs, sent
}

func randslice(size int) []int {
	out := make([]int, size)
	for i := 0; i < size; i++ {
		out[i] = rand.Intn(1000)
	}
	return out
}

func contains[T comparable](haystack []T, needle T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func equalslice[T comparable](want, got []T) error {
	if len(want) != len(got) {
		return fmt.Errorf("unequal slice lengths: want %d, got %d", len(want), len(got))
	}
	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			return fmt.Errorf("wrong slice element at index %d: want %v, got %v", i, want[i], got[i])
		}
	}
	return nil
}
