package cache

type Cache[T any] struct {
	capacity uint
	length   uint
	head     *Node[T]
}

type Node[T any] struct {
	key     string
	payload T
	prev    *Node[T]
	next    *Node[T]
}

func NewCache[T any](capacity uint) Cache[T] {
	return Cache[T]{
		capacity: capacity,
		length:   0,
		head:     nil,
	}
}

func (c *Cache[T]) Length() uint {
	return c.length
}

func (c *Cache[T]) Items() []*Node[T] {
	node := c.head

	items := make([]*Node[T], 0, c.length)

	for node != nil {
		items = append(items, node)
		node = node.next
	}

	return items
}

func (c *Cache[T]) Set(key string, payload T) {
	if c.length >= c.capacity {
		node := c.head

		for node != nil && node.next != nil {
			node = node.next
		}

		tail := node

		tail.prev.next = nil
		c.length -= 1
	}

	if c.head == nil {
		c.head = &Node[T]{
			key:     key,
			payload: payload,
			prev:    nil,
			next:    nil,
		}
		c.length += 1
	} else {

		next := c.head

		newNode := Node[T]{
			key:     key,
			payload: payload,
			prev:    nil,
			next:    next,
		}

		next.prev = &newNode

		c.head = &newNode
		c.length += 1
	}

}

func (c *Cache[T]) Has(key string) bool {
	node := c.head

	for node != nil {
		if node.key == key {
			return true
		} else {
			node = node.next
		}
	}

	return false
}

func (c *Cache[T]) Get(key string) *T {
	node := c.head

	for node != nil {
		if node.key == key {
			next := c.head

			if next.next == node {
				next.next = node.next
			}

			if node.prev != nil {
				node.prev.next = node.next
			}

			c.head = node
			c.head.prev = nil
			c.head.next = next

			next.prev = c.head

			return &node.payload
		} else {
			node = node.next
		}
	}

	return nil
}
