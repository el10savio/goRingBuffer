package buffer

type BasicRingBuffer[T any] struct {
	capacity uint
	array    []*T
	head     int
	tail     int
}

func NewBasicRingBuffer[T any](capacity uint) *BasicRingBuffer[T] {
	buffer := &BasicRingBuffer[T]{
		capacity: capacity,
		array:    make([]*T, capacity),
		head:     0,
		tail:     0,
	}
	return buffer
}

func (rb *BasicRingBuffer[T]) GetCapacity() uint {
	return rb.capacity
}

func (rb *BasicRingBuffer[T]) Read() *T {
	element := rb.array[rb.tail]
	if rb.tail == rb.head {
		return element
	}

	rb.tail = (rb.tail + 1) % int(rb.capacity)
	return element
}

func (rb *BasicRingBuffer[T]) Write(element *T) {
	rb.array[rb.head] = element
	rb.head = (rb.head + 1) % int(rb.capacity)
}
