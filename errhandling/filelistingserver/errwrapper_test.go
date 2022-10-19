package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

type userError string

func (u userError) Error() string {
	return u.message()
}

func (u userError) message() string {
	return string(u)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return userError("testing user error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown Error")
}

func noError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "testing user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "https://www.baidu.com", nil)
		f(response, request)

		verifyResponse(t, response.Result(), tt.code, tt.message)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f)) // http.HandlerFunc(): convert func to interface
		resp, _ := http.Get(server.URL)

		verifyResponse(t, resp, tt.code, tt.message)
	}
}

func verifyResponse(t *testing.T, resp *http.Response, expectedCode int, expectedMessage string) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")

	if resp.StatusCode != expectedCode || body != expectedMessage {
		t.Errorf("Expected (%d, %s); got (%d, %s)!", expectedCode, expectedMessage, resp.StatusCode, body)
	}
}
