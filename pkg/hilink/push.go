package hilink

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// PushSms payloads to an endpoint.
func PushSms(endpoint string, s SmsMessage) {
	fmt.Println("Pushing message", s)
	reqBodyBytes := new(bytes.Buffer)
	err := json.NewEncoder(reqBodyBytes).Encode(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = http.Post(endpoint, "application/json", reqBodyBytes)
	if err != nil {
		fmt.Println(err)
	}
}
