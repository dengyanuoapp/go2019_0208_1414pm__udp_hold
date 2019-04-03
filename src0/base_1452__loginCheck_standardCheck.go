package main

func (___Vlc *_TloginCheck) _FloginCheck__900201x__standardCheck() {
	_Fsleep(_T1s)
	if nil == ___Vlc.ucCHSendLO {
		_Pn()
		_Pn()
		_FpfNdb(" 838392 01 why no Chan output ? exit loop")
		_Pn()
		_Pn()
		return
	}

	_FpfN(" 838392 02 start")

	var __Vdecode _Tdecode
	for {
		//_Fsleep_100s()
		select {
		case __Vdecode = <-___Vlc.ucDecodeI: // _Tdecode
			//_FpfNdb(" 838392 05 : %s", __Vdecode.String()) // 15540463611554046361
			if true == __Vdecode.ok {
				switch __Vdecode.Type {
				case Cmd__loginS1ReqTryNoToken:
					//_FpfNdb(" 838392 06 : %x", __Vdecode.Dlogin.TokenL)
					//_FpfNdb(" 838392 07 : %s", __Vdecode.String()) // 15540463611554046361
					___Vlc._FloginCheck_step900201y__s2Reply_tokenB(&__Vdecode)
				default:
					_FpfNdb(" 838392 08 : unknow how to deal with : %d", __Vdecode.Type)
				}
			} else {
				_FpfNdb(" 838392 09 : why not ok ?")
			}
		}
	}
}
