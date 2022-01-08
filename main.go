package main

import (
	"C"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"unsafe"

	http "github.com/Carcraftz/fhttp"

	cclient "github.com/Carcraftz/cclient"
	tls "github.com/Carcraftz/utls"
)
import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"io"
	"time"

	"github.com/andybalholm/brotli"
)

func main() {}

//export SendRequest
func SendRequest(requestC *C.char) unsafe.Pointer {
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

	length := make([]byte, 8)

	binary.LittleEndian.PutUint64(length, uint64(len(respBytes)))
	return C.CBytes(append(length, respBytes...))
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
		Headers: respHeaders,
		Body:    string(respBody),
	}
}
