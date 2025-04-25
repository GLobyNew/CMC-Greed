package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	cmcGreedURL = "https://pro-api.coinmarketcap.com/v3/fear-and-greed/latest"
	YEKT_TZ     = "Asia/Yekaterinburg"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", cmcGreedURL, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	req.Header.Set("Accepts", "application/json")
	err = godotenv.Load()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	req.Header.Add("X-CMC_PRO_API_KEY", os.Getenv("CMC_API_KEY"))
	req.Header.Add("limit", "1")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	greedData := Greed{}
	respBody, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(respBody, &greedData)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	log.Printf("CMC Greed value: %v\n", greedData.Data.Value)
	if greedData.Data.Value >= 60 || greedData.Data.Value <= 23 {
		greedData.sendData()
	}
}
