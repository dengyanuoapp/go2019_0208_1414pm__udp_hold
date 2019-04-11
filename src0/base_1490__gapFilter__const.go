package main

import (
	"sync"
	"time"
)

type _TgapFilterX struct {
}
type _TgapFilter struct {
	gfMux                 sync.Mutex
	gfGap                 time.Duration // default _T10s
	gfCHudpNodeDataReceLO *chan _TudpNodeDataRece
	gfCHbyteLO            *chan []byte
	gfCHudpNodeDataReceI  chan _TudpNodeDataRece
	gfCHbyteI             chan []byte
}
