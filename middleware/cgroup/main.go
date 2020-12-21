package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/containerd/cgroups"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	"github.com/qiniu/log"

	"gowhole/middleware/cgroup/mock"
)

func main() {
	Create()

	// Stat()
}

func Create() {
	mock, err := mock.NewMock()
	if err != nil {
		log.Fatal(err)
	}
	defer mock.Delete()
	//cgroups.V1 mock.Hierarchy
	control, err := cgroups.New(mock.Hierarchy, cgroups.StaticPath("test"), &specs.LinuxResources{})
	if err != nil {
		log.Error(err)
		return
	}
	if control == nil {
		log.Error("control is nil")
		return
	}

	for _, s := range cgroups.Subsystems() {
		fmt.Println(filepath.Join(mock.Root, string(s), "test"))
		if _, err := os.Stat(filepath.Join(mock.Root, string(s), "test")); err != nil {
			if os.IsNotExist(err) {
				log.Errorf("group %s was not created", s)
				return
			}
			log.Errorf("group %s was not created correctly %s", s, err)
			return
		}
	}
}

func Stat() {
	mock, err := mock.NewMock()
	if err != nil {
		log.Fatal(err)
	}
	defer mock.Delete()
	control, err := cgroups.New(mock.Hierarchy, cgroups.StaticPath("test"), &specs.LinuxResources{})
	if err != nil {
		log.Error(err)
		return
	}

	s, err := control.Stat(cgroups.IgnoreNotExist)
	if err != nil {
		log.Error(err)
		return
	}
	if s == nil {
		log.Error("stat result is nil")
		return
	}
}
