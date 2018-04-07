package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	c := http.DefaultClient
	fmt.Println("connecting to service...")
	err := retry(30, 3*time.Second, func() error {
		err0 := makeRequest(c)
		return err0
	})
	if err != nil {
		log.Fatal("could not connect to service: ", err)
	}
	fmt.Println("connected to service!")
	fmt.Println("testing...")
	for i := 0; i < 10; i++ {
		err := makeRequest(c)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
	}
	os.Exit(0)
}

func makeRequest(c *http.Client) error {
	fmt.Println("making request...")
	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	r, err := http.NewRequest(http.MethodPost, "http://transpose:8080", bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	log.Printf("request: %+v", r)
	resp, err := c.Do(r)
	if err != nil {
		return err
	}
	fmt.Printf("resp: %+v", resp)
	return nil
}

func retry(attempts int, sleep time.Duration, callback func() error) (err error) {
	for i := 0; ; i++ {
		err = callback()
		if err == nil {
			return
		}

		if i >= (attempts - 1) {
			break
		}

		time.Sleep(sleep)

		log.Println("retrying after error:", err)
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}
