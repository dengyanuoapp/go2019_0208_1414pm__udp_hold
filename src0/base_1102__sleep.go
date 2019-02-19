package main

import (
    "time"
)

func _Fsleep_1ms()              { time.Sleep(                                           time.Millisecond)   } //
func _Fsleep_10ms()             { time.Sleep(   10 *                                    time.Millisecond)   } //
func _Fsleep_100ms()            { time.Sleep(   100 *                                   time.Millisecond)   } //
func _Fsleep_1s()               { time.Sleep(                                           time.Second)        } //
func _Fsleep_10s()              { time.Sleep(   10 *                                    time.Second)        } //
func _Fsleep_100s()             { time.Sleep(   100 *                                   time.Second)        } //
func _FtimeNow() string         { return time.Now().String()                } //

func _Fsleep_1sX( ___Vs int)    { for i:=0; i < ___Vs ; i++ { time.Sleep(               time.Second)        }} //
func _Fsleep_10sX( ___Vs int)   { for i:=0; i < ___Vs ; i++ { time.Sleep( 10 *          time.Second)        }} //
