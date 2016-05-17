package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Jeffail/gabs"
	"github.com/olekukonko/tablewriter"
)

func parseJson(body []byte, path string) ([]*gabs.Container, error) {
	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		return nil, errors.New("error while parsing body")
	}
	entities, err := jsonParsed.S("_embedded", path).Children()
	if err != nil {
		return nil, errors.New("error while getting entity")
	}

	return entities, nil
}

func getEntity(url string, token string, path string) ([]*gabs.Container, error) {
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

	entities, err := parseJson(body, path)

	return entities, err
}

func renderTable(entities []*gabs.Container) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Status"})
	table.SetAlignment(tablewriter.ALIGN_CENTRE)

	for _, entity := range entities {
		newData := entity.Data().(map[string]interface{})
		table.Append([]string{strconv.FormatFloat(newData["id"].(float64), 'f', 0, 64), newData["name"].(string), newData["_status"].(string)})
	}
	table.Render()
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
	ePath := os.Args[1]
	url := host + ePath

	}

	renderTable(entities)
}
