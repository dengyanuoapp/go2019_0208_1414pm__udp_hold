package main

func (___Vbnb *_TbyteNoteBuf) _FbyteNoteBuf__1300201x__chan_rece__default() {
	select {
	case __Vb := <-___Vbnb.bnbCHinI:

		_CFpfN(" 838191 01 _TbyteNoteBuf ")

		if nil != ___Vbnb.bnbCHoutLO1 {
			_CFpfN(" 838191 11 _TbyteNoteBuf out : %d , %d", len(*___Vbnb.bnbCHoutLO1), cap(*___Vbnb.bnbCHoutLO1))
			(*___Vbnb.bnbCHoutLO1) <- __Vb
		}
		if nil != ___Vbnb.bnbCHoutLO2 {
			_CFpfN(" 838191 12 _TbyteNoteBuf out : %d , %d", len(*___Vbnb.bnbCHoutLO1), cap(*___Vbnb.bnbCHoutLO1))
			(*___Vbnb.bnbCHoutLO2) <- __Vb
		}
		if nil != ___Vbnb.bnbCHoutLO3 {
			_CFpfN(" 838191 13 _TbyteNoteBuf out : %d , %d", len(*___Vbnb.bnbCHoutLO1), cap(*___Vbnb.bnbCHoutLO1))
			(*___Vbnb.bnbCHoutLO3) <- __Vb
		}
		if nil != ___Vbnb.bnbCHoutLO4 {
			_CFpfN(" 838191 14 _TbyteNoteBuf out : %d , %d", len(*___Vbnb.bnbCHoutLO1), cap(*___Vbnb.bnbCHoutLO1))
			(*___Vbnb.bnbCHoutLO4) <- __Vb
		}
		if nil != ___Vbnb.bnbCHoutLO5 {
			_CFpfN(" 838191 15 _TbyteNoteBuf out : %d , %d", len(*___Vbnb.bnbCHoutLO1), cap(*___Vbnb.bnbCHoutLO1))
			(*___Vbnb.bnbCHoutLO5) <- __Vb
		}
	}
}
