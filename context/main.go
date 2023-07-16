package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func thirdPartyHTTPCall() (string, error) {
	time.Sleep(time.Millisecond * 90)
	return "user id: 200", nil
}

func fetchUserID(ctx context.Context) (string, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()

	val := ctx.Value("username")
	fmt.Println("fetching for user: ", val)

	type result struct {
		userId string
		err    error
	}

	resultch := make(chan result, 1)

	go func() {
		res, err := thirdPartyHTTPCall()
		resultch <- result{
			userId: res,
			err:    err,
		}
	}()

	select {
	case <-ctx.Done():
		//1) The context timeout is exceeded
		//2) The context has been manually cancelled -> cancel()
		return "", ctx.Err()
	case res := <-resultch:
		return res.userId, res.err
	}
}

func main() {
	start := time.Now()
	// res, err := thirdPartyHTTPCall()
	//* Goroutines are susceptible to race conditions, but passing values from context is safe.
	ctx := context.WithValue(context.Background(), "username", "rohitdubey22")
	userID, err := fetchUserID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the response took %v -> %+v", time.Since(start), userID)
}
