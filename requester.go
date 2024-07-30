package requester

import "net/url"

type Requester struct {
	// ...

	URL url.URL
}

func NewRequester() *Requester {
	return &Requester{}
}
