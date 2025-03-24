package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Greed struct {
	Data struct {
		Value               int       `json:"value"`
		UpdateTime          time.Time `json:"update_time"`
		ValueClassification string    `json:"value_classification"`
	} `json:"data"`
	Status struct {
		Timestamp    time.Time `json:"timestamp"`
		ErrorCode    string    `json:"error_code"`
		ErrorMessage string    `json:"error_message"`
		Elapsed      int       `json:"elapsed"`
		CreditCount  int       `json:"credit_count"`
	} `json:"status"`
}

func (g Greed) generateResult() string {
	yektTime, _ := time.LoadLocation(YEKT_TZ)
	time := g.Data.UpdateTime.In(yektTime).Format(time.DateTime)
	return fmt.Sprintf("%s - %d - %s\n", time, g.Data.Value, g.Data.ValueClassification)
}

func (g Greed) sendData() {
	type telegramMessage struct {
		Message string `json:"message"`
	}
	msg := telegramMessage{
		Message: g.generateResult(),
	}
	js, err := json.Marshal(msg)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	resp, err := http.Post(os.Getenv("WEBHOOK_URL"), "application/json", bytes.NewBuffer(js))
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Print("there was an error while sending post req")
	}

}
