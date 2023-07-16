package main

import "fmt"

//Stor(er)
// io.Red(er)
// io.Writ(er)

type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

type PostgresNumberStore struct {
	//postgres values (db connection)
}

func (s PostgresNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3, 4, 5, 6, 7}, nil
}

func (s PostgresNumberStore) Put(number int) error {
	fmt.Println("store the number into the mongoDB storage", number)
	return nil
}

type MongoDBNumberStore struct {
	//some values
	//To implement an interface a struct / value of any type must implement the methods
	//mentioned in the interface
}

func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3, 4}, nil
}

func (m MongoDBNumberStore) Put(number int) error {
	fmt.Println("store the number into the mongoDB storage", number)
	return nil
}

type ApiServer struct {
	numberStore NumberStorer
}

func main() {
	//Since we had a generic number store, it is easy to swap it from mongo to postgres.
	//This is the power of interfaces
	apiServer := ApiServer{
		// numberStore: MongoDBNumberStore{},
		numberStore: PostgresNumberStore{},
	}

	//Logic
	if err := apiServer.numberStore.Put(10); err != nil {
		panic(err)
	}

	numbers, err := apiServer.numberStore.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(numbers)
}
