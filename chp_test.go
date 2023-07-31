package chp

import (
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

func TestTake(t *testing.T) {
	cs, sent := nchan(1, 10)
	c := cs[0]

	d := take(c, 2)
	var got []int
	for v := range d {
		got = append(got, v)
	}

	equalslice(t, got, sent[:2])
}

func TestMerge(t *testing.T) {
	cs, sent := nchan(1000, 5)

	merged := merge(nil, 42, cs...)
	var got []int
	for v := range merged {
		got = append(got, v)
	}

	sort.Ints(got)
	sort.Ints(sent)
	equalslice(t, got, sent)
}

func TestLast(t *testing.T) {
	cs, sent := nchan(1, 10)
	got := last(cs...)

	if sent[len(sent)-1] != got {
		t.Errorf("unexpected value %d", got)
	}
}

func TestFirst(t *testing.T) {
	cs, sent := nchan(1, 10)
	got := first(cs...)

	if sent[0] != got {
		t.Errorf("unexpected value %d", got)
	}
}

func TestCollect(t *testing.T) {
	cs, sent := nchan(1, 100)
	got := collect(cs[0])

	sort.Ints(got)
	sort.Ints(sent)
	equalslice(t, got, sent)
}

// nchan creates n channels, and sends k random values into each of the
// n channels. It closes each channel after sending the values. nchan
// returns the created channels and all of the sent values.
func nchan(n, k int) ([]chan int, []int) {
	var cs []chan int
	var sent []int

	for i := 0; i < n; i++ {
		// create channel: i/n.
		c := make(chan int)
		cs = append(cs, c)

		// send k values on the channel.
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

func equalslice[T comparable](t *testing.T, got, want []T) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("len: got %d, want %d", len(got), len(want))
		return
	}
	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("[%d]: got %v, want %v", i, got[i], want[i])
		}
	}
}
