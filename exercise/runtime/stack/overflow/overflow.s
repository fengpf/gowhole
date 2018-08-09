"".val.String STEXT size=360 args=0x10 locals=0x90
	0x0000 00000 (main.go:14)	TEXT	"".val.String(SB), $144-16
	0x0000 00000 (main.go:14)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:14)	LEAQ	-16(SP), AX
	0x000e 00014 (main.go:14)	CMPQ	AX, 16(CX)
	0x0012 00018 (main.go:14)	JLS	350
	0x0018 00024 (main.go:14)	SUBQ	$144, SP
	0x001f 00031 (main.go:14)	MOVQ	BP, 136(SP)
	0x0027 00039 (main.go:14)	LEAQ	136(SP), BP
	0x002f 00047 (main.go:14)	FUNCDATA	$0, gclocals·d8b28f51bb91e05d264803f0f600a200(SB)
	0x002f 00047 (main.go:14)	FUNCDATA	$1, gclocals·b0938e6026b5c81628342d8ac78cb7a4(SB)
	0x002f 00047 (main.go:14)	XORPS	X0, X0
	0x0032 00050 (main.go:14)	MOVUPS	X0, "".~r0+152(SP)
	0x003a 00058 (main.go:15)	PCDATA	$0, $0
	0x003a 00058 (main.go:15)	CALL	"".val.String(SB)
	0x003f 00063 (main.go:15)	MOVQ	8(SP), AX
	0x0044 00068 (main.go:15)	MOVQ	(SP), CX
	0x0048 00072 (main.go:15)	MOVQ	CX, ""..autotmp_3+80(SP)
	0x004d 00077 (main.go:15)	MOVQ	AX, ""..autotmp_3+88(SP)
	0x0052 00082 (main.go:15)	XORPS	X0, X0
	0x0055 00085 (main.go:15)	MOVUPS	X0, ""..autotmp_2+96(SP)
	0x005a 00090 (main.go:15)	LEAQ	""..autotmp_2+96(SP), AX
	0x005f 00095 (main.go:15)	MOVQ	AX, ""..autotmp_6+56(SP)
	0x0064 00100 (main.go:15)	LEAQ	type.string(SB), AX
	0x006b 00107 (main.go:15)	MOVQ	AX, (SP)
	0x006f 00111 (main.go:15)	LEAQ	""..autotmp_3+80(SP), AX
	0x0074 00116 (main.go:15)	MOVQ	AX, 8(SP)
	0x0079 00121 (main.go:15)	PCDATA	$0, $1
	0x0079 00121 (main.go:15)	CALL	runtime.convT2Estring(SB)
	0x007e 00126 (main.go:15)	MOVQ	24(SP), AX
	0x0083 00131 (main.go:15)	MOVQ	16(SP), CX
	0x0088 00136 (main.go:15)	MOVQ	CX, ""..autotmp_7+64(SP)
	0x008d 00141 (main.go:15)	MOVQ	AX, ""..autotmp_7+72(SP)
	0x0092 00146 (main.go:15)	MOVQ	""..autotmp_6+56(SP), DX
	0x0097 00151 (main.go:15)	TESTB	AL, (DX)
	0x0099 00153 (main.go:15)	MOVQ	CX, (DX)
	0x009c 00156 (main.go:15)	MOVL	runtime.writeBarrier(SB), CX
	0x00a2 00162 (main.go:15)	LEAQ	8(DX), DI
	0x00a6 00166 (main.go:15)	TESTL	CX, CX
	0x00a8 00168 (main.go:15)	JNE	340
	0x00ae 00174 (main.go:15)	JMP	176
	0x00b0 00176 (main.go:15)	MOVQ	AX, 8(DX)
	0x00b4 00180 (main.go:15)	JMP	182
	0x00b6 00182 (main.go:15)	MOVQ	""..autotmp_6+56(SP), AX
	0x00bb 00187 (main.go:15)	TESTB	AL, (AX)
	0x00bd 00189 (main.go:15)	JMP	191
	0x00bf 00191 (main.go:15)	MOVQ	AX, ""..autotmp_5+112(SP)
	0x00c4 00196 (main.go:15)	MOVQ	$1, ""..autotmp_5+120(SP)
	0x00cd 00205 (main.go:15)	MOVQ	$1, ""..autotmp_5+128(SP)
	0x00d9 00217 (main.go:15)	MOVQ	AX, (SP)
	0x00dd 00221 (main.go:15)	MOVQ	$1, 8(SP)
	0x00e6 00230 (main.go:15)	MOVQ	$1, 16(SP)
	0x00ef 00239 (main.go:15)	PCDATA	$0, $2
	0x00ef 00239 (main.go:15)	CALL	fmt.Println(SB)
	0x00f4 00244 (main.go:16)	MOVQ	"".i(SB), AX
	0x00fb 00251 (main.go:16)	MOVQ	AX, ""..autotmp_4+48(SP)
	0x0100 00256 (main.go:16)	INCQ	AX
	0x0103 00259 (main.go:16)	MOVQ	AX, "".i(SB)
	0x010a 00266 (main.go:17)	PCDATA	$0, $0
	0x010a 00266 (main.go:17)	CALL	runtime.printlock(SB)
	0x010f 00271 (main.go:17)	MOVQ	"".i(SB), AX
	0x0116 00278 (main.go:17)	MOVQ	AX, (SP)
	0x011a 00282 (main.go:17)	PCDATA	$0, $0
	0x011a 00282 (main.go:17)	CALL	runtime.printint(SB)
	0x011f 00287 (main.go:17)	PCDATA	$0, $0
	0x011f 00287 (main.go:17)	CALL	runtime.printnl(SB)
	0x0124 00292 (main.go:17)	PCDATA	$0, $0
	0x0124 00292 (main.go:17)	CALL	runtime.printunlock(SB)
	0x0129 00297 (main.go:18)	LEAQ	go.string."aaa"(SB), AX
	0x0130 00304 (main.go:18)	MOVQ	AX, "".~r0+152(SP)
	0x0138 00312 (main.go:18)	MOVQ	$3, "".~r0+160(SP)
	0x0144 00324 (main.go:18)	MOVQ	136(SP), BP
	0x014c 00332 (main.go:18)	ADDQ	$144, SP
	0x0153 00339 (main.go:18)	RET
	0x0154 00340 (main.go:15)	CALL	runtime.gcWriteBarrier(SB)
	0x0159 00345 (main.go:15)	JMP	182
	0x015e 00350 (main.go:15)	NOP
	0x015e 00350 (main.go:14)	PCDATA	$0, $-1
	0x015e 00350 (main.go:14)	CALL	runtime.morestack_noctxt(SB)
	0x0163 00355 (main.go:14)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 8d 44 24 f0 48 3b  eH..%....H.D$.H;
	0x0010 41 10 0f 86 46 01 00 00 48 81 ec 90 00 00 00 48  A...F...H......H
	0x0020 89 ac 24 88 00 00 00 48 8d ac 24 88 00 00 00 0f  ..$....H..$.....
	0x0030 57 c0 0f 11 84 24 98 00 00 00 e8 00 00 00 00 48  W....$.........H
	0x0040 8b 44 24 08 48 8b 0c 24 48 89 4c 24 50 48 89 44  .D$.H..$H.L$PH.D
	0x0050 24 58 0f 57 c0 0f 11 44 24 60 48 8d 44 24 60 48  $X.W...D$`H.D$`H
	0x0060 89 44 24 38 48 8d 05 00 00 00 00 48 89 04 24 48  .D$8H......H..$H
	0x0070 8d 44 24 50 48 89 44 24 08 e8 00 00 00 00 48 8b  .D$PH.D$......H.
	0x0080 44 24 18 48 8b 4c 24 10 48 89 4c 24 40 48 89 44  D$.H.L$.H.L$@H.D
	0x0090 24 48 48 8b 54 24 38 84 02 48 89 0a 8b 0d 00 00  $HH.T$8..H......
	0x00a0 00 00 48 8d 7a 08 85 c9 0f 85 a6 00 00 00 eb 00  ..H.z...........
	0x00b0 48 89 42 08 eb 00 48 8b 44 24 38 84 00 eb 00 48  H.B...H.D$8....H
	0x00c0 89 44 24 70 48 c7 44 24 78 01 00 00 00 48 c7 84  .D$pH.D$x....H..
	0x00d0 24 80 00 00 00 01 00 00 00 48 89 04 24 48 c7 44  $........H..$H.D
	0x00e0 24 08 01 00 00 00 48 c7 44 24 10 01 00 00 00 e8  $.....H.D$......
	0x00f0 00 00 00 00 48 8b 05 00 00 00 00 48 89 44 24 30  ....H......H.D$0
	0x0100 48 ff c0 48 89 05 00 00 00 00 e8 00 00 00 00 48  H..H...........H
	0x0110 8b 05 00 00 00 00 48 89 04 24 e8 00 00 00 00 e8  ......H..$......
	0x0120 00 00 00 00 e8 00 00 00 00 48 8d 05 00 00 00 00  .........H......
	0x0130 48 89 84 24 98 00 00 00 48 c7 84 24 a0 00 00 00  H..$....H..$....
	0x0140 03 00 00 00 48 8b ac 24 88 00 00 00 48 81 c4 90  ....H..$....H...
	0x0150 00 00 00 c3 e8 00 00 00 00 e9 58 ff ff ff e8 00  ..........X.....
	0x0160 00 00 00 e9 98 fe ff ff                          ........
	rel 5+4 t=16 TLS+0
	rel 59+4 t=8 "".val.String+0
	rel 103+4 t=15 type.string+0
	rel 122+4 t=8 runtime.convT2Estring+0
	rel 158+4 t=15 runtime.writeBarrier+0
	rel 240+4 t=8 fmt.Println+0
	rel 247+4 t=15 "".i+0
	rel 262+4 t=15 "".i+0
	rel 267+4 t=8 runtime.printlock+0
	rel 274+4 t=15 "".i+0
	rel 283+4 t=8 runtime.printint+0
	rel 288+4 t=8 runtime.printnl+0
	rel 293+4 t=8 runtime.printunlock+0
	rel 300+4 t=15 go.string."aaa"+0
	rel 341+4 t=8 runtime.gcWriteBarrier+0
	rel 351+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=584 args=0x0 locals=0xe0
	0x0000 00000 (main.go:23)	TEXT	"".main(SB), $224-0
	0x0000 00000 (main.go:23)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:23)	LEAQ	-96(SP), AX
	0x000e 00014 (main.go:23)	CMPQ	AX, 16(CX)
	0x0012 00018 (main.go:23)	JLS	574
	0x0018 00024 (main.go:23)	SUBQ	$224, SP
	0x001f 00031 (main.go:23)	MOVQ	BP, 216(SP)
	0x0027 00039 (main.go:23)	LEAQ	216(SP), BP
	0x002f 00047 (main.go:23)	FUNCDATA	$0, gclocals·f6bd6b3389b872033d462029172c8612(SB)
	0x002f 00047 (main.go:23)	FUNCDATA	$1, gclocals·3e9d367ad8e894cbbe1024f46b88885d(SB)
	0x002f 00047 (main.go:26)	PCDATA	$0, $0
	0x002f 00047 (main.go:26)	CALL	"".val.String(SB)
	0x0034 00052 (main.go:26)	MOVQ	(SP), AX
	0x0038 00056 (main.go:26)	MOVQ	8(SP), CX
	0x003d 00061 (main.go:26)	MOVQ	AX, ""..autotmp_3+160(SP)
	0x0045 00069 (main.go:26)	MOVQ	CX, ""..autotmp_3+168(SP)
	0x004d 00077 (main.go:26)	PCDATA	$0, $1
	0x004d 00077 (main.go:26)	CALL	runtime.printlock(SB)
	0x0052 00082 (main.go:26)	MOVQ	""..autotmp_3+168(SP), AX
	0x005a 00090 (main.go:26)	MOVQ	""..autotmp_3+160(SP), CX
	0x0062 00098 (main.go:26)	MOVQ	CX, (SP)
	0x0066 00102 (main.go:26)	MOVQ	AX, 8(SP)
	0x006b 00107 (main.go:26)	PCDATA	$0, $0
	0x006b 00107 (main.go:26)	CALL	runtime.printstring(SB)
	0x0070 00112 (main.go:26)	PCDATA	$0, $0
	0x0070 00112 (main.go:26)	CALL	runtime.printnl(SB)
	0x0075 00117 (main.go:26)	PCDATA	$0, $0
	0x0075 00117 (main.go:26)	CALL	runtime.printunlock(SB)
	0x007a 00122 (main.go:27)	LEAQ	go.itab."".val,"".Stringer(SB), AX
	0x0081 00129 (main.go:27)	MOVQ	AX, "".value+80(SP)
	0x0086 00134 (main.go:27)	LEAQ	runtime.zerobase(SB), AX
	0x008d 00141 (main.go:27)	MOVQ	AX, "".value+88(SP)
	0x0092 00146 (main.go:28)	MOVQ	"".value+88(SP), AX
	0x0097 00151 (main.go:28)	MOVQ	"".value+80(SP), CX
	0x009c 00156 (main.go:28)	MOVQ	CX, ""..autotmp_7+112(SP)
	0x00a1 00161 (main.go:28)	MOVQ	AX, ""..autotmp_7+120(SP)
	0x00a6 00166 (main.go:28)	TESTQ	CX, CX
	0x00a9 00169 (main.go:28)	JNE	176
	0x00ab 00171 (main.go:28)	JMP	572
	0x00b0 00176 (main.go:28)	TESTB	AL, (CX)
	0x00b2 00178 (main.go:28)	MOVL	16(CX), AX
	0x00b5 00181 (main.go:28)	MOVL	AX, ""..autotmp_9+68(SP)
	0x00b9 00185 (main.go:28)	XORPS	X0, X0
	0x00bc 00188 (main.go:28)	MOVUPS	X0, "".str+96(SP)
	0x00c1 00193 (main.go:28)	MOVQ	""..autotmp_7+120(SP), AX
	0x00c6 00198 (main.go:28)	MOVQ	""..autotmp_7+112(SP), CX
	0x00cb 00203 (main.go:28)	LEAQ	type."".Stringer(SB), DX
	0x00d2 00210 (main.go:28)	MOVQ	DX, (SP)
	0x00d6 00214 (main.go:28)	MOVQ	CX, 8(SP)
	0x00db 00219 (main.go:28)	MOVQ	AX, 16(SP)
	0x00e0 00224 (main.go:28)	PCDATA	$0, $0
	0x00e0 00224 (main.go:28)	CALL	runtime.assertI2I2(SB)
	0x00e5 00229 (main.go:28)	MOVQ	32(SP), AX
	0x00ea 00234 (main.go:28)	MOVQ	24(SP), CX
	0x00ef 00239 (main.go:28)	MOVBLZX	40(SP), DX
	0x00f4 00244 (main.go:28)	MOVQ	CX, "".str+96(SP)
	0x00f9 00249 (main.go:28)	MOVQ	AX, "".str+104(SP)
	0x00fe 00254 (main.go:28)	MOVB	DL, ""..autotmp_8+67(SP)
	0x0102 00258 (main.go:28)	TESTB	DL, DL
	0x0104 00260 (main.go:28)	JNE	267
	0x0106 00262 (main.go:28)	JMP	568
	0x010b 00267 (main.go:31)	JMP	269
	0x010d 00269 (main.go:32)	MOVQ	"".str+96(SP), AX
	0x0112 00274 (main.go:32)	TESTB	AL, (AX)
	0x0114 00276 (main.go:32)	MOVQ	"".str+104(SP), CX
	0x0119 00281 (main.go:32)	MOVQ	24(AX), AX
	0x011d 00285 (main.go:32)	MOVQ	CX, (SP)
	0x0121 00289 (main.go:32)	PCDATA	$0, $0
	0x0121 00289 (main.go:32)	CALL	AX
	0x0123 00291 (main.go:32)	MOVQ	16(SP), AX
	0x0128 00296 (main.go:32)	MOVQ	8(SP), CX
	0x012d 00301 (main.go:32)	MOVQ	CX, ""..autotmp_6+128(SP)
	0x0135 00309 (main.go:32)	MOVQ	AX, ""..autotmp_6+136(SP)
	0x013d 00317 (main.go:32)	XORPS	X0, X0
	0x0140 00320 (main.go:32)	MOVUPS	X0, ""..autotmp_5+144(SP)
	0x0148 00328 (main.go:32)	LEAQ	""..autotmp_5+144(SP), AX
	0x0150 00336 (main.go:32)	MOVQ	AX, ""..autotmp_11+72(SP)
	0x0155 00341 (main.go:32)	LEAQ	type.string(SB), AX
	0x015c 00348 (main.go:32)	MOVQ	AX, (SP)
	0x0160 00352 (main.go:32)	LEAQ	""..autotmp_6+128(SP), AX
	0x0168 00360 (main.go:32)	MOVQ	AX, 8(SP)
	0x016d 00365 (main.go:32)	PCDATA	$0, $2
	0x016d 00365 (main.go:32)	CALL	runtime.convT2Estring(SB)
	0x0172 00370 (main.go:32)	MOVQ	16(SP), AX
	0x0177 00375 (main.go:32)	MOVQ	24(SP), CX
	0x017c 00380 (main.go:32)	MOVQ	AX, ""..autotmp_12+176(SP)
	0x0184 00388 (main.go:32)	MOVQ	CX, ""..autotmp_12+184(SP)
	0x018c 00396 (main.go:32)	MOVQ	""..autotmp_11+72(SP), DX
	0x0191 00401 (main.go:32)	TESTB	AL, (DX)
	0x0193 00403 (main.go:32)	MOVQ	AX, (DX)
	0x0196 00406 (main.go:32)	MOVL	runtime.writeBarrier(SB), AX
	0x019c 00412 (main.go:32)	LEAQ	8(DX), DI
	0x01a0 00416 (main.go:32)	TESTL	AX, AX
	0x01a2 00418 (main.go:32)	JNE	555
	0x01a8 00424 (main.go:32)	JMP	426
	0x01aa 00426 (main.go:32)	MOVQ	CX, 8(DX)
	0x01ae 00430 (main.go:32)	JMP	432
	0x01b0 00432 (main.go:32)	MOVQ	""..autotmp_11+72(SP), AX
	0x01b5 00437 (main.go:32)	TESTB	AL, (AX)
	0x01b7 00439 (main.go:32)	JMP	441
	0x01b9 00441 (main.go:32)	MOVQ	AX, ""..autotmp_10+192(SP)
	0x01c1 00449 (main.go:32)	MOVQ	$1, ""..autotmp_10+200(SP)
	0x01cd 00461 (main.go:32)	MOVQ	$1, ""..autotmp_10+208(SP)
	0x01d9 00473 (main.go:32)	LEAQ	go.string."str(%v) is Stringer\n"(SB), AX
	0x01e0 00480 (main.go:32)	MOVQ	AX, (SP)
	0x01e4 00484 (main.go:32)	MOVQ	$20, 8(SP)
	0x01ed 00493 (main.go:32)	MOVQ	""..autotmp_10+192(SP), AX
	0x01f5 00501 (main.go:32)	MOVQ	""..autotmp_10+200(SP), CX
	0x01fd 00509 (main.go:32)	MOVQ	""..autotmp_10+208(SP), DX
	0x0205 00517 (main.go:32)	MOVQ	AX, 16(SP)
	0x020a 00522 (main.go:32)	MOVQ	CX, 24(SP)
	0x020f 00527 (main.go:32)	MOVQ	DX, 32(SP)
	0x0214 00532 (main.go:32)	PCDATA	$0, $3
	0x0214 00532 (main.go:32)	CALL	fmt.Printf(SB)
	0x0219 00537 (main.go:28)	JMP	539
	0x021b 00539 (<unknown line number>)	MOVQ	216(SP), BP
	0x0223 00547 (<unknown line number>)	ADDQ	$224, SP
	0x022a 00554 (<unknown line number>)	RET
	0x022b 00555 (<unknown line number>)	MOVQ	CX, AX
	0x022e 00558 (main.go:32)	CALL	runtime.gcWriteBarrier(SB)
	0x0233 00563 (main.go:32)	JMP	432
	0x0238 00568 (main.go:28)	JMP	570
	0x023a 00570 (main.go:28)	JMP	539
	0x023c 00572 (main.go:28)	JMP	570
	0x023e 00574 (main.go:28)	NOP
	0x023e 00574 (main.go:23)	PCDATA	$0, $-1
	0x023e 00574 (main.go:23)	CALL	runtime.morestack_noctxt(SB)
	0x0243 00579 (main.go:23)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 8d 44 24 a0 48 3b  eH..%....H.D$.H;
	0x0010 41 10 0f 86 26 02 00 00 48 81 ec e0 00 00 00 48  A...&...H......H
	0x0020 89 ac 24 d8 00 00 00 48 8d ac 24 d8 00 00 00 e8  ..$....H..$.....
	0x0030 00 00 00 00 48 8b 04 24 48 8b 4c 24 08 48 89 84  ....H..$H.L$.H..
	0x0040 24 a0 00 00 00 48 89 8c 24 a8 00 00 00 e8 00 00  $....H..$.......
	0x0050 00 00 48 8b 84 24 a8 00 00 00 48 8b 8c 24 a0 00  ..H..$....H..$..
	0x0060 00 00 48 89 0c 24 48 89 44 24 08 e8 00 00 00 00  ..H..$H.D$......
	0x0070 e8 00 00 00 00 e8 00 00 00 00 48 8d 05 00 00 00  ..........H.....
	0x0080 00 48 89 44 24 50 48 8d 05 00 00 00 00 48 89 44  .H.D$PH......H.D
	0x0090 24 58 48 8b 44 24 58 48 8b 4c 24 50 48 89 4c 24  $XH.D$XH.L$PH.L$
	0x00a0 70 48 89 44 24 78 48 85 c9 75 05 e9 8c 01 00 00  pH.D$xH..u......
	0x00b0 84 01 8b 41 10 89 44 24 44 0f 57 c0 0f 11 44 24  ...A..D$D.W...D$
	0x00c0 60 48 8b 44 24 78 48 8b 4c 24 70 48 8d 15 00 00  `H.D$xH.L$pH....
	0x00d0 00 00 48 89 14 24 48 89 4c 24 08 48 89 44 24 10  ..H..$H.L$.H.D$.
	0x00e0 e8 00 00 00 00 48 8b 44 24 20 48 8b 4c 24 18 0f  .....H.D$ H.L$..
	0x00f0 b6 54 24 28 48 89 4c 24 60 48 89 44 24 68 88 54  .T$(H.L$`H.D$h.T
	0x0100 24 43 84 d2 75 05 e9 2d 01 00 00 eb 00 48 8b 44  $C..u..-.....H.D
	0x0110 24 60 84 00 48 8b 4c 24 68 48 8b 40 18 48 89 0c  $`..H.L$hH.@.H..
	0x0120 24 ff d0 48 8b 44 24 10 48 8b 4c 24 08 48 89 8c  $..H.D$.H.L$.H..
	0x0130 24 80 00 00 00 48 89 84 24 88 00 00 00 0f 57 c0  $....H..$.....W.
	0x0140 0f 11 84 24 90 00 00 00 48 8d 84 24 90 00 00 00  ...$....H..$....
	0x0150 48 89 44 24 48 48 8d 05 00 00 00 00 48 89 04 24  H.D$HH......H..$
	0x0160 48 8d 84 24 80 00 00 00 48 89 44 24 08 e8 00 00  H..$....H.D$....
	0x0170 00 00 48 8b 44 24 10 48 8b 4c 24 18 48 89 84 24  ..H.D$.H.L$.H..$
	0x0180 b0 00 00 00 48 89 8c 24 b8 00 00 00 48 8b 54 24  ....H..$....H.T$
	0x0190 48 84 02 48 89 02 8b 05 00 00 00 00 48 8d 7a 08  H..H........H.z.
	0x01a0 85 c0 0f 85 83 00 00 00 eb 00 48 89 4a 08 eb 00  ..........H.J...
	0x01b0 48 8b 44 24 48 84 00 eb 00 48 89 84 24 c0 00 00  H.D$H....H..$...
	0x01c0 00 48 c7 84 24 c8 00 00 00 01 00 00 00 48 c7 84  .H..$........H..
	0x01d0 24 d0 00 00 00 01 00 00 00 48 8d 05 00 00 00 00  $........H......
	0x01e0 48 89 04 24 48 c7 44 24 08 14 00 00 00 48 8b 84  H..$H.D$.....H..
	0x01f0 24 c0 00 00 00 48 8b 8c 24 c8 00 00 00 48 8b 94  $....H..$....H..
	0x0200 24 d0 00 00 00 48 89 44 24 10 48 89 4c 24 18 48  $....H.D$.H.L$.H
	0x0210 89 54 24 20 e8 00 00 00 00 eb 00 48 8b ac 24 d8  .T$ .......H..$.
	0x0220 00 00 00 48 81 c4 e0 00 00 00 c3 48 89 c8 e8 00  ...H.......H....
	0x0230 00 00 00 e9 78 ff ff ff eb 00 eb df eb fc e8 00  ....x...........
	0x0240 00 00 00 e9 b8 fd ff ff                          ........
	rel 5+4 t=16 TLS+0
	rel 48+4 t=8 "".val.String+0
	rel 78+4 t=8 runtime.printlock+0
	rel 108+4 t=8 runtime.printstring+0
	rel 113+4 t=8 runtime.printnl+0
	rel 118+4 t=8 runtime.printunlock+0
	rel 125+4 t=15 go.itab."".val,"".Stringer+0
	rel 137+4 t=15 runtime.zerobase+0
	rel 206+4 t=15 type."".Stringer+0
	rel 225+4 t=8 runtime.assertI2I2+0
	rel 289+0 t=11 +0
	rel 344+4 t=15 type.string+0
	rel 366+4 t=8 runtime.convT2Estring+0
	rel 408+4 t=15 runtime.writeBarrier+0
	rel 476+4 t=15 go.string."str(%v) is Stringer\n"+0
	rel 533+4 t=8 fmt.Printf+0
	rel 559+4 t=8 runtime.gcWriteBarrier+0
	rel 575+4 t=8 runtime.morestack_noctxt+0
"".init STEXT size=104 args=0x0 locals=0x8
	0x0000 00000 (<autogenerated>:1)	TEXT	"".init(SB), $8-0
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	97
	0x000f 00015 (<autogenerated>:1)	SUBQ	$8, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x0017 00023 (<autogenerated>:1)	LEAQ	(SP), BP
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	MOVBLZX	"".initdone·(SB), AX
	0x0022 00034 (<autogenerated>:1)	CMPB	AL, $1
	0x0024 00036 (<autogenerated>:1)	JHI	40
	0x0026 00038 (<autogenerated>:1)	JMP	49
	0x0028 00040 (<autogenerated>:1)	MOVQ	(SP), BP
	0x002c 00044 (<autogenerated>:1)	ADDQ	$8, SP
	0x0030 00048 (<autogenerated>:1)	RET
	0x0031 00049 (<autogenerated>:1)	MOVBLZX	"".initdone·(SB), AX
	0x0038 00056 (<autogenerated>:1)	CMPB	AL, $1
	0x003a 00058 (<autogenerated>:1)	JEQ	62
	0x003c 00060 (<autogenerated>:1)	JMP	69
	0x003e 00062 (<autogenerated>:1)	PCDATA	$0, $0
	0x003e 00062 (<autogenerated>:1)	CALL	runtime.throwinit(SB)
	0x0043 00067 (<autogenerated>:1)	UNDEF
	0x0045 00069 (<autogenerated>:1)	MOVB	$1, "".initdone·(SB)
	0x004c 00076 (<autogenerated>:1)	PCDATA	$0, $0
	0x004c 00076 (<autogenerated>:1)	CALL	fmt.init(SB)
	0x0051 00081 (<autogenerated>:1)	MOVB	$2, "".initdone·(SB)
	0x0058 00088 (<autogenerated>:1)	MOVQ	(SP), BP
	0x005c 00092 (<autogenerated>:1)	ADDQ	$8, SP
	0x0060 00096 (<autogenerated>:1)	RET
	0x0061 00097 (<autogenerated>:1)	NOP
	0x0061 00097 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0061 00097 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0066 00102 (<autogenerated>:1)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 52 48  eH..%....H;a.vRH
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 0f b6 05 00 00  ...H.,$H.,$.....
	0x0020 00 00 3c 01 77 02 eb 09 48 8b 2c 24 48 83 c4 08  ..<.w...H.,$H...
	0x0030 c3 0f b6 05 00 00 00 00 3c 01 74 02 eb 07 e8 00  ........<.t.....
	0x0040 00 00 00 0f 0b c6 05 00 00 00 00 01 e8 00 00 00  ................
	0x0050 00 c6 05 00 00 00 00 02 48 8b 2c 24 48 83 c4 08  ........H.,$H...
	0x0060 c3 e8 00 00 00 00 eb 98                          ........
	rel 5+4 t=16 TLS+0
	rel 30+4 t=15 "".initdone·+0
	rel 52+4 t=15 "".initdone·+0
	rel 63+4 t=8 runtime.throwinit+0
	rel 71+4 t=15 "".initdone·+-1
	rel 77+4 t=8 fmt.init+0
	rel 83+4 t=15 "".initdone·+-1
	rel 98+4 t=8 runtime.morestack_noctxt+0
"".Stringer.String STEXT dupok size=130 args=0x20 locals=0x30
	0x0000 00000 (<autogenerated>:1)	TEXT	"".Stringer.String(SB), DUPOK|WRAPPER, $48-32
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	108
	0x000f 00015 (<autogenerated>:1)	SUBQ	$48, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, 40(SP)
	0x0018 00024 (<autogenerated>:1)	LEAQ	40(SP), BP
	0x001d 00029 (<autogenerated>:1)	MOVQ	32(CX), BX
	0x0021 00033 (<autogenerated>:1)	TESTQ	BX, BX
	0x0024 00036 (<autogenerated>:1)	JNE	115
	0x0026 00038 (<autogenerated>:1)	NOP
	0x0026 00038 (<autogenerated>:1)	FUNCDATA	$0, gclocals·c55e845a0a62e9baae6c740db5a20866(SB)
	0x0026 00038 (<autogenerated>:1)	FUNCDATA	$1, gclocals·2589ca35330fc0fce83503f4569854a0(SB)
	0x0026 00038 (<autogenerated>:1)	XORPS	X0, X0
	0x0029 00041 (<autogenerated>:1)	MOVUPS	X0, "".~r1+72(SP)
	0x002e 00046 (<autogenerated>:1)	MOVQ	""..this+56(SP), AX
	0x0033 00051 (<autogenerated>:1)	TESTB	AL, (AX)
	0x0035 00053 (<autogenerated>:1)	MOVQ	24(AX), AX
	0x0039 00057 (<autogenerated>:1)	MOVQ	""..this+64(SP), CX
	0x003e 00062 (<autogenerated>:1)	MOVQ	CX, (SP)
	0x0042 00066 (<autogenerated>:1)	PCDATA	$0, $1
	0x0042 00066 (<autogenerated>:1)	CALL	AX
	0x0044 00068 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x0049 00073 (<autogenerated>:1)	MOVQ	16(SP), CX
	0x004e 00078 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_2+24(SP)
	0x0053 00083 (<autogenerated>:1)	MOVQ	CX, ""..autotmp_2+32(SP)
	0x0058 00088 (<autogenerated>:1)	MOVQ	AX, "".~r1+72(SP)
	0x005d 00093 (<autogenerated>:1)	MOVQ	CX, "".~r1+80(SP)
	0x0062 00098 (<autogenerated>:1)	MOVQ	40(SP), BP
	0x0067 00103 (<autogenerated>:1)	ADDQ	$48, SP
	0x006b 00107 (<autogenerated>:1)	RET
	0x006c 00108 (<autogenerated>:1)	NOP
	0x006c 00108 (<autogenerated>:1)	PCDATA	$0, $-1
	0x006c 00108 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0071 00113 (<autogenerated>:1)	JMP	0
	0x0073 00115 (<autogenerated>:1)	LEAQ	56(SP), DI
	0x0078 00120 (<autogenerated>:1)	CMPQ	(BX), DI
	0x007b 00123 (<autogenerated>:1)	JNE	38
	0x007d 00125 (<autogenerated>:1)	MOVQ	SP, (BX)
	0x0080 00128 (<autogenerated>:1)	JMP	38
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 5d 48  eH..%....H;a.v]H
	0x0010 83 ec 30 48 89 6c 24 28 48 8d 6c 24 28 48 8b 59  ..0H.l$(H.l$(H.Y
	0x0020 20 48 85 db 75 4d 0f 57 c0 0f 11 44 24 48 48 8b   H..uM.W...D$HH.
	0x0030 44 24 38 84 00 48 8b 40 18 48 8b 4c 24 40 48 89  D$8..H.@.H.L$@H.
	0x0040 0c 24 ff d0 48 8b 44 24 08 48 8b 4c 24 10 48 89  .$..H.D$.H.L$.H.
	0x0050 44 24 18 48 89 4c 24 20 48 89 44 24 48 48 89 4c  D$.H.L$ H.D$HH.L
	0x0060 24 50 48 8b 6c 24 28 48 83 c4 30 c3 e8 00 00 00  $PH.l$(H..0.....
	0x0070 00 eb 8d 48 8d 7c 24 38 48 39 3b 75 a9 48 89 23  ...H.|$8H9;u.H.#
	0x0080 eb a4                                            ..
	rel 5+4 t=16 TLS+0
	rel 66+0 t=11 +0
	rel 109+4 t=8 runtime.morestack_noctxt+0
"".(*val).String STEXT dupok size=138 args=0x18 locals=0x28
	0x0000 00000 (<autogenerated>:1)	TEXT	"".(*val).String(SB), DUPOK|WRAPPER, $40-24
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	116
	0x000f 00015 (<autogenerated>:1)	SUBQ	$40, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, 32(SP)
	0x0018 00024 (<autogenerated>:1)	LEAQ	32(SP), BP
	0x001d 00029 (<autogenerated>:1)	MOVQ	32(CX), BX
	0x0021 00033 (<autogenerated>:1)	TESTQ	BX, BX
	0x0024 00036 (<autogenerated>:1)	JNE	123
	0x0026 00038 (<autogenerated>:1)	NOP
	0x0026 00038 (<autogenerated>:1)	FUNCDATA	$0, gclocals·e6397a44f8e1b6e77d0f200b4fba5269(SB)
	0x0026 00038 (<autogenerated>:1)	FUNCDATA	$1, gclocals·2589ca35330fc0fce83503f4569854a0(SB)
	0x0026 00038 (<autogenerated>:1)	XORPS	X0, X0
	0x0029 00041 (<autogenerated>:1)	MOVUPS	X0, "".~r0+56(SP)
	0x002e 00046 (<autogenerated>:1)	MOVQ	""..this+48(SP), AX
	0x0033 00051 (<autogenerated>:1)	TESTQ	AX, AX
	0x0036 00054 (<autogenerated>:1)	JNE	58
	0x0038 00056 (<autogenerated>:1)	JMP	109
	0x003a 00058 (<autogenerated>:1)	MOVQ	""..this+48(SP), AX
	0x003f 00063 (<autogenerated>:1)	TESTB	AL, (AX)
	0x0041 00065 (<autogenerated>:1)	PCDATA	$0, $1
	0x0041 00065 (<autogenerated>:1)	CALL	"".val.String(SB)
	0x0046 00070 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x004b 00075 (<autogenerated>:1)	MOVQ	(SP), CX
	0x004f 00079 (<autogenerated>:1)	MOVQ	CX, ""..autotmp_2+16(SP)
	0x0054 00084 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_2+24(SP)
	0x0059 00089 (<autogenerated>:1)	MOVQ	CX, "".~r0+56(SP)
	0x005e 00094 (<autogenerated>:1)	MOVQ	AX, "".~r0+64(SP)
	0x0063 00099 (<autogenerated>:1)	MOVQ	32(SP), BP
	0x0068 00104 (<autogenerated>:1)	ADDQ	$40, SP
	0x006c 00108 (<autogenerated>:1)	RET
	0x006d 00109 (<autogenerated>:1)	PCDATA	$0, $1
	0x006d 00109 (<autogenerated>:1)	CALL	runtime.panicwrap(SB)
	0x0072 00114 (<autogenerated>:1)	UNDEF
	0x0074 00116 (<autogenerated>:1)	NOP
	0x0074 00116 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0074 00116 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0079 00121 (<autogenerated>:1)	JMP	0
	0x007b 00123 (<autogenerated>:1)	LEAQ	48(SP), DI
	0x0080 00128 (<autogenerated>:1)	CMPQ	(BX), DI
	0x0083 00131 (<autogenerated>:1)	JNE	38
	0x0085 00133 (<autogenerated>:1)	MOVQ	SP, (BX)
	0x0088 00136 (<autogenerated>:1)	JMP	38
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 65 48  eH..%....H;a.veH
	0x0010 83 ec 28 48 89 6c 24 20 48 8d 6c 24 20 48 8b 59  ..(H.l$ H.l$ H.Y
	0x0020 20 48 85 db 75 55 0f 57 c0 0f 11 44 24 38 48 8b   H..uU.W...D$8H.
	0x0030 44 24 30 48 85 c0 75 02 eb 33 48 8b 44 24 30 84  D$0H..u..3H.D$0.
	0x0040 00 e8 00 00 00 00 48 8b 44 24 08 48 8b 0c 24 48  ......H.D$.H..$H
	0x0050 89 4c 24 10 48 89 44 24 18 48 89 4c 24 38 48 89  .L$.H.D$.H.L$8H.
	0x0060 44 24 40 48 8b 6c 24 20 48 83 c4 28 c3 e8 00 00  D$@H.l$ H..(....
	0x0070 00 00 0f 0b e8 00 00 00 00 eb 85 48 8d 7c 24 30  ...........H.|$0
	0x0080 48 39 3b 75 a1 48 89 23 eb 9c                    H9;u.H.#..
	rel 5+4 t=16 TLS+0
	rel 66+4 t=8 "".val.String+0
	rel 110+4 t=8 runtime.panicwrap+0
	rel 117+4 t=8 runtime.morestack_noctxt+0
go.string."aaa" SRODATA dupok size=3
	0x0000 61 61 61                                         aaa
go.info."".val.String SDWARFINFO size=63
	0x0000 02 22 22 2e 76 61 6c 2e 53 74 72 69 6e 67 00 00  ."".val.String..
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01  ................
	0x0020 9c 00 00 00 00 01 0e 76 00 00 0e 00 00 00 00 01  .......v........
	0x0030 9c 0e 7e 72 30 00 01 0e 00 00 00 00 01 9c 00     ..~r0..........
	rel 15+8 t=1 "".val.String+0
	rel 23+8 t=1 "".val.String+360
	rel 33+4 t=29 gofile../data/app/go/src/gowhole/exercise/runtime/stack/overflow/main.go+0
	rel 43+4 t=28 go.info."".val+0
	rel 56+4 t=28 go.info.string+0
go.range."".val.String SDWARFRANGE size=0
go.string."str(%v) is Stringer\n" SRODATA dupok size=20
	0x0000 73 74 72 28 25 76 29 20 69 73 20 53 74 72 69 6e  str(%v) is Strin
	0x0010 67 65 72 0a                                      ger.
go.info."".main SDWARFINFO size=82
	0x0000 02 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 09 76 61 6c 75 65 00 1b 00 00 00 00 03 91 e8 7e  .value.........~
	0x0030 09 76 76 00 19 00 00 00 00 03 91 db 7e 13 00 00  .vv.........~...
	0x0040 00 00 09 73 74 72 00 1f 00 00 00 00 03 91 f8 7e  ...str.........~
	0x0050 00 00                                            ..
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+584
	rel 27+4 t=29 gofile../data/app/go/src/gowhole/exercise/runtime/stack/overflow/main.go+0
	rel 40+4 t=28 go.info."".Stringer+0
	rel 53+4 t=28 go.info."".val+0
	rel 62+4 t=28 go.range."".main+0
	rel 72+4 t=28 go.info."".Stringer+0
go.range."".main SDWARFRANGE size=48
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 0+8 t=44 "".main+267
	rel 8+8 t=44 "".main+537
	rel 16+8 t=44 "".main+558
	rel 24+8 t=44 "".main+568
go.info."".init SDWARFINFO size=33
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+104
	rel 27+4 t=29 gofile..<autogenerated>+0
go.range."".init SDWARFRANGE size=0
"".i SNOPTRBSS size=8
"".initdone· SNOPTRBSS size=1
go.info."".Stringer.String SDWARFINFO dupok size=73
	0x0000 02 22 22 2e 53 74 72 69 6e 67 65 72 2e 53 74 72  ."".Stringer.Str
	0x0010 69 6e 67 00 00 00 00 00 00 00 00 00 00 00 00 00  ing.............
	0x0020 00 00 00 00 01 9c 00 00 00 00 01 0e 2e 74 68 69  .............thi
	0x0030 73 00 00 01 00 00 00 00 01 9c 0e 7e 72 31 00 01  s..........~r1..
	0x0040 01 00 00 00 00 02 91 10 00                       .........
	rel 20+8 t=1 "".Stringer.String+0
	rel 28+8 t=1 "".Stringer.String+130
	rel 38+4 t=29 gofile..<autogenerated>+0
	rel 52+4 t=28 go.info."".Stringer+0
	rel 65+4 t=28 go.info.string+0
go.range."".Stringer.String SDWARFRANGE dupok size=0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*func() string- SRODATA dupok size=17
	0x0000 00 00 0e 2a 66 75 6e 63 28 29 20 73 74 72 69 6e  ...*func() strin
	0x0010 67                                               g
type.*func() string SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 bc f4 77 69 00 08 08 36 00 00 00 00 00 00 00 00  ..wi...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func() string-+0
	rel 48+8 t=1 type.func() string+0
type.func() string SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a2 6d cb 1e 02 08 08 33 00 00 00 00 00 00 00 00  .m.....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func() string-+0
	rel 44+4 t=6 type.*func() string+0
	rel 56+8 t=1 type.string+0
type..namedata.*main.Stringer. SRODATA dupok size=17
	0x0000 01 00 0e 2a 6d 61 69 6e 2e 53 74 72 69 6e 67 65  ...*main.Stringe
	0x0010 72                                               r
type.*"".Stringer SRODATA size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 26 d3 e3 8d 00 08 08 36 00 00 00 00 00 00 00 00  &......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.Stringer.+0
	rel 48+8 t=1 type."".Stringer+0
runtime.gcbits.03 SRODATA dupok size=1
	0x0000 03                                               .
type..namedata.String. SRODATA dupok size=9
	0x0000 01 00 06 53 74 72 69 6e 67                       ...String
type."".Stringer SRODATA size=104
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 57 94 f4 eb 07 08 08 14 00 00 00 00 00 00 00 00  W...............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 18 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+128
	rel 32+8 t=1 runtime.gcbits.03+0
	rel 40+4 t=5 type..namedata.*main.Stringer.+0
	rel 44+4 t=5 type.*"".Stringer+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type."".Stringer+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+4 t=5 type..namedata.String.+0
	rel 100+4 t=5 type.func() string+0
go.info."".(*val).String SDWARFINFO dupok size=71
	0x0000 02 22 22 2e 28 2a 76 61 6c 29 2e 53 74 72 69 6e  ."".(*val).Strin
	0x0010 67 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  g...............
	0x0020 00 00 01 9c 00 00 00 00 01 0e 2e 74 68 69 73 00  ...........this.
	0x0030 00 01 00 00 00 00 01 9c 0e 7e 72 30 00 01 01 00  .........~r0....
	0x0040 00 00 00 02 91 08 00                             .......
	rel 18+8 t=1 "".(*val).String+0
	rel 26+8 t=1 "".(*val).String+138
	rel 36+4 t=29 gofile..<autogenerated>+0
	rel 50+4 t=28 go.info.*"".val+0
	rel 63+4 t=28 go.info.string+0
go.range."".(*val).String SDWARFRANGE dupok size=0
type..namedata.*main.val- SRODATA dupok size=12
	0x0000 00 00 09 2a 6d 61 69 6e 2e 76 61 6c              ...*main.val
type..namedata.*func(*main.val) string- SRODATA dupok size=26
	0x0000 00 00 17 2a 66 75 6e 63 28 2a 6d 61 69 6e 2e 76  ...*func(*main.v
	0x0010 61 6c 29 20 73 74 72 69 6e 67                    al) string
type.*func(*"".val) string SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 c9 00 67 a7 00 08 08 36 00 00 00 00 00 00 00 00  ..g....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.val) string-+0
	rel 48+8 t=1 type.func(*"".val) string+0
type.func(*"".val) string SRODATA dupok size=72
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a5 92 7e 19 02 08 08 33 00 00 00 00 00 00 00 00  ..~....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.val) string-+0
	rel 44+4 t=6 type.*func(*"".val) string+0
	rel 56+8 t=1 type.*"".val+0
	rel 64+8 t=1 type.string+0
type.*"".val SRODATA size=88
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 b7 32 c9 86 01 08 08 36 00 00 00 00 00 00 00 00  .2.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 01 00 00 00  ................
	0x0040 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.val-+0
	rel 48+8 t=1 type."".val+0
	rel 56+4 t=5 type..importpath."".+0
	rel 72+4 t=5 type..namedata.String.+0
	rel 76+4 t=24 type.func() string+0
	rel 80+4 t=24 "".(*val).String+0
	rel 84+4 t=24 "".(*val).String+0
runtime.gcbits. SRODATA dupok size=0
type..namedata.*func(main.val) string- SRODATA dupok size=25
	0x0000 00 00 16 2a 66 75 6e 63 28 6d 61 69 6e 2e 76 61  ...*func(main.va
	0x0010 6c 29 20 73 74 72 69 6e 67                       l) string
type.*func("".val) string SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 6a 11 12 9a 00 08 08 36 00 00 00 00 00 00 00 00  j......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(main.val) string-+0
	rel 48+8 t=1 type.func("".val) string+0
type.func("".val) string SRODATA dupok size=72
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 d7 e1 28 a6 02 08 08 33 00 00 00 00 00 00 00 00  ..(....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(main.val) string-+0
	rel 44+4 t=6 type.*func("".val) string+0
	rel 56+8 t=1 type."".val+0
	rel 64+8 t=1 type.string+0
type."".val SRODATA size=112
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 a5 ea 44 91 07 01 01 99 00 00 00 00 00 00 00 00  ..D.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 01 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+16
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*main.val-+0
	rel 44+4 t=5 type.*"".val+0
	rel 56+8 t=1 type."".val+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+4 t=5 type..namedata.String.+0
	rel 100+4 t=24 type.func() string+0
	rel 104+4 t=24 "".(*val).String+0
	rel 108+4 t=24 "".val.String+0
type..namedata.*interface {}- SRODATA dupok size=16
	0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 00 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.03+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=6 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
	0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
	0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 00 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=6 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=19
	0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
	0x0010 20 7b 7d                                          {}
type.[1]interface {} SRODATA dupok size=72
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.03+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 44+4 t=6 type.*[1]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
type.*[1]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 bf 03 a8 35 00 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 48+8 t=1 type.[1]interface {}+0
go.itab."".val,"".Stringer SRODATA dupok size=32
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 a5 ea 44 91 00 00 00 00 00 00 00 00 00 00 00 00  ..D.............
	rel 0+8 t=1 type."".Stringer+0
	rel 8+8 t=1 type."".val+0
	rel 24+8 t=1 "".(*val).String+0
go.itablink."".val,"".Stringer SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 go.itab."".val,"".Stringer+0
type..importpath.fmt. SRODATA dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
gclocals·d8b28f51bb91e05d264803f0f600a200 SRODATA dupok size=11
	0x0000 03 00 00 00 02 00 00 00 00 00 00                 ...........
gclocals·b0938e6026b5c81628342d8ac78cb7a4 SRODATA dupok size=14
	0x0000 03 00 00 00 0a 00 00 00 00 00 69 00 68 00        ..........i.h.
gclocals·f6bd6b3389b872033d462029172c8612 SRODATA dupok size=8
	0x0000 04 00 00 00 00 00 00 00                          ........
gclocals·3e9d367ad8e894cbbe1024f46b88885d SRODATA dupok size=20
	0x0000 04 00 00 00 12 00 00 00 00 00 00 00 08 00 81 06  ................
	0x0010 00 80 06 00                                      ....
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·c55e845a0a62e9baae6c740db5a20866 SRODATA dupok size=10
	0x0000 02 00 00 00 04 00 00 00 03 00                    ..........
gclocals·2589ca35330fc0fce83503f4569854a0 SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 00 00                    ..........
gclocals·e6397a44f8e1b6e77d0f200b4fba5269 SRODATA dupok size=10
	0x0000 02 00 00 00 03 00 00 00 01 00                    ..........
