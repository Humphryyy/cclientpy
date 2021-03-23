package main

import (
	"C"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"unsafe"

	/* I take no credit for cclient I just had a hard time go getting it and had to change the .mod, sorry x04 */
	cclient "github.com/IHaveNothingg/cclientwtf"
	tls "github.com/Titanium-ctrl/utls"
)

var sb string

type Headers struct {
	Headers []Header `json:"Headers"`
}
type Header struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

func Split(str, sep string) []string

//export cclientpy
func cclientpy(url *C.char, headersfrompy *C.char) unsafe.Pointer {
	headers := C.GoString(headersfrompy)

	s := C.GoString(url)

	request := UrlGET(s, headers)

	length := make([]byte, 8)

	binary.LittleEndian.PutUint64(length, uint64(len(request)))
	return C.CBytes(append(length, request...))
}

func UrlGET(urlsent string, headerslices string) string {

	client, err := cclient.NewClient(tls.HelloChrome_Auto)

	if err != nil {
		println(err)
		return ""

	}

	var url = urlsent
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		println(err)
		return ""

	}

	var headers Headers

	fer, err := []byte(headerslices), err
	if err != nil {
		println(err)
		return ""

	}

	json.Unmarshal(fer, &headers)

	for i := 0; i < len(headers.Headers); i++ {
		req.Header.Add(headers.Headers[i].Name, headers.Headers[i].Value)
	}

	resp, err := client.Do(req)

	if err != nil {
		println(err)
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err)
		return ""
	}

	sb := string(body)

	return sb
}

func main() {
}
