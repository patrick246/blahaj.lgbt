package prometheus

import (
	"net/http"
)

type BasicAuthRoundTripper struct {
	Username     string
	Password     string
	RoundTripper http.RoundTripper
}

func (b *BasicAuthRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	if b.Username != "" && b.Password != "" {
		request.SetBasicAuth(b.Username, b.Password)
	}

	if b.RoundTripper != nil {
		return b.RoundTripper.RoundTrip(request)
	}

	return http.DefaultTransport.RoundTrip(request)
}
