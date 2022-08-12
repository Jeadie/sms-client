package hilink

import (
	"context"
	"fmt"
	"github.com/jeadie/hilink"
	"os"
)

// ReceiveSms from a list of endpoints.
func ReceiveSms(endpoints []string) {
	for _, endpoint := range endpoints {
		hiClient, err := CreateHilink(endpoint)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sms, err := hiClient.ListSms()
		if err != nil {
			fmt.Println(err)
		}
		for _, s := range sms {
			fmt.Println(s)
		}
	}
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
