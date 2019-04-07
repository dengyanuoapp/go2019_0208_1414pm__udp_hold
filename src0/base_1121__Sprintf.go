package main

import "fmt"

var (
	_Spf  func(___Vfmt string, ___Vpara ...interface{}) string = fmt.Sprintf
	_Spt  func(___Vpara ...interface{}) string                 = fmt.Sprint
	_Spln func(___Vpara ...interface{}) string                 = fmt.Sprintln
)

func _Sph() string {
	if "" != _Vself.ProjName {
		return _Spf("%s:", _Vself.ProjName)
	}
	return ""
} // _Sph
