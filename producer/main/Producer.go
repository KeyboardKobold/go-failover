package main

import (
	"net/url"
	"io/ioutil"
	"fmt"
	"strconv"
	"time"
	"net/http"
	"strings"
)

func main() {
	count := 1
	ticker := time.NewTicker(1000 * time.Millisecond)
	for range ticker.C {
		postEvent(count)
		count = count + 1
	}
}

func postEvent(count int){
	requestUrl := "http://localhost:8080/addEvent"
	form := url.Values{
		"eventId": {strconv.Itoa(count)},
	}
	fmt.Printf(strconv.Itoa(count)+ "\n")
	fmt.Printf(form.Get("eventId") + "\n")

	req, err := http.NewRequest("POST", requestUrl, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer rsp.Body.Close()
	bodyByte, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bodyByte))
}
