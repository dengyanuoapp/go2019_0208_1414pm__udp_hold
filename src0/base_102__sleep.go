package main

import (
    "time"
)

func _Fsleep_1ms()      { time.Sleep(           time.Millisecond)   } //
func _Fsleep_10ms()     { time.Sleep(   10 *    time.Millisecond)   } //
func _Fsleep_100ms()    { time.Sleep(   100 *   time.Millisecond)   } //
func _Fsleep_1s()       { time.Sleep(           time.Second)        } //
func _Fsleep_10s()      { time.Sleep(   10 *    time.Second)        } //
func _Fsleep_100s()     { time.Sleep(   100 *   time.Second)        } //
