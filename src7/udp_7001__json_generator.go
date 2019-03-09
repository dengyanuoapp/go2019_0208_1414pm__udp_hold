package main

type _TjFn struct {
	N string
	G string
	U []string
} // _TjFn

var (
	_VjFn _TjFn = _TjFn{"Fn", "https://github.com/jasas78/jsonOnly/raw/master/Makefile",
		[]string{}}
	_Vself   _Tself
	_Vconfig _Tconfig
)

// _FencAesRand_only
// _FencAesRandExit
func _Ftest01(___Vfname string, ___Vkey *[]byte, ___Vdst interface{}) {

	if 2 == 3 {
		// _FreGenRandBuf___
		// _FgenRand_nByte__
		_FgenRand_nByte__testExit(1) // 1s
		//_FgenRand_nByte__testExit(1000)       // 6s
		//_FgenRand_nByte__testExit(10000)      // 37s
		//_FgenRand_nByte__testExit(100000)     // 300s
		//_FgenRand_nByte__testExit(1000000)    // 3360s
		_Fex1(" 389181 debug stop ")
	}

	// _FdecAesCbc__only___
	// _FdecAesRand__only
	//var __VloopAmount = 10000000 // 2:23 == 120+23 == 143 s
	//var __VloopAmount = 1000000 // 18s
	var __VloopAmount = 100000 // 5s
	_FaesRand_test__en_de_Exit("838189183jsadas78asklas9p/avh7o8asadhasi,s,hasjsdjdsa873n1kadsuikl189034jasd", __VloopAmount)
	//_Fex1(" 389183 debug stop ")

	//_Fwrite_json_and_rand_Exit(" 389189 ", ___Vkey, "json/"+___Vfname+".json", ___Vdst)
} // _Ftest01

func main() {
	_FpfN("\n Start \n")

	_Ftest01("fn", &_Vpasswd_udp_Fn_waitForCliens01, &_VjFn)

} // main
