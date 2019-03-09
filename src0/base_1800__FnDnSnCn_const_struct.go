package main

/*
// https://www.callicoder.com/golang-basic-types-operators-type-conversion/
// Type    Size        Range
// int8    8 bits      -128 to 127
// int16   16 bits     -2^^15 to 2^^15 -1
// int32   32 bits     -2^^31 to 2^^31 -1
// int64   64 bits     -2^^63 to 2^^63 -1
// int     Platform dependent     Platform dependent
//
// Unsigned Integers
// Type    Size        Range
// uint8   8 bits      0 to 255
// uint16  16 bits     0 to 2^^16 -1
// uint32  32 bits     0 to 2^^32 -1
// uint64  64 bits     0 to 2^^64 -1
// uint    Platform dependent     Platform dependent

//                --====----====----====----====----
//	_FnPasswd = "718291a81b893a8a8a8921a8dad892a1"
//	_DnPasswd = "b87288938a81838e828291823719aba2"
//	_SnPasswd = "ca79319381829379abe9abe889ac8aa3"
//	_CnPasswd = "3019303f301ab921b8a8c8a9018a78a4"
*/

type _TsrvInfo struct {
	Name    string   // name
	Guri    string   // refresh-uri-Github
	UriArrs []string // try-Uris
	K256    []byte   // passwd to connect the this server
} // _TsrvInfo

type _TUreqNewSession struct {
	Enabled bool

	skipCnt   int
	remainCnt int

	updateUri    string
	updatePasswd *[]byte

	srvLen  int
	srvIdx  int
	srvInfo *_TsrvInfo

	serviceUdP        *_TserviceUDP
	UnewSessionCall01 func(*_TUreqNewSession) // 1: _Fhandle_u03x__udpListen__timer__main_top__default : deal with timer ARRAY in udp
	UnewSessionCall04 func(*_TUreqNewSession) // 4: _FuserCallback_u04__
	UnewSessionCall08 func(*_TUreqNewSession) // 8:

	sendBuf08 *[]byte // send tmp buf gen by UnewSessionCall08
} //    _TUreqNewSession