package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"log"
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
		return nil, errors.New("error while parsing body")
	}

	return data.PoolData, nil
}

func getPool(url string, token string) ([]Pool, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("x-auth-token", token)
	req.Header.Set("Content-Type", "application/json")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	pool := make([]Pool, 0)
	data := jsonData{Embedded{pool}}

	pools, err := render(body, data)

	return pools, err
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

	pools, _ := getPool(url, token)
	for _, pool := range pools {
		fmt.Printf("Id = %v, Name = %v, Status = %v\n", pool.Id, pool.Name, pool.Status)
	}
}
