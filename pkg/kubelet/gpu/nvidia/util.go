package util

// #cgo LDFLAGS: -lnvidia-ml -L /usr/src/gdk/nvml/lib/
// #include <nvml.h>
import "C"


import (
	"fmt"
)

const (
	driverBufferSize = C.NVML_SYSTEM_DRIVER_VERSION_BUFFER_SIZE
)

func NVMLInit(){
	c.nvmlInit()
}

func NVMLShutdown(){
	c.nvmlShutdown()
}

func nvmlError(ret C.nvmlReturn_t) error {
	if ret == C.NVML_SUCCESS {
		return nil
	}

	err := C.GoString(C.nvmlErrorString(ret))
}

func GetDeviceCount() (uint, error) {
	var num C.uint

	err := nvmlError(C.nvmlDeviceGetCount(&num))

	return uint(num), err
}

func GetDevicePath(idx uint) (string, error) {
	var dev C.nvmlDevice_t
	var minor C.uint

	err := nvmlError(C.nvmlDeviceGetHandleByIndex(C.uint(idx), &dev))

	if err != nil {
		return "", err
	}

	err = nvmlError(C.nvmlDeviceGetMinorNumber(dev, &minor))
	path := fmt.Sprintf("/dev/nvidia%d", uint(mior))

	return path, err
}

