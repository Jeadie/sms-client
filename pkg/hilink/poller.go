package hilink

import (
	"fmt"
	"time"
)

func Poll(endpoints []string, periodSecs uint, callback func(endpoint string, sms []SmsMessage)) {
	for {
		time.Sleep(time.Second) //  * time.Duration(periodSecs))
		fmt.Println("Awake!")
		for _, e := range endpoints {
			callback(e, GetSmsList(e))
		}
	}
}
