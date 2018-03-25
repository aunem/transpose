package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	c := http.DefaultClient
	for {
		fmt.Println("making request...")
		var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
		r, err := http.NewRequest(http.MethodPost, "http://transpose:8080", bytes.NewBuffer(jsonStr))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("request: %+v", r)
		resp, err := c.Do(r)
		// resp, err := http.Post("http://transpose:8080", "application/json", bytes.NewBuffer(jsonStr))
		if err != nil {
			log.Println("err: ", err)
		}
		fmt.Printf("resp: %+v", resp)
		time.Sleep(3 * time.Second)
	}
}
