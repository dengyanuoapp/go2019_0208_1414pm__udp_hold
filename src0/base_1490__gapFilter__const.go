package main

import (
	"sync"
	"time"
)

type _TgapFilter_ReceCnt struct {
	cnt int
	unr _TudpNodeDataRece
}

type _TgapFilter_ReceStX struct {
}
type _TgapFilter_ReceSt struct { // _TudpNodeDataReceX
	now_ map[string]_TgapFilter_ReceCnt
	last map[string]_TgapFilter_ReceCnt
	las2 map[string]_TgapFilter_ReceCnt
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
	gfR                   _TgapFilter_ReceSt
}
