package hilink

import (
	"fmt"
	"time"
)

func Poll(endpoints []string, periodSecs uint, callback func(endpoint string, sms []Sms)) {
	for {
		time.Sleep(time.Second) //  * time.Duration(periodSecs))
		fmt.Println("Awake!")
		for _, e := range endpoints {
			c, err := CreateHilink(e)
			if err != nil {
				fmt.Println(err)
				continue
			}
			sms, err := c.ListSms()
			if err != nil {
				fmt.Println(err)
				continue
			}

			callback(e, sms)
		}
	}
}
