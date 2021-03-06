package request

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// New is the function to build a new RequestBuilder
func New() RequestBuilder {
	client := &http.Client{}
	r := RequestInput{
		Client: client,
	}

	return RequestBuilder{
		RequestInput: r,
	}
}

// NewWithClient is to simplify proxy request. Insteading of creating a new http client, we can use the request's http.Client to save time to mimick request.
func NewWithClient(c *http.Client) RequestBuilder {
	var client *http.Client
	if c == nil {
		client = &http.Client{}
	} else {
		client = c
	}
	r := RequestInput{
		Client: client,
	}
	return RequestBuilder{
		RequestInput: r,
	}
}

// SetURL sets the url
func (r *RequestBuilder) SetURL(b string) *RequestBuilder {
	r.RequestInput.URL = &b
	return r
}

// AddCert adds client cert
func (r *RequestBuilder) AddCert(b []byte) *RequestBuilder {
	r.certs = append(r.certs, b)
	return r
}

// AddKey adds the client key
func (r *RequestBuilder) AddKey(b []byte) *RequestBuilder {
	r.keys = append(r.keys, b)
	return r
}

// AddTrust adds CA
func (r *RequestBuilder) AddTrust(b []byte) *RequestBuilder {
	r.trusts = append(r.trusts, b)
	return r
}

// SetCookie sets the cookie
func (r *RequestBuilder) SetCookie(b http.CookieJar) *RequestBuilder {
	r.RequestInput.Client.Jar = b
	return r
}

// SetBodyBytes set the body in terms of bytes
func (r *RequestBuilder) SetBodyBytes(b []byte) *RequestBuilder {
	r.RequestInput.BodyBytes = b
	return r
}

// SetHeaders sets the header with input variable map[string]string
func (r *RequestBuilder) SetHeaders(h map[string]string) *RequestBuilder {
	r.RequestInput.Headers = h
	return r
}

// SetRawHeaders sets the header but passing down the http.Header as a variable
func (r *RequestBuilder) SetRawHeaders(h http.Header) *RequestBuilder {
	r.RequestInput.RawHeaders = h
	return r
}

// SetTimeOut sets the timeout
func (r *RequestBuilder) SetTimeOut(h int) *RequestBuilder {
	r.RequestInput.TimeOut = h
	return r
}

// SetMethod sets the http method
func (r *RequestBuilder) SetMethod(m string) *RequestBuilder {
	r.RequestInput.Method = &m
	return r
}

// NoVerify sets whether we should do https verifying when doing the request
func (r *RequestBuilder) NoVerify() *RequestBuilder {
	r.RequestInput.NoVerify = true
	return r
}

// SetJson sets the json as body
func (r *RequestBuilder) SetJson(j interface{}) *RequestBuilder {
	r.RequestInput.Json = &j
	return r
}

// Do is executing the request
func (r *RequestBuilder) Do() (*http.Response, error) {
	var client *http.Client
	if r.RequestInput.Client == nil {
		client = &http.Client{}
	} else {
		client = r.RequestInput.Client
	}
	var (
		h     *http.Response
		ebody *bytes.Reader
	)
	tlsConfig, err := r.getTlsConfig()
	if err != nil {
		return h, err
	}
	err = r._check()
	if err != nil {
		return h, err
	}

	// check if json exist
	if r.Json != nil {
		eJson, err := json.Marshal(*r.RequestInput.Json)
		if err != nil {
			return h, err
		}
		ebody = bytes.NewReader(eJson)
	} else {
		ebody = bytes.NewReader([]byte(""))
	}
	// check if bodyBytes exist bodybytes overwrite everything
	if r.BodyBytes != nil {
		ebody = bytes.NewReader(r.BodyBytes)
	}
	req, err := http.NewRequest(*r.RequestInput.Method, *r.RequestInput.URL, ebody)
	if err != nil {
		return h, err
	}

	if r.RequestInput.Headers != nil {
		for k, v := range r.RequestInput.Headers {
			req.Header.Set(k, v)
		}
	}
	if r.RequestInput.RawHeaders != nil {
		req.Header = r.RequestInput.RawHeaders
	}

	client.Transport = &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	// Set Timeout
	client.Timeout = time.Duration(r.RequestInput.TimeOut) * time.Second
	h, err = client.Do(req)
	if err != nil {
		return h, err
	}
	return h, nil
}

func (r *RequestBuilder) _check() error {
	// make GET as default
	if r.RequestInput.Method == nil {
		method := "GET"
		r.RequestInput.Method = &method
	}

	// check if url is valid
	_, err := url.ParseRequestURI(*r.RequestInput.URL)
	if err != nil {
		return err
	}
	if r.RequestInput.URL == nil {
		return fmt.Errorf("url not valid")
	}
	return nil
}
func (r *RequestBuilder) getTlsConfig() (*tls.Config, error) {
	// Mtls request
	if len(r.certs) != len(r.keys) {
		return &tls.Config{}, fmt.Errorf("number of certs and keys are not the same")
	}
	var certList []tls.Certificate
	if len(r.certs) != 0 && len(r.keys) != 0 && len(r.trusts) != 0 {
		for i := 0; i < len(r.certs); i++ {
			cert, err := tls.X509KeyPair(r.certs[i], r.keys[i])
			if err != nil {
				return &tls.Config{}, err
			}
			certList = append(certList, cert)
		}
	}
	toReturn := &tls.Config{
		InsecureSkipVerify: r.RequestInput.NoVerify,
		Certificates:       certList,
	}
	if len(r.trusts) != 0 {
		clientCertPool := x509.NewCertPool()
		for _, trust := range r.trusts {
			clientCertPool.AppendCertsFromPEM(trust)
		}
		toReturn.RootCAs = clientCertPool
	}
	return toReturn, nil

}
