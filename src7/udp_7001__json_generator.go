package main

type _TjFn struct {
	N string
	G string
	U []string
} // _TjFn

var (
	_VjFn _TjFn = _TjFn{"Fn", "https://github.com/jasas78/jsonOnly/raw/master/Makefile",
		[]string{}}
)

// _FencAesRand_only
// _FencAesRandExit
func _Ftry_gen_json01(___Vfname string, ___Vkey *[]byte, ___Vdst interface{}) {

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

	_FaesRand_test__en_de_Exit("asa90as90sa90as90as9as82391")
	_Fex1(" 389181 debug stop ")

	_Fwrite_json_and_rand_Exit(" 389189 ", ___Vkey, "json/"+___Vfname+".json", ___Vdst)
} // _Ftry_gen_json01

func main() {
	_FpfN("\n Start \n")

	_Ftry_gen_json01("fn", &_VpwFile_Dn, &_VjFn)

} // main