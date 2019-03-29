package main

func (___Vlc *_TloginCheck) _FloginCheck__900201x__standardCheck() {
	_Fsleep(_T1s)
	if nil == ___Vlc.ucCHSendLO {
		_FpfNdb(" 838392 01 why no Chan output ? exit loop")
		return
	}

	_FpfN(" 838392 02 start")

	var __Vdecode _Tdecode
	for {
		//_Fsleep_100s()
		select {
		case __Vdecode = <-___Vlc.ucDecodeI: // _Tdecode
			_FpfNdb(" 838392 05 : %v", __Vdecode)
			if true == __Vdecode.ok {
				switch __Vdecode.Type {
				case Cmd__loginS1ReqTryNoToken:
				default:
					_FpfNdb(" 838392 07 : unknow how to deal with : %d", __Vdecode.Type)
				}
			} else {
				_FpfNdb(" 838392 08 : why not ok ?")
			}
		}
	}
}
