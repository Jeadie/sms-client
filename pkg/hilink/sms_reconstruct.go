package hilink

import (
	"fmt"
	"sort"
	"time"
)

func GetSmsTime(x SmsMessage) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", x.Date)
}

func CombineSmss(x []SmsMessage) SmsMessage {
	if len(x) == 0 {
		return SmsMessage{}
	}

	sort.Slice(x, func(i, j int) bool {
		return x[i].Index < x[j].Index
	})

	for _, sms := range x[1:] {
		x[0].Content += sms.Content
	}
	return x[0]
}

func ReconstructSms(x []SmsMessage) []SmsMessage {
	sort.Slice(x, func(i, j int) bool {
		return x[i].Date < x[j].Date
	})

	var result []SmsMessage

	for i := 0; i < len(x); i++ {
		sms := x[i]

		d, err := GetSmsTime(sms)
		if err != nil {
			fmt.Println(fmt.Errorf("[ReconstructSms]: Could not parse date %s for SMS. Error %w\n", sms.Date, err))
			continue
		}

		// Get groups of SMSs within a 1sec time frame
		d1 := d
		j := i
		for d1.Sub(d).Seconds() <= 1 && j < len(x) && sms.Phone == x[j].Phone {
			j++
			if j < len(x) {
				d1, err = GetSmsTime(x[j])
			}
		}

		// Group detected
		if j > i { // Account for initial j++
			result = append(result, CombineSmss(x[i:j]))
			i = j - 1

		} else {
			result = append(result, sms)
		}
	}
	return result
}
