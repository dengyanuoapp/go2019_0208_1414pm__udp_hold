package main

type _Tcount struct {
	try    uint64
	ok     uint64
	eofErr uint64
	enErr  uint64
	deErr  uint64
	otErr  uint64
}

func (___Vtc0 *_Tcount) String() string {
	return _Spf(
		"try:%d ok:%d eofE:%d enE:%d deE:%d otE:%d",
		___Vtc0.try,
		___Vtc0.ok,
		___Vtc0.eofErr,
		___Vtc0.enErr,
		___Vtc0.deErr,
		___Vtc0.otErr,
	)
}
