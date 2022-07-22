package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"sync"
)

var addr = flag.String("addr", "http://localhost:4000", "server's address")

func main() {
	flag.Parse()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getPage()
		}()
	}
	wg.Wait()
}

func getPage() {
	res, err := http.Get(*addr + "?n=40")
	if err != nil {
		return
	}
	var data interface{}
	json.NewDecoder(res.Body).Decode(&data)
	fmt.Println(data)
}
