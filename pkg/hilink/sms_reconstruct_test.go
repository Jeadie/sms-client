package hilink

import (
	"encoding/xml"
	"fmt"
	"testing"
)

// <Smstat>0</Smstat>
// <Index>40008</Index>
// <Phone>...</Phone>
// <Content>Hello</Content>
// <Date>2022-08-15 12:09:24</Date>
// <Sca></Sca>
// <SaveType>4</SaveType>
// <Priority>0</Priority>
// <SmsType>1</SmsType>
func TestReconstructSms_twoSMSs_shouldCombine(t *testing.T) {
	msgs := []SmsMessage{{
		XMLName:  xml.Name{},
		Smstat:   0,
		Index:    40008,
		Phone:    "12345678",
		Content:  "Hello",
		Date:     "2022-08-15 12:09:24",
		Sca:      nil,
		SaveType: 4,
		Priority: 0,
		SmsType:  1,
	}, {
		XMLName:  xml.Name{},
		Smstat:   0,
		Index:    40009,
		Phone:    "12345678",
		Content:  " World",
		Date:     "2022-08-15 12:09:23",
		Sca:      nil,
		SaveType: 4,
		Priority: 0,
		SmsType:  1,
	}, {
		XMLName:  xml.Name{},
		Smstat:   0,
		Index:    40009,
		Phone:    "12345678",
		Content:  " Random message",
		Date:     "2022-08-15 12:19:23",
		Sca:      nil,
		SaveType: 4,
		Priority: 0,
		SmsType:  1,
	}}
	actual := ReconstructSms(msgs)
	if len(actual) != 2 {
		fmt.Println(len(actual), actual)
		t.Errorf("Messages were not combined, but they should be.")
	}
	if actual[0].Content != "Hello World" {
		t.Errorf("Combined message was incorrect. Expected 'Hello World', found %s", actual[0].Content)
	}
}

func TestReconstructSms_twoSMSs_lastPosition_shouldCombine(t *testing.T) {
	msgs := []SmsMessage{{
		XMLName:  xml.Name{},
		Smstat:   0,
		Index:    40008,
		Phone:    "12345678",
		Content:  "Hello",
		Date:     "2022-08-15 12:09:24",
		Sca:      nil,
		SaveType: 4,
		Priority: 0,
		SmsType:  1,
	}, {
		XMLName:  xml.Name{},
		Smstat:   0,
		Index:    40009,
		Phone:    "12345678",
		Content:  " World",
		Date:     "2022-08-15 12:09:23",
		Sca:      nil,
		SaveType: 4,
		Priority: 0,
		SmsType:  1,
	}}
	actual := ReconstructSms(msgs)
	if len(actual) != 1 {
		fmt.Println(actual)
		t.Errorf("Messages were not combined, but they should be.")
	}
	if actual[0].Content != "Hello World" {
		t.Errorf("Combined message was incorrect. Expected 'Hello World', found %s", actual[0].Content)
	}
}

func TestReconstructSms_twoSMSs_shouldntCombine_timestamp(t *testing.T) {
	msgs := []SmsMessage{{
		XMLName:  xml.Name{},
		Smstat:   0,
		Index:    40008,
		Phone:    "12345678",
		Content:  "Hello",
		Date:     "2022-08-15 12:09:24",
		Sca:      nil,
		SaveType: 4,
		Priority: 0,
		SmsType:  1,
	}, {
		XMLName:  xml.Name{},
		Smstat:   0,
		Index:    40009,
		Phone:    "12345678",
		Content:  " World",
		Date:     "2022-08-15 12:09:26",
		Sca:      nil,
		SaveType: 4,
		Priority: 0,
		SmsType:  1,
	}, {
		XMLName:  xml.Name{},
		Smstat:   0,
		Index:    40009,
		Phone:    "12345678",
		Content:  " Random message",
		Date:     "2022-08-15 12:11:26",
		Sca:      nil,
		SaveType: 4,
		Priority: 0,
		SmsType:  1,
	}}
	actual := ReconstructSms(msgs)
	if len(actual) != 3 {
		t.Errorf("Messages should not be combined (timestamps too far apart)")
	}
}

func TestReconstructSms_twoSMSs_shouldntCombine_differentNumbers(t *testing.T) {
	msgs := []SmsMessage{{
		XMLName:  xml.Name{},
		Smstat:   0,
		Index:    40008,
		Phone:    "12345678",
		Content:  "Hello",
		Date:     "2022-08-15 12:09:24",
		Sca:      nil,
		SaveType: 4,
		Priority: 0,
		SmsType:  1,
	}, {
		XMLName:  xml.Name{},
		Smstat:   0,
		Index:    40009,
		Phone:    "12345679",
		Content:  " World",
		Date:     "2022-08-15 12:09:23",
		Sca:      nil,
		SaveType: 4,
		Priority: 0,
		SmsType:  1,
	}}
	actual := ReconstructSms(msgs)
	if len(actual) == 1 {
		t.Errorf("Messages should not be combined (different phone numbers)")
	}
}
