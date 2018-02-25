// Copyright (c) 2017 Philipp Weber
// Use of this source code is governed by the MIT license
// which can be found in the repositorys LICENSE file.

package hibpgo

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"time"
)

const (
	// APIUrl contains the base URL of the RESTful endpoint.
	APIUrl = "https://haveibeenpwned.com/api/"
	// APIContentNeg contains the content negotiation string for the endpoint.
	APIContentNeg = "application/vnd.haveibeenpwned.v2+json"
	// APIRateLimit contains the minimum amount of time between two requests in milliseconds. Should be used as sleep time when performing multiple requests sequentially!
	APIRateLimit = time.Duration(1500) * time.Millisecond
)

// UserAgent is used in each requests header. Could be set to individual string.
var UserAgent = "hibpgo-" + runtime.Version()

// transport includes multiple timeouts for reliable http clients.
var transport = &http.Transport{
	Dial: (&net.Dialer{
		Timeout: 5 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 5 * time.Second,
}

// client contains default client for http requests with timeouts.
var client = &http.Client{
	Timeout:   time.Second * 10,
	Transport: transport,
}

var (
	// ErrUnknownEndpoint occures if the specified endpoint does not exist.
	ErrUnknownEndpoint = errors.New("the specified endpoint does not exist")
	// ErrMalformedAccount occures if the specified account does not meet the required format.
	ErrMalformedAccount = errors.New("the account does not comply with an acceptable format (i.e. it's an empty string)")
	// ErrNoUserAgent occures if the header field user-agent does not exist or is invalid.
	ErrNoUserAgent = errors.New("no user agent has been specified in the request")
	// ErrInvalidStatus occures if the endpoint returns an http status code which is not specified in the API documentation.
	ErrInvalidStatus = errors.New("got an unspecified http status code from endpoint")
)

// getURL builds up an URL with optional querys to a specific endpoint.
func getURL(endpoint, keyword string, querys []Query) (*url.URL, error) {
	URL, err := url.Parse(APIUrl)
	if err != nil {
		return nil, err
	}
	URL.Path += endpoint + "/" + keyword

	queryString := url.Values{}
	for _, q := range querys {
		queryString.Add(q.Parameter, q.Value)
	}
	URL.RawQuery = queryString.Encode()

	return URL, nil
}

// callEndpoint calls an endpoint and returns the request body as string.
func callEndpoint(endpoint, keyword string, querys []Query) ([]byte, error) {
	var req = &http.Request{}

	// check if requested endpoint is specified and build up new request
	switch endpoint {
	case "breachedaccount":
		fallthrough
	case "breaches":
		fallthrough
	case "breach":
		fallthrough
	case "dataclasses":
		fallthrough
	case "pasteaccount":
		// get URL to endpoint
		URL, err := getURL(endpoint, keyword, querys)
		if err != nil {
			return []byte{}, err
		}
		// build up request
		req, err = http.NewRequest("GET", URL.String(), nil)
		if err != nil {
			return []byte{}, err
		}
		break
	case "pwnedpassword":
		// get URL to endpoint
		URL, err := getURL(endpoint, keyword, querys)
		if err != nil {
			return []byte{}, err
		}
		// build up request
		req, err = http.NewRequest("GET", URL.String(), nil)
		if err != nil {
			return []byte{}, err
		}
	default:
		return []byte{}, ErrUnknownEndpoint
	}

	// set user-agent and content negotiation header
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Accept", APIContentNeg)

	// do request with global client (with timeouts)
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	// suck in content from response body and close it
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	err = resp.Body.Close()
	if err != nil {
		return []byte{}, err
	}

	// switch through http status codes
	switch resp.StatusCode {
	// 200 = everything worked and there's a string array of pwned sites for the account; password found
	case http.StatusOK:
		// if content is empty with status code 200 it was a password search, return password
		if len(content) == 0 {
			content = []byte(keyword)
		}
		return content, nil
	// 400 = the account does not comply with an acceptable format (i.e. it's an empty string)
	case http.StatusBadRequest:
		return []byte{}, ErrMalformedAccount
	// 403 = no user agent has been specified in the request
	case http.StatusForbidden:
		return []byte{}, ErrNoUserAgent
	// 404 = the account could not be found and has therefore not been pwned; password not found
	case http.StatusNotFound:
		return []byte{}, nil
	// 429 = the rate limit has been exceeded
	case http.StatusTooManyRequests:
		// get Retry-After header field and calculate sleep time
		i, err := strconv.Atoi(resp.Header.Get("Retry-After"))
		if err != nil {
			return []byte{}, err
		}
		time.Sleep(time.Duration(i)*time.Second + 100*time.Millisecond)
		return callEndpoint(endpoint, keyword, querys)
	// should not happen, string contains further information from RESTful service
	default:
		return content, ErrInvalidStatus
	}
}
