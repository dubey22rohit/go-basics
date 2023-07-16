package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type UserProfile struct {
	ID       int
	Comments []string
	Likes    int
	Friends  []int
}

type Response struct {
	data any
	err  error
}

func handleGetUserProfile(id int) (*UserProfile, error) {
	var (
		respch = make(chan Response, 3)
		wg     = &sync.WaitGroup{}
	)

	go getComments(id, respch, wg)
	go getLikes(id, respch, wg)
	go getFriends(id, respch, wg)

	wg.Add(3)
	wg.Wait() //* When counter is 0 we unblock
	close(respch)

	userProfile := &UserProfile{}

	for resp := range respch {
		if resp.err != nil {
			return nil, resp.err
		}
		switch msg := resp.data.(type) {
		//can be a better check, instead of checking the data type
		case int:
			userProfile.Likes = msg
		case []int:
			userProfile.Friends = msg
		case []string:
			userProfile.Comments = msg
		}
	}
	return userProfile, nil
}

func getComments(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	comments := []string{"Wow cool!", "looking good!"}

	respch <- Response{
		data: comments,
		err:  nil,
	}
	wg.Done()
}

func getLikes(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	respch <- Response{
		data: 100,
		err:  nil,
	}
	wg.Done()
}

func getFriends(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	friends := []int{123, 456, 789}

	respch <- Response{
		data: friends,
		err:  nil,
	}
	wg.Done()
}

func main() {
	start := time.Now()
	userProfile, err := handleGetUserProfile(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("time took: %v to get the userProfile: %+v", time.Since(start), userProfile)
}
