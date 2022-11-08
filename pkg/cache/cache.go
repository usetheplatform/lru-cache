package cache

type Cache[K comparable, V any] struct {
	capacity uint
	length   uint
	head     *Node[K, V]
}

type Node[K comparable, V any] struct {
	key     K
	payload V
	prev    *Node[K, V]
	next    *Node[K, V]
}

func NewCache[K comparable, V any](capacity uint) Cache[K, V] {
	return Cache[K, V]{
		capacity: capacity,
		length:   0,
		head:     nil,
	}
}

func (c *Cache[K, V]) Length() uint {
	return c.length
}

func (c *Cache[K, V]) Items() []*Node[K, V] {
	node := c.head

	items := make([]*Node[K, V], 0, c.length)

	for node != nil {
		items = append(items, node)
		node = node.next
	}

	return items
}

func (c *Cache[K, V]) Set(key K, payload V) {
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
		c.head = &Node[K, V]{
			key:     key,
			payload: payload,
			prev:    nil,
			next:    nil,
		}
		c.length += 1
	} else {

		next := c.head

		newNode := Node[K, V]{
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

func (c *Cache[K, V]) Has(key K) bool {
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

func (c *Cache[K, V]) Get(key K) *V {
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
