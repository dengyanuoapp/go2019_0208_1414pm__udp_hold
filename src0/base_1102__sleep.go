package main

import (
	"time"
)

const (
	_T1s    = 1 * time.Second
	_T5s    = 5 * time.Second
	_T10s   = 10 * time.Second
	_T15s   = 15 * time.Second
	_T20s   = 20 * time.Second
	_T25s   = 25 * time.Second
	_T50s   = 50 * time.Second
	_T60s   = 60 * time.Second
	_T80s   = 80 * time.Second
	_T100s  = 100 * time.Second
	_T110s  = 110 * time.Second
	_T120s  = 120 * time.Second
	_T140s  = 140 * time.Second
	_T360s  = 360 * time.Second
	_T600s  = 600 * time.Second
	_T1200s = 1200 * time.Second
	_T2400s = 2400 * time.Second
	_T3600s = 3600 * time.Second
	_T7200s = 7200 * time.Second
	_T1ms   = 1 * time.Millisecond
	_T5ms   = 5 * time.Millisecond
	_T10ms  = 10 * time.Millisecond
	_T15ms  = 15 * time.Millisecond
	_T20ms  = 20 * time.Millisecond
	_T25ms  = 25 * time.Millisecond
	_T50ms  = 50 * time.Millisecond
	_T80ms  = 80 * time.Millisecond
	_T100ms = 100 * time.Millisecond
	_T110ms = 110 * time.Millisecond
	_T120ms = 120 * time.Millisecond
	_T140ms = 140 * time.Millisecond
	_T240ms = 240 * time.Millisecond
	_T360ms = 360 * time.Millisecond
	_T720ms = 720 * time.Millisecond
)

// func time.Sleep(d Duration)

func _Fsleep_1ms()               { time.Sleep(time.Millisecond) }       //
func _Fsleep_10ms()              { time.Sleep(10 * time.Millisecond) }  //
func _Fsleep_100ms()             { time.Sleep(100 * time.Millisecond) } //
func _Fsleep_1s()                { time.Sleep(time.Second) }            //
func _Fsleep_10s()               { time.Sleep(10 * time.Second) }       //
func _Fsleep_30s()               { time.Sleep(30 * time.Second) }       //
func _Fsleep_50s()               { time.Sleep(50 * time.Second) }       //
func _Fsleep_100s()              { time.Sleep(100 * time.Second) }      //
func _FtimeNow() string          { return time.Now().String() }         //
func _FtimeI64() int64           { return time.Now().Unix() }           //
func _FtimeInt() int             { return int(time.Now().Unix()) }      //
func _Fsleep(___d time.Duration) { time.Sleep(___d) }                   // _T1s

func _Fsleep_1sX(___Vs int) {
	for i := 0; i < ___Vs; i++ {
		time.Sleep(time.Second)
	}
} //
func _Fsleep_10sX(___Vs int) {
	for i := 0; i < ___Vs; i++ {
		time.Sleep(10 * time.Second)
	}
} //
