package priority

type Item[T any] struct {
	value    T
	priority int
	index    int
}

type Queue[T any] []*Item[T]

func (q Queue[T]) Get() (*Item[T], error) {
	return nil, nil
}

func (q Queue[T]) Add(i *Item[T]) error {
	return nil
}
