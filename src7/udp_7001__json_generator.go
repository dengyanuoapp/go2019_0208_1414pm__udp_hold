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

	// _FreGenRandBuf___
	// _FgenRand_nByte__
	_FgenRand_nByte__testExit(8)
	_FaesRand_test__en_de_Exit("asa90as90sa90as90as9as82391")

	_Fwrite_json_and_rand_Exit(" 192181 ", ___Vkey, "json/"+___Vfname+".json", ___Vdst)
} // _Ftry_gen_json01

func main() {
	_FpfN("\n Start \n")

	_Ftry_gen_json01("fn", &_VpwFile_Dn, &_VjFn)

} // main
