package hilink

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// response>
//
//	<Count>1</Count>
//	<Messages>
//		<Message>
//			<Smstat>0</Smstat>
//			<Index>40008</Index>
//			<Phone>...</Phone>
//			<Content>Hello</Content>
//			<Date>2022-08-15 12:09:24</Date>
//			<Sca></Sca>
//			<SaveType>4</SaveType>
//			<Priority>0</Priority>
//			<SmsType>1</SmsType>
//		</Message>
//	</Messages>
//
// </response>
type SmsListResponse struct {
	XMLName  xml.Name        `xml:"response"`
	Count    int             `xml:"Count"`
	Messages SmsListMessages `xml:"Messages"`
}

type SmsListMessages struct {
	Message []SmsMessage `xml:"Message"`
}

// <Smstat>0</Smstat>
// <Index>40008</Index>
// <Phone>...</Phone>
// <Content>Hello</Content>
// <Date>2022-08-15 12:09:24</Date>
// <Sca></Sca>
// <SaveType>4</SaveType>
// <Priority>0</Priority>
// <SmsType>1</SmsType>
type SmsMessage struct {
	XMLName  xml.Name `xml:"Message"`
	Smstat   uint     `xml:"Smstat"`
	Index    uint     `xml:"Index"`
	Phone    string   `xml:"Phone"`
	Content  string   `xml:"Content"`
	Date     string   `xml:"Date"`
	Sca      any      `xml:"Sca"`
	SaveType uint     `xml:"SaveType"`
	Priority uint     `xml:"Priority"`
	SmsType  uint     `xml:"SmsType"`
}

// <response>
// <SesInfo>SessionID=4Xr6pqD9k5i8bTSX32YGrVGirdhK7zyNYFVjzP38q8/0JTxgAmOKfDpMJ8lmeJbstBJIvR7JLU5wc7zejlpn8kpuzRsh/oajMHwMklaXrF3RKiUTU5v4tI9tAZxVL6tm</SesInfo>
// <TokInfo>ttws6CQ9WT5EftzoWYS6yVCEfRNK3BhE</TokInfo>
// </response>
type SesTokInfo struct {
	XMLName xml.Name `xml:"response"`
	SesInfo string   `xml:"SesInfo"`
	TokInfo string   `xml:"TokInfo"`
}

func GetSesTokInfo(url string) SesTokInfo {
	url = fmt.Sprintf(`%s/api/webserver/SesTokInfo`, url)
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)

	var stInfo SesTokInfo

	err := xml.Unmarshal(body, &stInfo)
	if err != nil {
		fmt.Println(err)
		return SesTokInfo{}
	}
	return stInfo
}

func GetHeaders(url string) map[string][]string {
	stInfo := GetSesTokInfo(url)
	header := make(map[string][]string)
	header["Content-Type"] = []string{"text/xml; charset=UTF-8"}
	header["Cookie"] = []string{stInfo.SesInfo}
	header["__RequestVerificationToken"] = []string{stInfo.TokInfo}
	return header

}

func ParseResponse(resp *http.Response) SmsListResponse {
	body, _ := ioutil.ReadAll(resp.Body)
	var smsList SmsListResponse

	err := xml.Unmarshal(body, &smsList)
	if err != nil {
		fmt.Println(err)
		return SmsListResponse{Count: 0}
	}
	return smsList
}

func GetSmsList(url string) []SmsMessage {
	nb := 50
	postData := fmt.Sprintf(`<?xml version = "1.0" encoding = "UTF-8"?>%c<request><PageIndex>%d</PageIndex><ReadCount>%d</ReadCount><BoxType>%d</BoxType><SortType>%d</SortType><Ascending>%d</Ascending><UnreadPreferred>%d</UnreadPreferred></request>`, '\n', 1, nb, 1, 0, 0, 1)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(`%s/api/sms/sms-list`, url), bytes.NewBufferString(postData))
	if err != nil {
		fmt.Println(err)
		return []SmsMessage{}
	}
	req.Header = GetHeaders(url)
	c := &http.Client{Timeout: 10 * time.Second}
	resp, err := c.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return []SmsMessage{}
	}
	smsList := ParseResponse(resp)

	return smsList.Messages.Message
}
