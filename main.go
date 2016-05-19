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

//type Context struct {
//	Args   []string
//	Stdout io.Writer
//	Stderr io.Writer
//	Stdin  io.Reader
//}

type Entities struct {
	Id     int
	Name   string
	Status string
}

func parseJson(body []byte, path string) ([]Entities, error) {
	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		return nil, errors.New("error while parsing body")
	}
	embedded, err := jsonParsed.S("_embedded", path).Children()
	if err != nil {
		return nil, errors.New("error while getting entity")
	}

	var entities []Entities
	for _, entity := range embedded {
		data := entity.Data().(map[string]interface{})
		entities = append(entities, Entities{
			Id:     int(data["id"].(float64)),
			Name:   data["name"].(string),
			Status: data["_status"].(string)})
	}

	return entities, nil
}

func getEntity(url string, token string, path string) ([]Entities, error) {
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

func renderTable(entities []Entities) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Status"})
	table.SetAlignment(tablewriter.ALIGN_CENTER)

	for _, entity := range entities {
		table.Append([]string{strconv.Itoa(entity.Id), entity.Name, entity.Status})

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

	entities, err := getEntity(url, token, ePath)
	if err != nil {
		log.Fatal(err)
	}

	//var stdout, stderr bytes.Buffer
	//context := Context{
	//	Stdout: &stdout,
	//	Stderr: &stderr,
	//}
	//renderTable(&context, entities)

	renderTable(entities)
}
