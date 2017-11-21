package bitslice

import (
	"testing"
)

func TestSetBit(t *testing.T) {
	const (
		max = 10
		off = 9
	)
	s := NewBits(max)
	if err := s.Set(off); err != nil {
		t.Fatal(err)
	}
	if set, err := s.Get(off); err != nil {
		t.Fatal(err)
	} else if !set {
		x := off / 64
		t.Fatalf("expected index: %d to be set (%b)", off, s.data[x])
	}
}

func TestNotSet(t *testing.T) {
	const (
		max = 10
		off = 9
	)
	s := NewBits(max)
	if err := s.Set(off); err != nil {
		t.Fatal(err)
	}
	less := off - 1
	if set, err := s.Get(less); err != nil {
		t.Fatal(err)
	} else if set {
		x := less / 64
		t.Fatalf("expected index: %d to not be set (%b)", less, s.data[x])
	}
}

func TestRangeError(t *testing.T) {
	const (
		max = 1024
	)
	t.Log("MAX:", max)
	s := NewBits(max)
	if _, err := s.Get(max + 1); err == nil {
		t.Fatal("expected range error")
	} else {
		t.Log("got expected range error: " + err.Error())
	}
}

func TestSeries(t *testing.T) {
	const max = 32 * 1024
	fib := []int{
		0, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657,
	}
	lookup := make(map[int]struct{})
	for _, x := range fib {
		lookup[x] = struct{}{}
	}
	s := NewBits(max)
	for _, bit := range fib {
		if err := s.Set(bit); err != nil {
			t.Fatal(err)
		}
	}
	for bit := 0; bit < max; bit++ {
		set, err := s.Get(bit)
		if err != nil {
			t.Fatal(err)
		}
		_, ok := lookup[bit]
		if set != ok {
			t.Errorf("offset %d should be %t but is %t\n", bit, ok, set)
		}
	}
}

func TestSetGet(t *testing.T) {
	const max = 32 * 1024
	fib := []int{
		0, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657,
	}
	lookup := make(map[int]struct{})
	for _, x := range fib {
		lookup[x] = struct{}{}
	}
	s := NewBits(max)
	for _, bit := range fib {
		if err := s.Set(bit); err != nil {
			t.Fatal(err)
		}
	}
	for bit := 0; bit < max; bit++ {
		set, err := s.SetGet(bit)
		if err != nil {
			t.Fatal(err)
		}
		_, ok := lookup[bit]
		if set != ok {
			t.Errorf("offset %d should be %t but is %t\n", bit, ok, set)
		}
	}
}
