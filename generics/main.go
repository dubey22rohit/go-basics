package main

import "fmt"

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{
		data: make(map[K]V),
	}
}

func (m *CustomMap[K, V]) Insert(k K, v V) error {
	m.data[k] = v
	return nil
}

func foo[T any](val T) {
	fmt.Println(val)
}

func main() {
	m1 := NewCustomMap[string, int]()
	m1.Insert("foo", 10)
	m1.Insert("bar", 20)

	m2 := NewCustomMap[int, float64]()
	m2.Insert(10, 10.99)
	m2.Insert(69, 69.69)

	fmt.Println(m1)
	fmt.Println(m2)

	foo[int](10)
	foo("generics") //If type not provided it will be infered
}
