package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://www.thepaper.cn/"
	response, err := http.Get(url)

	if err != nil {
		log.Printf("fetch url error:%v", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("Error status code:/%v", err)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("read content failed:%v", err)
		return
	}

	log.Println("body:", string(body))
}
