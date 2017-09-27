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
	            "bool": {
					"must": {
						"range": {
							"level": {
								"gte": 10,
								"lte": 20
							}
						}
					},
					"should": [
						{
							"range": {
								"joined_at": {
									"gte": "now-3M",
									"boost": 5
								}
							}
						},
						{
							"wildcard": {
								"alliance.name": {
									"value": "*on*",
									"boost": 2
								}
							}
						}
					],
					"must_not": {
						"match": {
							"level": 15
						}
					}
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
