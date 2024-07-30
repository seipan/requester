package requester

import "net/url"

type Requester struct {
	// ...

	URL url.URL
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
