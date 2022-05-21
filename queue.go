package priority

//generic implementation of priority queue example found at https://pkg.go.dev/container/heap#example-package-PriorityQueue

type Item[T any] struct {
	Value T

	priority int
	index    int
}

type Queue[T any] []*Item[T]

func (q Queue[T]) Len() int {
	return len(q)
}

func (q Queue[T]) Less(i, j int) bool {
	return q[i].priority > q[j].priority
}

func (q Queue[T]) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *Queue[T]) Push(item *Item[T]) {
	n := len(*q)
	item.index = n
	*q = append(*q, item)
}

func (q *Queue[T]) Pop() *Item[T] {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]
	return item
}

func (q *Queue[T]) Update(item *Item[T], value T, priority int) {
	item.Value = value
	item.priority = priority
	Fix[*Item[T]](q, item.index)
}
