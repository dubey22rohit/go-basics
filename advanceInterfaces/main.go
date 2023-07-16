package main

import "fmt"

type Putter interface {
	Put(id int, val any) error
	fmt.Stringer
}

type Storage interface {
	Putter
	Get(id int) (any, error)
}

type simplePutter struct{}

func (s *simplePutter) Put(id int, val any) error {
	return nil
}

func (s *simplePutter) String() string {
	return ""
}

type FooStorage struct{}

func (s *FooStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s *FooStorage) Put(id int, val any) error {
	return nil
}

func (s *FooStorage) String() string {
	return ""
}

type BarStorage struct{}

func (s *BarStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s *BarStorage) Put(id int, val any) error {
	return nil
}

//This way of defining the updateValue method has a drawback that
// This method is dependent upon FooStorage
// func updateValue(id int, val any) error {
// 	store := &FooStorage{}
// 	return store.Put(id, val)
// }

// A better implementation of the above method would be:
// using DI we'll inject the interface Storage in it
// func updateValue(id int, val any, storage Storage) error {
//But why copy the entire storage interface which can have say a 100 methods?
//solution ? Create a new interface putter which only has the Put method
//The storage interface also impements the putter interface
// 	return storage.Put(id, val)
// }

func updateValue(id int, val any, p Putter) error {
	return p.Put(id, val)
}

type Server struct {
	store Storage
}

func main() {
	s := &Server{
		store: &FooStorage{},
	}
	sputter := &simplePutter{}

	updateValue(1, "foo", s.store)
	updateValue(1, "foo", sputter)
}
