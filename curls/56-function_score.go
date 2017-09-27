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
				"function_score": {
					"functions": [
						{
							"filter": {
								"range": {
									"joined_at": {
										"gte": "now-3M"
									}
								}
							},
							"weight": 10
						},
						{
							"field_value_factor": {
								"field": "alliance.active_members",
								"factor": 1,
								"modifier": "none",
								"missing": 1
							}
						},
						{
							"random_score": {},
							"weight": 1
						}
					],
					"score_mode": "sum",
					"boost_mode":"sum",
					"max_boost": 1000
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
