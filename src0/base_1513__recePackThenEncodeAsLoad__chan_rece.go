package main

func (___Vpel *_TrecePackThenEncodeAsLoad) _FrecePackThenEncodeAsLoad__1400201x__chan_rece__default() {
	for {

		select {

		case __Vp := <-___Vpel.pelCHinI:

			_NpfN(" 638191 01 _TrecePackThenEncodeAsLoad ")

			if nil != ___Vpel.pelCHoutLO {
				//_CFpfN(" 638191 11 _TrecePackThenEncodeAsLoad out : %d , %d", len(*___Vpel.pelCHoutLO), cap(*___Vpel.pelCHoutLO))
				(*___Vpel.pelCHoutLO) <- __Vp
				//_CFpfN(" 638191 21 _TrecePackThenEncodeAsLoad out : %d , %d", len(*___Vpel.pelCHoutLO), cap(*___Vpel.pelCHoutLO))
			}
		}
	}
}
