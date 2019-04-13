package main

import (
	"sync"
	"time"
)

type _TgapFilter_ReceST struct { // _TudpNodeDataReceX
	now_ map[string]_TudpNodeDataRece
	last map[string]_TudpNodeDataRece
	las2 map[string]_TudpNodeDataRece
}

type _TgapFilterX struct {
}
type _TgapFilter struct {
	gfMux                 sync.Mutex
	gfGap                 time.Duration           // default _T10s
	gfCHudpNodeDataReceLO *chan _TudpNodeDataRece // _TudpNodeDataReceX
	gfCHbyteLO            *chan []byte
	gfCHudpNodeDataReceI  chan _TudpNodeDataRece
	gfCHbyteI             chan []byte
	gfCHdelay             chan byte
	gfCBinit              func(*_TgapFilter) // _FgapFilter__1200101x__init_default
	gfCBdelay             func(*_TgapFilter)
	gfCBudpNodeDataRece   func(*_TgapFilter)
	gfCBbyte              func(*_TgapFilter)
	gfR                   _TgapFilter_ReceST
}
