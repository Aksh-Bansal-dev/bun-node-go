package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	test("Golang", 4000)
	test("Bun", 4001)
	test("Nodejs(express)", 4002)
}

func test(name string, port int) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getPage(fmt.Sprintf("http://localhost:%d", port))
		}()
	}
	wg.Wait()
	fmt.Printf("%s: %v\n", name, time.Since(start))
}

func getPage(addr string) {
	res, err := http.Get(addr)
	if err != nil {
		return
	}
	var data interface{}
	json.NewDecoder(res.Body).Decode(&data)
}
