// +build darwin freebsd netbsd openbsd
// +build !cgo

package loadavg

import (
	"fmt"
	"syscall"
	"unsafe"
)

func get() (*Loadavg, error) {
	ret, err := syscall.Sysctl("vm.loadavg")
	if err != nil {
		return nil, fmt.Errorf("failed in sysctl vm.loadavg: %s", err)
	}
	return collectLoadavgStats([]byte(ret))
}

// loadavg in sysctl.h
type loadStruct struct {
	Ldavg  [3]uint32
	Fscale uint64
}

func collectLoadavgStats(out []byte) (*Loadavg, error) {
	load := *(*loadStruct)(unsafe.Pointer(&out[0]))
	return &Loadavg{
		Loadavg1:  float64(load.Ldavg[0]) / float64(load.Fscale),
		Loadavg5:  float64(load.Ldavg[1]) / float64(load.Fscale),
		Loadavg15: float64(load.Ldavg[2]) / float64(load.Fscale),
	}, nil
}