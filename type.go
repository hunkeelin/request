package request

import (
	"crypto/tls"
	"net/http"
)

// RequestInput is request Input.
type RequestInput struct {
	URL        *string           // URL
	Headers    map[string]string //Headers
	RawHeaders http.Header       //RawHeaders
	Method     *string           // Method
	TimeOut    int               // TimeOut
	NoVerify   bool              // NoVerify
	BodyBytes  []byte            // BodyBytes
	Json       *interface{}      // Json
	Client     *http.Client      // Client
	TlsConfig  tls.Config        // TlsConfig
	certs      [][]byte
	keys       [][]byte
	trusts     [][]byte
}

// RequestBuilder is the request builder struct for method chaining
type RequestBuilder struct {
	RequestInput
}
