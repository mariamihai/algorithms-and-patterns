package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// https://www.youtube.com/watch?v=B9uR2gLM80E

type User struct {
	Name struct {
		Title string `json:"title"`
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
	Email string `json:"email"`
	Nat   string `json:"nat"`
}

func retrieveRandomUser() interface{} {
	var data struct {
		Results []User `json:"results"`
	}

	url := fmt.Sprintf("https://randomuser.me/api/")

	resp, err := http.Get(url)
	if err != nil {
		return nil
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil
	}

	return data
}

func retrieveRandomUserGo(cnt int, ch chan<- string, wg *sync.WaitGroup) interface{} {
	var data struct {
		Results []User `json:"results"`
	}

	defer wg.Done()

	url := fmt.Sprintf("https://randomuser.me/api/")

	resp, err := http.Get(url)
	if err != nil {
		return nil
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil
	}

	ch <- fmt.Sprintf("This is the %d user", cnt)

	return data
}

func main() {
	startTime := time.Now()

	for i := 0; i < 5; i++ {
		fmt.Println(retrieveRandomUser())
	}

	fmt.Println("Total time: ", time.Since(startTime))
	fmt.Println()

	// ---------------------------------

	startTime = time.Now()

	ch := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go retrieveRandomUserGo(i, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("Total time: ", time.Since(startTime))
}
