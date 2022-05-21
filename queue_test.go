package priority

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(Queue[string], len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item[string]{
			Value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	Init[*Item[string]](&pq)

	// Insert a new item and then modify its priority.
	item := &Item[string]{
		Value:    "orange",
		priority: 1,
	}
	Push[*Item[string]](&pq, item)

	pq.Update(item, item.Value, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := Pop[*Item[string]](&pq)
		fmt.Printf("%.2d:%s ", item.priority, item.Value)
	}
}
