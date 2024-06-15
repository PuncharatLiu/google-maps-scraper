package cref

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var names []string
var cursor string

const mondayAPIURL = "https://api.monday.com/v2"

type RequestBody struct {
	Query string `json:"query"`
}

func GetAllName(w http.ResponseWriter, r *http.Request) {
	// Test
	fmt.Println("call /names")
	doQuery1()
	for cursor != "" {
		doQuery2(cursor)
	}

	// Test
	fmt.Println("names len: ", len(names))
}

func doQuery1() {
	apiKey := getApiKey()

	client := &http.Client{}
	query1 := `
		{
			boards(ids: [4594651101]) {
				items_page(limit: 500) {
					cursor
					items {
						name
					}
				}
			}
		}`

	reqBody := RequestBody{
		Query: query1,
	}

	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalf("Error marshaling request body: %v", err)
	}

	req, err := http.NewRequest("POST", mondayAPIURL, bytes.NewBuffer(jsonReqBody))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", apiKey)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	unmarshalData1(body)
}

func doQuery2(cursor string) {
	apiKey := getApiKey()

	client := &http.Client{}
	query2 := fmt.Sprintf(`
		{
			next_items_page (limit: 500, cursor: "%s") {
				cursor
				items {
					name
				}
			}
		}`, cursor)

	reqBody := RequestBody{
		Query: query2,
	}

	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalf("Error marshaling request body: %v", err)
	}

	req, err := http.NewRequest("POST", mondayAPIURL, bytes.NewBuffer(jsonReqBody))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", apiKey)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer res.Body.Close()

	body2, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	unmarshalData2(body2)
}

func unmarshalData1(body []byte) {
	type Root struct {
		Data struct {
			Boards []struct {
				ItemsPage struct {
					Cursor string `json:"cursor"`
					Items  []struct {
						Name string `json:"name"`
					} `json:"items"`
				} `json:"items_page"`
			} `json:"boards"`
		} `json:"data"`
		AccountID int `json:"account_id"`
	}

	var root Root
	err := json.Unmarshal(body, &root)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	cursor = root.Data.Boards[0].ItemsPage.Cursor
	for _, item := range root.Data.Boards[0].ItemsPage.Items {
		names = append(names, item.Name)
	}

	// Test
	fmt.Println("Names:", len(names))
	// fmt.Println("Names:", names)
	fmt.Println("Cursor:", cursor)
}

func unmarshalData2(body2 []byte) {
	type Root2 struct {
		Data struct {
			NextItemsPage struct {
				Cursor string `json:"cursor"`
				Items  []struct {
					Name string `json:"name"`
				} `json:"items"`
			} `json:"next_items_page"`
		} `json:"data"`
		AccountID int `json:"account_id"`
	}

	var root2 Root2
	err := json.Unmarshal(body2, &root2)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	cursor = root2.Data.NextItemsPage.Cursor
	for _, item := range root2.Data.NextItemsPage.Items {
		names = append(names, item.Name)
	}

	// Test
	fmt.Println("len of names: ", len(names))
	fmt.Println("cursor: ", cursor)
}

func getApiKey() string {
	err := godotenv.Load("../../monday.env")
	if err != nil {
		log.Fatalf("from get_all_name: Error loading .env file: %v", err)
		return ""
	}

	apiKey := os.Getenv("MONDAY_API_TOKEN")
	if apiKey == "" {
		log.Fatalf("MONDAY_API_TOKEN is require")
	}

	return apiKey
}
