package main

import "testing"

func TestSign(t *testing.T) {
	token := "token1234567890"
	timestamp := "1458864389"
	nonce := "2066297436"
	wantSignature := "f9c725922f6844701ba71e98031978e40023c09f"
	haveSignature := sign(token, timestamp, nonce)
	if haveSignature != wantSignature {
		t.Errorf("test Sign failed,\nhave signature: %s\nwant signature: %s\n", haveSignature, wantSignature)
		return
	}
}
