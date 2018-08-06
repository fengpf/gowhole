package tests

import (
	"fmt"
	"testing"
)

const (
	// User user handler
	User = iota
	// UserMobile mobile handler
	UserMobile
	// UserWeb user web handler
	UserWeb
	// Internal internal handler
	Internal
	// Internet internet handler
	Internet
	// Guest guest handler
	Guest
	// GuestMobile guest mobile handler
	GuestMobile
	// GuestWeb guest web handler
	GuestWeb

	// CtxProto proto context key.
	CtxProto = "proto"
	// CtxTTL ttl context key.
	CtxTTL = "ttl"
	// CtxCSRF csrf context key.
	CtxCSRF = "csrf"
	// CtxVerify verify context key.
	CtxVerify = "verify"
	// CtxMID mid context key.
	CtxMID = "mid"
	// CtxIdentify identify type key.
	CtxIdentify = "identify"
	// CtxSupervisor supervisor context key.
	CtxSupervisor = "supervisor"
	// CtxAntispam antispam for api
	CtxAntispam = "antispam"
	// CtxDegrade degrade context key
	CtxDegrade = "degrade"
)

func print(args ...int) {
	fmt.Printf("t  %p\n", args)
}

func Test_enum(t *testing.T) {
	a := []int{1, 2, 3}
	b := a[1:]

	print(a...)
	print(b...)

	fmt.Printf("main %p\n", a)
	fmt.Printf("main %p\n", b)

	fmt.Printf("User %v\n", User)
	fmt.Printf("UserMobile %v\n", UserMobile)
	fmt.Printf("CtxDegrade %+v\n", CtxDegrade)

}
