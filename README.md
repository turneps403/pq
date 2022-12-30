# ⑉PrQueue
- generics + comparator-func
- time complexity O(n\*log(n)) as any other priority queue
- [examples](#erremptyqueue)

## Usage

```go
import "github.com/turneps403/prqueue"

pqMin := prqueue.New(func(a, b int) bool { return a < b })
pqMax := prqueue.New(func(a, b int) bool { return a > b })

type custom struct {
    w int
}
pqMin := prqueue.New(func(a, b custom) bool { return a.w < b.w })
pqMax := prqueue.New(func(a, b custom) bool { return a.w > b.w })

pq.Add(el)
el, err := pq.Poll()	// Retrieves and removes
el, err := pq.Peek()	// Retrieves, but does not remove
pq.IsEmpty()		   // bool
pq.Len()			   // bool
```

## Contents

### Methods
- [New(func(T),[capacity])](#new)
- [Add(T)](#add)
- [Poll() T](#poll)
- [Peek() T](#peek)
- [IsEmpty() bool](#isempty)

### Errors
- [ErrEmptyQueue](#erremptyqueue)


## Methods

### New
recieves a **comparator-func** and optional **capacity**.
```go
func New[T any]( func(a, b T) bool, ...int ) *pq[T]
```
e.g.
```go
pq := prqueue.New(1000, func(a, b int) bool { return a < b })
pq := prqueue.New(func(a, b int) bool { return a < b })

pqmin := prqueue.New(func(a, b int) bool { return a < b })
pqmax := prqueue.New(func(a, b int) bool { return a > b })
```

### Add
inserts the specified element into this priority queue. Thread safe.
```go
func (*pq[T]) Add(el T)
```

### Poll
retrieves and removes the head of this queue, or return error ErrEmptyQueue. Thread safe.
```go
func (*pq[T]) Poll() (el T, er error)
```

### Peek
retrieves, but does not remove, the head of this queue, or return error ErrEmptyQueue. Thread safe.
```go
func (*pq[T]) Peek() (el T, er error)
```

### IsEmpty
returns true if this collection contains no elements.
```go
func (*pq[T]) IsEmpty() bool
```

### Len
current length of queue
```go
func (*pq[T]) Len() int
```

## Errors

### ErrEmptyQueue 
indicates about empty queue. contains string "priority queue is empty"