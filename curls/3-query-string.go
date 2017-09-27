package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	request := `
	    {
			"size": 400,
	        "query": {
	            "query_string": {
					"default_field": "alliance.name",
					"query": "Garg* AND *cker"
				}
	        }
	    }
	`

	body := strings.NewReader(request)
	req, err := http.NewRequest("GET", "http://0.0.0.0:9222/socialpoint/players/_search", body)
	if err != nil {
		log.Printf("#%v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("#%v", err)
	}
	defer resp.Body.Close()

	out, err := os.Create("result.json")
	if err != nil {
		log.Printf("#%v", err)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Printf("#%v", err)
	}
}
