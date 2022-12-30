# ⑉PrQueue
#### Package prqueue implements a priority queue with **generics** and ability to send **comparator-func** as a constructor argument.

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
pq.IsEmpty()		    // bool
pq.Len()			    // bool
```
