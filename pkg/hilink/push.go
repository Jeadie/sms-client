package hilink

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// PushSms payloads to an endpoint.
func PushSms(endpoint string, sms []SmsMessage) {
	fmt.Printf("Pushing %d smses to %s\n", len(sms), endpoint)
	for _, s := range sms {
		reqBodyBytes := new(bytes.Buffer)
		err := json.NewEncoder(reqBodyBytes).Encode(s)
		if err != nil {
			fmt.Println(err)
			continue
		}
		_, err = http.Post(endpoint, "application/json", reqBodyBytes)
		if err != nil {
			fmt.Println(err)
		}
	}
}
