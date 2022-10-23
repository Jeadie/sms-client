package hilink

// ReceiveSms from a list of endpoints. If a Hilink client cannot be connected to, or Smss cannot
// be retrieved from a device, no error is thrown. if `pktReconstruct` is true, ReceiveSms will
// attempt to remediate SMS packet issues, combining SMSs that should be together
func ReceiveSms(endpoints []string, pktReconstruct bool) chan SmsMessage {
	output := make(chan SmsMessage)

	go func(output chan SmsMessage, endpoints []string) {
		defer close(output)
		for _, endpoint := range endpoints {
			smss := GetSmsList(endpoint)
			if pktReconstruct {
				smss = ReconstructSms(smss)
			}
			for _, s := range smss {
				output <- s
			}
		}
	}(output, endpoints)
	return output
}
