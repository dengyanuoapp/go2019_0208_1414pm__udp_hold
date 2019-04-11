package main

import (
	"sync"
)

type _TgapFilterX struct {
}
type _TgapFilter struct {
	gfMux sync.Mutex
}
