package safehttp

import (
	"net"
	"net/http"
	"time"
)

func NewClient(client *http.Client, filter Filter) *http.Client {
	if client == nil {
		client = &http.Client{}
	}

	var transport *http.Transport

	if client.Transport == nil {
		transport = &(*(http.DefaultTransport.(*http.Transport)))
		client.Transport = transport
	} else {
		transport = client.Transport.(*http.Transport)
	}

	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}

	transport.DialContext = (&FilterDialer{dialer, filter}).DialContext

	return client
}
