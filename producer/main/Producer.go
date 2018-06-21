package main

import (
	"net/url"
	"io/ioutil"
	"fmt"
	"strconv"
	"time"
	"net/http"
	"strings"
	"os"
	"errors"
)

func main() {
	count := 1
	ticker := time.NewTicker(1000 * time.Millisecond)
	for range ticker.C {
		err := postEvent(count)
		if err != nil {
			fmt.Println(err)
		} else {
			count = count + 1
		}
	}
}

func postEvent(count int) error {
	requestHost := os.Getenv("CONSUMER_HOST")
	if requestHost == "" {
		panic(errors.New("CONSUMER_HOST was not set"))
	}
	requestUrl := "http://" + requestHost + "/addEvent"
	form := url.Values{
		"eventId": {strconv.Itoa(count)},
	}
	fmt.Printf(strconv.Itoa(count) + "\n")
	fmt.Printf(form.Get("eventId") + "\n")

	req, err := http.NewRequest("POST", requestUrl, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		//panic(err)
		return err
	}

	defer rsp.Body.Close()
	bodyByte, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bodyByte))
	return nil
}
