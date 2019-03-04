package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func _Fhttp_getAll(___skipTLS bool, ___Vuri string) ([]byte, error) {

	// https://golang.org/pkg/net/http/#Get
	// "net/http" : func Get(url string) (resp *Response, err error)
	//__Vresp, __Verr := http.Get(___Vuri)
	var __Vresp *http.Response
	var __Verr error

	if ___skipTLS == true { // skip tls
		__Vtransport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
		__Vtimeout := time.Duration(15 * time.Second)
		__Vclient := &http.Client{
			Timeout:   __Vtimeout,
			Transport: __Vtransport}
		__Vresp, __Verr = __Vclient.Get(___Vuri)
	} else { // using tls
		__Vresp, __Verr = http.Get(___Vuri)
	}

	if __Verr != nil {
		return nil, __Verr
	}
	defer __Vresp.Body.Close() // make sure close.

	// Check server response
	if __Vresp.StatusCode != http.StatusOK {
		_FpfN(" 381911 01 bad status: %s", __Vresp.Status)
		return nil, fmt.Errorf("381912 bad status: %s", __Vresp.Status)
	}

	// https://golang.org/pkg/net/http/#Response
	// import "io/ioutil" : func ReadAll(r io.Reader) ([]byte, error)
	__Vout, __Verr := ioutil.ReadAll(__Vresp.Body)
	if nil != __Verr {
		_FpfN(" 381911 04 %v", __Verr)
		return nil, __Verr
	}
	return __Vout, nil
} // _Fhttp_getAll

func _Fhttp_getLimit(___skipTLS bool, ___Vuri string, ___VmaxLen int64) ([]byte, error) {

	// https://golang.org/pkg/net/http/#Get
	// "net/http" : func Get(url string) (resp *Response, err error)
	//__Vresp, __Verr := http.Get(___Vuri)
	var __Vresp *http.Response
	var __Verr error

	if ___skipTLS == true { // skip tls
		__Vtransport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
		__Vtimeout := time.Duration(15 * time.Second)
		__Vclient := &http.Client{
			Timeout:   __Vtimeout,
			Transport: __Vtransport}
		__Vresp, __Verr = __Vclient.Get(___Vuri)
	} else { // using tls
		__Vresp, __Verr = http.Get(___Vuri)
	}

	if __Verr != nil {
		return nil, __Verr
	}
	defer __Vresp.Body.Close() // make sure close.

	// Check server response
	if __Vresp.StatusCode != http.StatusOK {
		_FpfN(" 381914 01 bad status: %s", __Vresp.Status)
		return nil, fmt.Errorf("381912 bad status: %s", __Vresp.Status)
	}

	// https://golang.org/pkg/io/#LimitReader
	// func LimitReader(r Reader, n int64) Reader
	__Vlimit_1m := io.LimitReader(__Vresp.Body, ___VmaxLen)

	// https://golang.org/pkg/net/http/#Response
	// import "io/ioutil" : func ReadAll(r io.Reader) ([]byte, error)
	__Vout, __Verr := ioutil.ReadAll(__Vlimit_1m)
	if nil != __Verr {
		_FpfN(" 381914 04 %v", __Verr)
		return nil, __Verr
	}
	return __Vout, nil
} // _Fhttp_getLimit
