package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	for {
		fmt.Println("making request...")
		var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
		resp, err := http.Post("http://transpose", "application/json", bytes.NewBuffer(jsonStr))
		if err != nil {
			log.Println("err: ", err)
		}
		fmt.Printf("resp: %+v", resp)
		time.Sleep(3 * time.Second)
	}
}
