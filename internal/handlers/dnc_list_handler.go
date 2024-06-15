package handlers

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

const mondayAPIURL = "https://api.monday.com/v2"

type RequestBody struct {
	Query string `json:"query"`
}

type Data struct {
	Data struct {
		Boards []struct {
			ItemsPage struct {
				Items []struct {
					Name string `json:"name"`
				} `json:"items"`
			} `json:"items_page"`
		} `json:"boards"`
	} `json:"data"`
}

func GetDNCList(w http.ResponseWriter, r *http.Request) {
	// Load environment variable from .env file
	err := godotenv.Load("monday.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	apiToken := os.Getenv("MONDAY_API_TOKEN")
	if apiToken == "" {
		log.Fatalf("MONDAY_API_TOKEN is required")
	}

	client := &http.Client{}
	query := `
	{
		boards(ids: [4609772656]) {
			items_page(limit: 500){
				items{
					name
				}
			}
		}
	}`

	reqBody := RequestBody{
		Query: query,
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
	req.Header.Set("Authorization", apiToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Decode the response body into a struct
	var response Data
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("Error unmarshaling response body: %v", err)
	}

	fmt.Println("Unmarshalled response: ", response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

	// fmt.Println("Response: ", string(body))
	fmt.Println("GetDNCList called!")
}
