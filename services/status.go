package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	for {
		time.Sleep(15 * time.Second)
		getStatus()
	}
}

func getStatus() {
	var (
		waterStatus string
		windStatus  string
	)
	statusNumber := map[string]int{}

	statusNumber["water"] = rand.Intn(100)
	statusNumber["air"] = rand.Intn(100)

	if statusNumber["water"] <= 5 {
		waterStatus = "safe"
	}

	if statusNumber["water"] >= 6 && statusNumber["water"] <= 8 {
		waterStatus = "standby"
	}

	if statusNumber["water"] > 8 {
		waterStatus = "danger"
	}

	if statusNumber["air"] <= 6 {
		windStatus = "safe"
	}

	if statusNumber["air"] >= 7 && statusNumber["air"] <= 15 {
		windStatus = "standby"
	}

	if statusNumber["air"] > 15 {
		windStatus = "danger"
	}

	fmt.Println(statusNumber)
	fmt.Printf("Water status : %s \n", waterStatus)
	fmt.Printf("Wind status : %s \n\n\n", windStatus)

	data := map[string]interface{}{
		"water": waterStatus,
		"air":   windStatus,
	}

	requestJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("PUT", "http://localhost:8080/save/1", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
