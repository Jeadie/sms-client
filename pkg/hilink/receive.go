package hilink

// ReceiveSms from a list of endpoints. If a Hilink client cannot be connected to, or Smss cannot
// be retrieved from a device, no error is thrown
func ReceiveSms(endpoints []string) chan SmsMessage {
	output := make(chan SmsMessage)

	go func(output chan SmsMessage, endpoints []string) {
		defer close(output)
		for _, endpoint := range endpoints {
			for _, s := range GetSmsList(endpoint) {
				output <- s
			}
		}
	}(output, endpoints)
	return output
}
