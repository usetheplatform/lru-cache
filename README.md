# LRU Cache

## Get Started

### Create a cache

```go
// Initialize a cache of size 3 that can hold payload of type V associated with key of type K
// where:
//  K = string
//  V = string
c := cache.NewCache[string, string](3)
```

### Store a value in the cache

```go
c.Set("key", "value")
```

### Retreive a value from the cache

```go
c.Get("key") // type *V or nil, if not found
```

### Some other supported operations, that may be helpful

#### Check if there is such key in the cache

```go
c.Has("key") // boolean
```

#### Check the current length of the cache

```go
c.Length() // uint
```

#### Turn the list into an array of type \*Node[K, V]

```go
c.Items() // []*Node[K, V]
```

### Internal types

```go
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

// Type Constructor
func NewCache[K comparable, V any](capacity uint) Cache[K, V] {
	return Cache[K, V]{
		capacity: capacity,
		length:   0,
		head:     nil,
	}
}
```
