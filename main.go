package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"errors"
)

type jsonData struct {
	Embedded `json:"_embedded"`
}

type Embedded struct {
	PoolData []Pool `json:"pool"`
}

type Pool struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"_status"`
}

func render(body []byte, data jsonData) ([]Pool, error) {
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, errors.New("Error while parsing body!")
	}

	return data.PoolData, nil
}

func main() {
	host := os.Getenv("GALEB_HOST")
	token := os.Getenv("GALEB_TOKEN")
	if host == "" {
		fmt.Println("GALEB_HOST undefined.")
		return
	}
	if token == "" {
		fmt.Println("GALEB_TOKEN undefined.")
		return
	}
	url := host + os.Args[1]
	pool := make([]Pool, 0)
	data := jsonData{Embedded{pool}}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("x-auth-token", token)
	req.Header.Set("Content-Type", "application/json")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	pools,_ := render(body, data)
	for _, pool := range pools {
		fmt.Printf("Id = %v, Name = %v, Status = %v\n", pool.Id, pool.Name, pool.Status)
	}
}
