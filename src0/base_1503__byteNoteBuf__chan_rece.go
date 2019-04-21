package main

func (___Vbnb *_TbyteNoteBuf) _FbyteNoteBuf__1300201x__chan_rece__default() {
	select {
	case __Vb := <-___Vbnb.bnbCHinI:
		if nil != ___Vbnb.bnbCHoutLO1 {
			(*___Vbnb.bnbCHoutLO1) <- __Vb
		}
		if nil != ___Vbnb.bnbCHoutLO2 {
			(*___Vbnb.bnbCHoutLO2) <- __Vb
		}
		if nil != ___Vbnb.bnbCHoutLO3 {
			(*___Vbnb.bnbCHoutLO3) <- __Vb
		}
		if nil != ___Vbnb.bnbCHoutLO4 {
			(*___Vbnb.bnbCHoutLO4) <- __Vb
		}
		if nil != ___Vbnb.bnbCHoutLO5 {
			(*___Vbnb.bnbCHoutLO5) <- __Vb
		}
	}
}
