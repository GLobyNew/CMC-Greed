package main

import (
	"fmt"
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

func (g Greed) printData() {
	yektTime, _ := time.LoadLocation(YEKT_TZ)
	time := g.Data.UpdateTime.In(yektTime).Format(time.DateTime)
	fmt.Printf("%s - %d - %s\n", time, g.Data.Value, g.Data.ValueClassification)
}
