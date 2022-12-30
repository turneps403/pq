package prqueue_test

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/turneps403/prqueue"
)

// TestNew_GenericType will check generics specific on a time of creation object
func TestNew_GenericType(t *testing.T) {
	t.Run("New for generic int", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("func New is in a panic: %v", r)
			}
		}()
		pq := prqueue.New(func(a, b int) bool {
			return a < b
		})
		_ = pq
	})

	t.Run("New for generic string", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("func New is in a panic: %v", r)
			}
		}()
		pq := prqueue.New(func(a, b string) bool {
			return a < b
		})
		_ = pq
	})

	t.Run("New for generic struct", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("func New is in a panic: %v", r)
			}
		}()
		type tmp struct {
			t int
		}
		pq := prqueue.New(func(a, b tmp) bool {
			return a.t < b.t
		})
		_ = pq
	})

	t.Run("New for generic pointer", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("func New is in a panic: %v", r)
			}
		}()
		type tmp struct {
			t int
		}
		pq := prqueue.New(func(a, b *tmp) bool {
			return a.t < b.t
		})
		_ = pq
	})

}

func TestAdd(t *testing.T) {
	t.Run("Add for int", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("func Add is in a panic: %v", r)
			}
		}()
		pq := prqueue.New(func(a, b int) bool {
			return a < b
		})
		for i := 0; i < 10; i++ {
			pq.Add(i)
		}
	})

	t.Run("Add for struct", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("func Add is in a panic: %v", r)
			}
		}()
		type tmp struct {
			t int
		}
		pq := prqueue.New[tmp](func(a, b tmp) bool {
			return a.t < b.t
		})
		for i := 0; i < 10; i++ {
			pq.Add(tmp{i})
		}
	})
}

func TestLen(t *testing.T) {
	pq := prqueue.New(func(a, b int) bool {
		return a < b
	})
	max := rand.Intn(100) + 100
	for i := 0; i < max; i++ {
		pq.Push(i)
		if pq.Len() != i+1 {
			t.Fatalf("func Len isn't correct: expected %v but got %v", i, pq.Len())
		}
	}
	t.Logf("func Len was successfully checked in %d loops", max)
}

func TestPoll0(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("func Poll is in a panic: %v", r)
		}
	}()
	pq := prqueue.New(func(a, b int) bool {
		return a < b
	})
	if _, err := pq.Poll(); err == nil {
		t.Errorf("func Poll from empty queue doesnt return error")
	}
	tint := 100500
	pq.Add(tint)
	v, err := pq.Poll()
	if err != nil {
		t.Errorf("func Poll from queue with 1 element return error: %v", err)
	}
	if v != tint {
		t.Errorf("func Poll from queue with 1 element return wron value: expected %v got %v", tint, v)
	}
	if _, err := pq.Poll(); err == nil {
		t.Errorf("func Poll from empty queue (second) doesnt return error")
	}
}

func TestPoll1(t *testing.T) {
	t.Run("Min heap correctness", func(t *testing.T) {
		l := []int{1, 3, 5, 2, 4}
		ordl := make([]int, len(l))
		copy(ordl, l)
		sort.Ints(ordl)

		pq := prqueue.New(func(a, b int) bool {
			return a < b
		})
		for _, v := range l {
			pq.Add(v)
		}

		for i := 0; i < len(ordl); i++ {
			e, _ := pq.Poll()
			if ordl[i] != e {
				t.Errorf("func Poll as min heap broken: expected %v got %v", ordl[i], e)
			}
		}
	})

	t.Run("Max heap correctness", func(t *testing.T) {
		l := []int{1, 3, 5, 2, 4}
		ordl := make([]int, len(l))
		copy(ordl, l)
		sort.Sort(sort.Reverse(sort.IntSlice(ordl)))

		pq := prqueue.New(func(a, b int) bool {
			return a > b
		})
		for _, v := range l {
			pq.Add(v)
		}

		for i := 0; i < len(ordl); i++ {
			e, _ := pq.Poll()
			if ordl[i] != e {
				t.Errorf("func Poll as max heap broken: expected %v got %v", ordl[i], e)
			}
		}
	})
}

func TestPoll2(t *testing.T) {
	pq := prqueue.New(func(a, b int) bool {
		return a > b
	})
	_, err := pq.Poll()
	if err == nil || !errors.Is(err, prqueue.ErrEmptyQueue) {
		t.Errorf("func Poll returned unexpected error: expected %v got %v", "prqueue.ErrEmptyQueue", err)
	}
}

func TestIsEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("func IsEmpty is in a panic: %v", r)
		}
	}()
	pq := prqueue.New(func(a, b int) bool {
		return a < b
	})
	if !pq.IsEmpty() {
		t.Errorf("func IsEmpty has broken: has to be empty")
	}
	pq.Add(100500)
	if pq.IsEmpty() {
		t.Errorf("func IsEmpty has broken: has to be not empty")
	}
	_, _ = pq.Poll()
	if !pq.IsEmpty() {
		t.Errorf("func IsEmpty has broken: has to be empty")
	}
}

func TestPeek0(t *testing.T) {
	t.Run("Min heap correctness", func(t *testing.T) {
		l := []int{1, 3, 5, 2, 4}

		pq := prqueue.New(func(a, b int) bool {
			return a < b
		})
		for _, v := range l {
			pq.Add(v)
		}

		for i := 0; i < len(l); i++ {
			v, _ := pq.Peek()
			e, _ := pq.Poll()
			if v != e {
				t.Errorf("func Peek as min heap broken: expected %v got %v", e, v)
			}
		}
	})

	t.Run("Max heap correctness", func(t *testing.T) {
		l := []int{1, 3, 5, 2, 4}

		pq := prqueue.New(func(a, b int) bool {
			return a > b
		})
		for _, v := range l {
			pq.Add(v)
		}

		for i := 0; i < len(l); i++ {
			v, _ := pq.Peek()
			e, _ := pq.Poll()
			if v != e {
				t.Errorf("func Peek as min heap broken: expected %v got %v", e, v)
			}
		}
	})
}

func TestPeek1(t *testing.T) {
	pq := prqueue.New(func(a, b int) bool {
		return a > b
	})
	_, err := pq.Peek()
	if err == nil || !errors.Is(err, prqueue.ErrEmptyQueue) {
		t.Errorf("func Peek returned unexpected error: expected %v got %v", "prqueue.ErrEmptyQueue", err)
	}
}

func TestString(t *testing.T) {
	t.Run("String representation for int", func(t *testing.T) {
		l := []int{1, 3, 5, 2, 4}
		heapVer := "[1 2 5 3 4]"
		pq := prqueue.New(func(a, b int) bool {
			return a < b
		})
		for _, v := range l {
			pq.Add(v)
		}
		if fmt.Sprintf("%v", pq) != heapVer {
			t.Errorf("func String has broken format: expected %v got %v", heapVer, fmt.Sprintf("%v", pq))
		}
	})

	t.Run("String representation for rune", func(t *testing.T) {
		l := []string{"a", "z", "b", "k", "c"}
		heapVer := "[a c b z k]"
		pq := prqueue.New[string](func(a, b string) bool {
			return a < b
		})
		for _, v := range l {
			pq.Add(v)
		}
		if fmt.Sprintf("%v", pq) != heapVer {
			t.Errorf("func String has broken format: expected %v got %v", heapVer, fmt.Sprintf("%v", pq))
		}
	})

}
