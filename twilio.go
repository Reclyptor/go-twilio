package twilio

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Twilio struct {
	BaseURL string
	SID     string
	Token   string
	Number  string
}

func NewClient(sid string, token string, number string) Twilio {
	return Twilio {
		BaseURL: "https://api.twilio.com/2010-04-01",
		SID: sid,
		Token: token,
		Number: number,
	}
}

func (twilio Twilio) SendSMS(number string, message string) {
	endpoint := fmt.Sprintf("%s/Accounts/%s/Messages.json", twilio.BaseURL, twilio.SID)

	v := url.Values{}
	v.Set("From", twilio.Number)
	v.Set("To", number)
	v.Set("Body", message)
	rb := *strings.NewReader(v.Encode())

	req, _ := http.NewRequest("POST", endpoint, &rb)
	req.SetBasicAuth(twilio.SID, twilio.Token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	resp, _ := client.Do(req)
	if resp.StatusCode != 201 {
		bytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(bytes))
		return
	}

	fmt.Println(resp.Status)
}