package hilink

import (
	"fmt"
	"time"
)

func Poll(endpoints []string, periodSecs uint, pktReconstruct bool) chan SmsMessage {
	smsOutput := make(chan SmsMessage, 10)

	go func(endpoints []string, periodSecs uint, output chan SmsMessage) {
		lastTimestamp := "1970-01-01 12:00:00"
		for {
			time.Sleep(time.Second * time.Duration(periodSecs))
			fmt.Println("Awake!")

			for _, e := range endpoints {
				localMostRecentTimestamp := "1970-01-01 12:00:00"
				smsList := GetSmsList(e)
				if pktReconstruct {
					smsList = ReconstructSms(smsList)
				}
				for _, m := range smsList {
					// Update local last sms timestamp
					if m.Date > localMostRecentTimestamp {
						localMostRecentTimestamp = m.Date
					}
					if m.Date > lastTimestamp {
						smsOutput <- m
					}
				}
				fmt.Println(localMostRecentTimestamp, lastTimestamp)
				if localMostRecentTimestamp > lastTimestamp {
					lastTimestamp = localMostRecentTimestamp
				}
			}
		}
	}(endpoints, periodSecs, smsOutput)

	return smsOutput
}
