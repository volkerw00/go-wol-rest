package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_shouldFailOnMissingMAC(t *testing.T) {
	test = t

	server := NewServer()
	response := httptest.NewRecorder()
	request := newRequest("POST", "/wake")

	server.ServeHTTP(response, request)

	responseBody := bodyAsString(response)
	assertThatInt(response.Code, isInt(http.StatusBadRequest))
	assertThatString(responseBody, containsString("mac is missing"))
}

func Test_shouldFailOnUnparsableMAC(t *testing.T) {
	test = t

	server := NewServer()
	response := httptest.NewRecorder()
	request := newRequest("POST", "/wake?mac=XXX")

	server.ServeHTTP(response, request)

	responseBody := bodyAsString(response)
	assertThatInt(response.Code, isInt(http.StatusBadRequest))
	assertThatString(responseBody, containsString("failed to parse XXX as a MAC adress"))
}

func Test_shouldFailOnUnparsableIP(t *testing.T) {
	test = t

	server := NewServer()
	response := httptest.NewRecorder()
	request := newRequest("POST", "/wake?mac=12:34:56:78:9A:BC&broadcastIP=XXX")

	server.ServeHTTP(response, request)

	responseBody := bodyAsString(response)
	assertThatInt(response.Code, isInt(http.StatusBadRequest))
	assertThatString(responseBody, containsString("failed to parse XXX as a IPv4 adress"))
}

func Test_shouldSendMagicPacket(t *testing.T) {
	test = t

	server := NewServer()
	response := httptest.NewRecorder()
	request := newRequest("POST", "/wake?mac=12:34:56:78:9A:BC&broadcastIP=127.0.0.255")

	server.ServeHTTP(response, request)

	assertThatInt(response.Code, isInt(http.StatusNoContent))
}

func newRequest(method, url string) *http.Request {
	request, error := http.NewRequest(method, url, nil)
	if error != nil {
		test.Fatal(error)
	}
	return request
}

func bodyAsString(response *httptest.ResponseRecorder) string {
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		test.Fatal(error)
	}
	return string(body)
}
