package main

// _TreqIneedToLogin
type _TconnInfo struct {
	ciId128 []byte
}

type _TloginCheck struct {
	ucDecodeI         chan _Tdecode
	ucCHSendLO        *chan _TudpNodeDataSend // ugCHSendI
	ucCB900101init    func(*_TloginCheck)     // if nil , use the default init procedure
	ucCB900201stCheck func(*_TloginCheck)     // if nil , use the default receive
	ucMapConnA        map[[16]byte]_TconnInfo // _TreqIneedToLogin MeIdx128
	ucMapConnAcnt     int
}
