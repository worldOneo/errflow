package errflow

import (
	"io"
	"net/http"
	"net/url"
)

type HttpClient struct {
	httpClient *http.Client
	errs       errChain
}

func (httpClient *HttpClient) CloseIdleConnections() *HttpClient {
	httpClient.httpClient.CloseIdleConnections()
	return httpClient
}

func (httpClient *HttpClient) DoFlow(req *http.Request) Splitted[*http.Response, *HttpClient] {
	return SplitOf(httpClient.Do(req), httpClient)
}

func (httpClient *HttpClient) Do(req *http.Request) *http.Response {
	return Do(func() (*http.Response, error) {
		result, err := httpClient.httpClient.Do(req)
		return result, err
	}, httpClient, empty[*http.Response])
}

func (httpClient *HttpClient) GetFlow(url string) Splitted[*http.Response, *HttpClient] {
	return SplitOf(httpClient.Get(url), httpClient)
}

func (httpClient *HttpClient) Get(url string) *http.Response {
	return Do(func() (*http.Response, error) {
		result, err := httpClient.httpClient.Get(url)
		return result, err
	}, httpClient, empty[*http.Response])
}

func (httpClient *HttpClient) HeadFlow(url string) Splitted[*http.Response, *HttpClient] {
	return SplitOf(httpClient.Head(url), httpClient)
}

func (httpClient *HttpClient) Head(url string) *http.Response {
	return Do(func() (*http.Response, error) {
		result, err := httpClient.httpClient.Head(url)
		return result, err
	}, httpClient, empty[*http.Response])
}

func (httpClient *HttpClient) PostFlow(url string, contentType string, body io.Reader) Splitted[*http.Response, *HttpClient] {
	return SplitOf(httpClient.Post(url, contentType, body), httpClient)
}

func (httpClient *HttpClient) Post(url string, contentType string, body io.Reader) *http.Response {
	return Do(func() (*http.Response, error) {
		result, err := httpClient.httpClient.Post(url, contentType, body)
		return result, err
	}, httpClient, empty[*http.Response])
}

func (httpClient *HttpClient) PostFormFlow(url string, data url.Values) Splitted[*http.Response, *HttpClient] {
	return SplitOf(httpClient.PostForm(url, data), httpClient)
}

func (httpClient *HttpClient) PostForm(url string, data url.Values) *http.Response {
	return Do(func() (*http.Response, error) {
		result, err := httpClient.httpClient.PostForm(url, data)
		return result, err
	}, httpClient, empty[*http.Response])
}

func (httpClient *HttpClient) Err() error {
	return httpClient.errs.Err()
}

func (httpClient *HttpClient) Fail(err error) {
	httpClient.errs.Fail(err)
}

func (httpClient *HttpClient) Link() *error {
	return httpClient.errs.Link()
}

func (httpClient *HttpClient) LinkTo(err *error) {
	httpClient.errs.LinkTo(err)
}

func (httpClient *HttpClient) Unwrap() (*http.Client, error) {
	return httpClient.httpClient, httpClient.Err()
}

func (httpClient *HttpClient) Raw() *http.Client {
	return httpClient.httpClient
}

func HttpClientOf(httpClient *http.Client, flow Linkable) *HttpClient {
	return &HttpClient{httpClient: httpClient, errs: errChainOf(flow)}
}

func NewHttpClient(httpClient *http.Client) *HttpClient {
	return &HttpClient{httpClient: httpClient, errs: emptyChain()}
}

func EmptyHttpClientOf(err error) *HttpClient {
	return &HttpClient{errs: errChainOfErr(err)}
}
