package errflow

import (
	"io"
	"net/http"
	"net/url"
)

// HTTPClient is the flow wrapper of *http.Client
type HTTPClient struct {
	hTTPClient *http.Client
	errs       errChain
}

// CloseIdleConnections does the same as *http.Client.CloseIdleConnections but is a noop if this flow
// failed.
// Returns this flow.
func (hTTPClient *HTTPClient) CloseIdleConnections() *HTTPClient {
	if hTTPClient.errs.Err() != nil {
		return hTTPClient
	}
	hTTPClient.hTTPClient.CloseIdleConnections()
	return hTTPClient
}

// DoFlow does the same as *http.Client.Do but splits the flow.
func (hTTPClient *HTTPClient) DoFlow(req *http.Request) Splitted[*http.Response, *HTTPClient] {
	return SplitOf(hTTPClient.Do(req), hTTPClient)
}

// Do does the same as *http.Client but is a noop if this flow already failed.
func (hTTPClient *HTTPClient) Do(req *http.Request) *http.Response {
	return Do(func() (*http.Response, error) {
		result, err := hTTPClient.hTTPClient.Do(req)
		return result, err
	}, hTTPClient, empty[*http.Response])
}

// GetFlow does the same as *http.Client.Get but splits the flow.
func (hTTPClient *HTTPClient) GetFlow(url string) Splitted[*http.Response, *HTTPClient] {
	return SplitOf(hTTPClient.Get(url), hTTPClient)
}

// Get does the same as *http.Client but is a noop if this flow already failed.
func (hTTPClient *HTTPClient) Get(url string) *http.Response {
	return Do(func() (*http.Response, error) {
		result, err := hTTPClient.hTTPClient.Get(url)
		return result, err
	}, hTTPClient, empty[*http.Response])
}

// HeadFlow does the same as *http.Client.Head but splits the flow.
func (hTTPClient *HTTPClient) HeadFlow(url string) Splitted[*http.Response, *HTTPClient] {
	return SplitOf(hTTPClient.Head(url), hTTPClient)
}

// Head does the same as *http.Client but is a noop if this flow already failed.
func (hTTPClient *HTTPClient) Head(url string) *http.Response {
	return Do(func() (*http.Response, error) {
		result, err := hTTPClient.hTTPClient.Head(url)
		return result, err
	}, hTTPClient, empty[*http.Response])
}

// PostFlow does the same as *http.Client.Post but splits the flow.
func (hTTPClient *HTTPClient) PostFlow(url string, contentType string, body io.Reader) Splitted[*http.Response, *HTTPClient] {
	return SplitOf(hTTPClient.Post(url, contentType, body), hTTPClient)
}

// Post does the same as *http.Client but is a noop if this flow already failed.
func (hTTPClient *HTTPClient) Post(url string, contentType string, body io.Reader) *http.Response {
	return Do(func() (*http.Response, error) {
		result, err := hTTPClient.hTTPClient.Post(url, contentType, body)
		return result, err
	}, hTTPClient, empty[*http.Response])
}

// PostFormFlow does the same as *http.Client.PostForm but splits the flow.
func (hTTPClient *HTTPClient) PostFormFlow(url string, data url.Values) Splitted[*http.Response, *HTTPClient] {
	return SplitOf(hTTPClient.PostForm(url, data), hTTPClient)
}

// PostForm does the same as *http.Client but is a noop if this flow already failed.
func (hTTPClient *HTTPClient) PostForm(url string, data url.Values) *http.Response {
	return Do(func() (*http.Response, error) {
		result, err := hTTPClient.hTTPClient.PostForm(url, data)
		return result, err
	}, hTTPClient, empty[*http.Response])
}

// Err returns the error of this flow if any happend.
func (hTTPClient *HTTPClient) Err() error {
	return hTTPClient.errs.Err()
}

// Fail ends this flow with err
func (hTTPClient *HTTPClient) Fail(err error) {
	hTTPClient.errs.Fail(err)
}

// Link returns the base error of this flow.
func (hTTPClient *HTTPClient) Link() *error {
	return hTTPClient.errs.Link()
}

// LinkTo merges err as base into this flow.
func (hTTPClient *HTTPClient) LinkTo(err *error) {
	hTTPClient.errs.LinkTo(err)
}

// Unwrap returns the internal value and the produced error if any.
func (hTTPClient *HTTPClient) Unwrap() (*http.Client, error) {
	return hTTPClient.hTTPClient, hTTPClient.Err()
}

// Raw returns the internal value.
// No Guarantees.
func (hTTPClient *HTTPClient) Raw() *http.Client {
	return hTTPClient.hTTPClient
}

// HTTPClientOf create a new HTTPClient and is linked to flow.
// They are linked to gether and if one errors the error is propagated.
func HTTPClientOf(hTTPClient *http.Client, flow Linkable) *HTTPClient {
	return &HTTPClient{hTTPClient: hTTPClient, errs: errChainOf(flow)}
}

// NewHTTPClient create a new HTTPClient and is the root of a flow.
// It will catch any error that happens in the future.
func NewHTTPClient(hTTPClient *http.Client) *HTTPClient {
	return &HTTPClient{hTTPClient: hTTPClient, errs: emptyChain()}
}

// EmptyHTTPClientOf returns an already failed HTTPClient.
// Calls will have no effects on it.
func EmptyHTTPClientOf(err error) *HTTPClient {
	return &HTTPClient{errs: errChainOfErr(err)}
}
