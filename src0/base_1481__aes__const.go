package main

import "sync"

type _TaesX struct {
	aesKey        []byte
	aesEncodeMode bool // true --> encode mode , false --> decode mode
	aesRandMode   bool // true --> add rand pad , false --> do nothing
	aesOk         bool // encode / decode result
	aesBufIn      []byte
	aesBufOut     []byte
	aesMux        sync.Mutex
	aesErr        string
}
type _Taes struct {
}

func (___Vaes *_Taes) String() string {
	return _Spf(
		"ok:%t k:%s En:%t Rn:%t bIn(%d){%s} bOut(%d){%s} err[%s]",
		___Vaes.aesOk,
		String5(&___Vaes.aesKey),
		___Vaes.aesEncodeMode,
		___Vaes.aesRandMode,
		len(___Vaes.aesBufIn),
		String5(&___Vaes.aesBufIn),
		len(___Vaes.aesBufOut),
		String5(&___Vaes.aesBufOut),
		___Vaes.aesErr,
	)
}
