package cmd

import (
	"os"
	"strconv"
	"strings"
)

const POLL_ENDPOINTS_FLAG = "poll-endpoints"
const PUSH_ENDPOINTS_FLAG = "push-endpoints"
const POLL_PERIOD_FLAG = "poll-period"
const RECONSTRUCT_PACKETS_FLAG = "reconstruct-packets"

const POLL_ENDPOINTS_ENV_KEY = "POLL_ENDPOINTS"
const PUSH_ENDPOINTS_ENV_KEY = "PUSH_ENDPOINTS"
const POLL_PERIOD_ENV_KEY = "POLL_PERIOD"
const RECONSTRUCT_PACKETS_ENV_KEY = "RECONSTRUCT_PACKETS"

func GetDefaultReconstructPackets() bool {
	v, exists := os.LookupEnv(RECONSTRUCT_PACKETS_ENV_KEY)
	if !exists {
		return false
	}
	return v == "true" || v == "True" || v == "T"
}

func GetDefaultPollEndpoints() []string {
	v, exists := os.LookupEnv(POLL_ENDPOINTS_ENV_KEY)
	if !exists {
		return []string{"http://192.168.8.1"}
	}
	return strings.Split(v, ",")
}

func GetDefaultPushEndpoints() []string {
	v, exists := os.LookupEnv(PUSH_ENDPOINTS_ENV_KEY)
	if !exists {
		return []string{"127.0.0.1"}
	}
	return strings.Split(v, ",")
}

func GetDefaultPollPeriod() uint {
	v, exists := os.LookupEnv(POLL_PERIOD_ENV_KEY)
	if exists {
		period, err := strconv.Atoi(v)
		if err == nil {
			return uint(period)
		}
	}
	return 10
}
