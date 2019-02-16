package main

import (
    "time"
)

func _Fsleep_1ms() {
    time.Sleep(time.Millisecond)
} // _Fsleep_1ms
func _Fsleep_10ms() {
    time.Sleep(10 * time.Millisecond)
} // _Fsleep_10ms
func _Fsleep_100ms() {
    time.Sleep(time.Millisecond)
} // _Fsleep_100ms
func _Fsleep_1s() {
    time.Sleep(time.Second)
} // _Fsleep_1s
func _Fsleep_10s() {
    time.Sleep(10 * time.Second)
} // _Fsleep_10s
func _Fsleep_100s() {
    time.Sleep(100 * time.Second)
} // _Fsleep_100s
