package integration_test

import (
	"context"
	"gowhole/middleware/grpc/gw/internal/clients/echo"
	"testing"

)

func TestEchoClient(t *testing.T) {
	if testing.Short() {
		t.Skip()
		return
	}

	cfg := echo.NewConfiguration()
	cfg.BasePath = "http://localhost:8088"

	cl := echo.NewAPIClient(cfg)
	resp, _, err := cl.EchoServiceApi.EchoServiceEcho(context.Background(), "foo")
	if err != nil {
		t.Errorf(`cl.EchoServiceApi.Echo("foo") failed with %v; want success`, err)
	}
	if got, want := resp.Id, "foo"; got != want {
		t.Errorf("resp.Id = %q; want %q", got, want)
	}
}

func TestEchoBodyClient(t *testing.T) {
	if testing.Short() {
		t.Skip()
		return
	}

	cfg := echo.NewConfiguration()
	cfg.BasePath = "http://localhost:8088"

	cl := echo.NewAPIClient(cfg)
	req := echo.ExamplepbSimpleMessage{Id: "foo"}
	resp, _, err := cl.EchoServiceApi.EchoServiceEchoBody(context.Background(), req)
	if err != nil {
		t.Errorf("cl.EchoBody(%#v) failed with %v; want success", req, err)
	}
	if got, want := resp.Id, "foo"; got != want {
		t.Errorf("resp.Id = %q; want %q", got, want)
	}
}
