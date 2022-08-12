package hilink

import (
	"context"
	"fmt"
	"github.com/jeadie/hilink"
)

// ReceiveSms from a list of endpoints. If a Hilink client cannot be connected to, or Smss cannot
// be retrieved from a device, no error is thrown
func ReceiveSms(endpoints []string) chan Sms {
	output := make(chan Sms)

	go func(output chan Sms, endpoints []string) {
		defer close(output)
		for _, endpoint := range endpoints {
			hiClient, err := CreateHilink(endpoint)
			if err != nil {
				fmt.Println(err)
				continue
			}

			sms, err := hiClient.ListSms()
			if err != nil {
				fmt.Println(err)
			}

			for _, s := range sms {
				output <- s
			}
		}
	}(output, endpoints)
	return output
}

// ListSms from the Hilink device.
func (h *Hilink) ListSms() ([]Sms, error) {
	// See reference: https://github.com/kenshaw/hilink/blob/ccc8a1ffb07ae4ac3368650d840249c2f85238b3/cmd/sms/main.go#L65
	xml, err := h.c.SmsList(context.TODO(), uint(hilink.SmsBoxTypeInbox), 1, 1000, false, false, true)
	if err != nil {
		return []Sms{}, err
	}

	sms := make([]Sms, len(xml))
	i := 0

	//TODO: actually parse SMS xml
	for k, v := range xml {
		sms[i] = Sms{
			From:    k,
			Content: v.(string),
		}
		i++
	}
	return sms, nil
}
