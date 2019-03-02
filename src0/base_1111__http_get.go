package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func _Fhttp_get(___Vuri string) ([]byte, error) {

	// https://golang.org/pkg/net/http/#Get
	// "net/http" : func Get(url string) (resp *Response, err error)
	__Vresp, __Verr := http.Get(___Vuri)
	if __Verr != nil {
		return nil, __Verr
	}
	defer __Vresp.Body.Close() // make sure close.

	// Check server response
	if __Vresp.StatusCode != http.StatusOK {
		_FpfN(" 381911 bad status: %s", __Vresp.Status)
		return nil, fmt.Errorf("381912 bad status: %s", __Vresp.Status)
	}

	// https://golang.org/pkg/net/http/#Response
	// import "io/ioutil" : func ReadAll(r io.Reader) ([]byte, error)
	__Vout, __Verr := ioutil.ReadAll(__Vresp.Body)
	if nil != __Verr {
		_FpfN(" 381914 %v", __Verr)
		return nil, __Verr
	}
	return __Vout, nil
} // _Fhttp_get
