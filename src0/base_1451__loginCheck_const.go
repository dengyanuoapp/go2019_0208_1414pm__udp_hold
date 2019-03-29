package main

type _TloginCheck struct {
	ucDecodeI         chan _Tdecode
	ucCHSendLO        *chan _TudpNodeDataSend // ugCHSendI
	ucCB900101init    func(*_TloginCheck)     // if nil , use the default init procedure
	ucCB900201stCheck func(*_TloginCheck)     // if nil , use the default receive
	ucMapStep101      map[[48]byte]_Tdecode
}
