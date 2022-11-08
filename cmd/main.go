package main

import (
	"fmt"

	"github.com/usetheplatform/lru-cache/pkg/cache"
)

func main() {

	c := cache.NewCache[string, string](3)
	fmt.Printf("Items are: %v\n", c.Items())

	c.Set("1", "one")
	c.Set("2", "two")
	c.Set("3", "three")

	fmt.Printf("Length is: %d\n", c.Length())
	fmt.Printf("Items are: %v\n", c.Items())

	c.Set("4", "four")

	fmt.Printf("Key 1 is: %v\n", c.Get("1"))
	fmt.Printf("Items are: %v\n", c.Items())
}
