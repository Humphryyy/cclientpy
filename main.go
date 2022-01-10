package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"unsafe"

	http "github.com/Carcraftz/fhttp"

	"bytes"
	"compress/gzip"
	"compress/zlib"
	"io"
	"time"

	cclient "github.com/Carcraftz/cclient"
	tls "github.com/Carcraftz/utls"

	"github.com/andybalholm/brotli"
)

func main() {
	for {
		fmt.Println(sendRequest(Request{
			URL:           "http://www.httpbin.org/get",
			Method:        "GET",
			Headers:       [][]string{},
			Body:          "",
			AllowRedirect: true,
			Proxy:         "",
			Timeout:       10000,
			PseudoHeaderOrder: []string{
				":method",
				":authority",
				":scheme",
				":path",
			},
		}))
	}
}

//export SendRequest
func SendRequest(requestC *C.char) *C.char {
	requestString := C.GoString(requestC)

	request := Request{}
	if err := json.Unmarshal([]byte(requestString), &request); err != nil {
		panic(err)
	}

	response := sendRequest(request)

	respBytes, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	return C.CString(string(respBytes))
}

//export FreePTR
func FreePTR(ptr *C.char) {
	C.free(unsafe.Pointer(ptr))
}

func sendRequest(request Request) Response {
	client, err := cclient.NewClient(tls.HelloChrome_Auto, request.Proxy, request.AllowRedirect, time.Duration(request.Timeout)*time.Millisecond)
	if err != nil {
		panic(err)
	}

	var body io.Reader
	if request.Body != "" {
		body = bytes.NewBufferString(request.Body)
	}

	req, err := http.NewRequest(request.Method, request.URL, body)
	if err != nil {
		panic(err)
	}

	headers := http.Header{
		http.PHeaderOrderKey: request.PseudoHeaderOrder,
	}

	var headerOrder []string
	for _, headerPair := range request.Headers {
		headerOrder = append(headerOrder, headerPair[0])
		headers[headerPair[0]] = []string{headerPair[1]}
	}

	headers[http.HeaderOrderKey] = headerOrder

	req.Header = headers

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var respBody []byte

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			panic(err)
		}
		defer reader.Close()

		respBody, err = ioutil.ReadAll(reader)
		if err != nil {
			panic(err)
		}
	case "deflate":
		reader, err := zlib.NewReader(resp.Body)
		if err != nil {
			panic(err)
		}
		defer reader.Close()

		respBody, err = ioutil.ReadAll(reader)
		if err != nil {
			panic(err)
		}
	case "br":
		reader := brotli.NewReader(resp.Body)

		respBody, err = ioutil.ReadAll(reader)
		if err != nil {
			panic(err)
		}
	default:
		respBody, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
	}

	var respHeaders [][]string
	for key, value := range resp.Header {
		respHeaders = append(respHeaders, []string{key, value[0]})
	}

	return Response{
		StatusCode: resp.StatusCode,
		Headers:    respHeaders,
		Body:       string(respBody),
	}
}
