// This example demonstrates a min priority queue with custom struct.
package prqueue_test

import (
	"fmt"

	"github.com/turneps403/prqueue"
)

type Employee struct {
	salary int
	name   string
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func Example_priorityQueue() {
	ems := []Employee{
		{salary: 100, name: "Mike"},
		{salary: 200, name: "Raph"},
		{salary: 250, name: "Donni"},
		{salary: 150, name: "Leo"},
	}
	pq := prqueue.New(func(a, b Employee) bool { return a.salary < b.salary })
	for _, el := range ems {
		pq.Add(el)
	}

	// get a size
	fmt.Printf("Queue contains now: %d element\n", pq.Len())

	// check a Peek value
	p, _ := pq.Peek()
	fmt.Printf("At the Peek now: %v with salary %d\n", p.name, p.salary)

	for !pq.IsEmpty() {
		el, _ := pq.Poll()
		fmt.Printf("Poll returned: %v with salary %d\n", el.name, el.salary)
	}

	if _, err := pq.Peek(); err != nil {
		fmt.Printf("Now queue is empty: %v\n", err)
	}

	// Output:
	// Queue contains now: 4 element
	// At the Peek now: Mike with salary 100
	// Poll returned: Mike with salary 100
	// Poll returned: Leo with salary 150
	// Poll returned: Raph with salary 200
	// Poll returned: Donni with salary 250
	// Now queue is empty: priority queue is empty
}
