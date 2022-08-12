package hilink

import (
	"github.com/jeadie/hilink"
	"net/url"
)

type Hilink struct {
	DeviceUrl url.URL
	c         hilink.Client
}

type Sms struct {
	From    string `json:"from"`
	Content string `json:"content"`
}

func CreateHilink(base_url string) (Hilink, error) {
	// TODO: update hilink package to use bas_url
	base, err := url.Parse(base_url)
	if err != nil {
		return Hilink{}, err
	}
	opts := []hilink.ClientOption{
		hilink.WithURL(base_url),
	}
	return Hilink{
		DeviceUrl: *base,
		c:         *hilink.NewClient(opts...),
	}, nil
}
