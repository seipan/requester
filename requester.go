package requester

import (
	"errors"
	"net/url"
)

type Requester struct {
	// ...

	URL url.URL

	Type string
}

func (r *Requester) ValidateType() error {
	if r.Type == "JSON" || r.Type == "XML" || r.Type == "Multipart" || r.Type == "Form" {
		return nil
	}
	return errors.New("invalid type")
}

func NewRequester(urlstr string) *Requester {
	url, err := url.Parse(urlstr)
	if err != nil {
		panic(err)
	}
	return &Requester{
		URL: *url,
	}
}
