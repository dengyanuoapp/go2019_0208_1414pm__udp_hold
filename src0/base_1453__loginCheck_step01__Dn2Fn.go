package main

/*
the Dn2Fn , step 01 , "Dn" gen a tokenL(A) and send to "Fn" , record in CmdMap(S) , sending : tokenL(A),-
the Fn2Dn , step 02 , "Fn" gen a tokenR(B) and send to "Dn" , record in CmdMap(C) , sending : tokenL(B),tokenR(A)
the Dn2Fn , step 03 , "Dn" check CmdMap(S) tokenR == A and delay less than 10s, then accept it : clear DataMap, gen new DataMap[ID128(Fn)]
the Fn2Dn , step 04 , "Fn" check CmdMap(C) tokenR == B and delay less than 10s, then accept it : clear DataMap, gen new DataMap[ID128(Dn)]

note :
1. setp 01 precess in :

2. CmdMap(S) and CmdMap(C) is the same, only difference working in Client and Server

*/
