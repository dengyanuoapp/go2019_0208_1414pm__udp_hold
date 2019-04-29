package main

import (
	"sync"
	"time"
)

type _TgapFilter_ReceCntX struct{}
type _TgapFilter_ReceCnt struct {
	gfrCnt int
	gfrUnr _TudpNodeDataRece
	gfrUnb []byte
}

type _TgapFilter_ReceStX struct{}
type _TgapFilter_ReceSt struct { // _TudpNodeDataReceX
	now_ map[string]_TgapFilter_ReceCnt
	last map[string]_TgapFilter_ReceCnt
	las2 map[string]_TgapFilter_ReceCnt
}

type _TgapFilterX struct{}
type _TgapFilter struct {
	gfMux                 sync.Mutex
	gfGap                 time.Duration           // default _T10s
	gfCHudpNodeDataReceLO *chan _TudpNodeDataRece // _TudpNodeDataReceX
	gfCHudpNodeDataReceI  chan _TudpNodeDataRece
	gfCHbyteLO            *chan []byte
	gfCHbyteI             chan []byte
	gfCHdelay             chan byte
	gfCBinit              func(*_TgapFilter) // _FgapFilter__1200101x__init_default
	gfCBdelay             func(*_TgapFilter)
	gfCBudpNodeDataRece   func(*_TgapFilter)
	gfCBbyte              func(*_TgapFilter)
	gfR                   _TgapFilter_ReceSt
}
