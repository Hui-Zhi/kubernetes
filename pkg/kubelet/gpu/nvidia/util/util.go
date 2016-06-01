package util

// #cgo LDFLAGS: -lnvidia-ml -L /usr/src/gdk/nvml/lib/
// #cgo CPPFLAGS: -I /usr/include/nvidia/gdk/
// #include <nvml.h>
import "C"

import (
	"fmt"
)

const (
	driverBufferSize = C.NVML_SYSTEM_DRIVER_VERSION_BUFFER_SIZE
)

func NVMLInit() error {
	ret := C.nvmlInit()

	if ret == C.NVML_ERROR_LIBRARY_NOT_FOUND {
		return fmt.Errorf("Could not find NVML library!")
	}

	return nvmlError(ret)
}

func NVMLShutdown() error {
	return nvmlError(C.nvmlShutdown())
}

func nvmlError(ret C.nvmlReturn_t) error {
	if ret == C.NVML_SUCCESS {
		return nil
	}

	err := C.GoString(C.nvmlErrorString(ret))

	return fmt.Errorf("NVML error: %v", err)
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
	path := fmt.Sprintf("/dev/nvidia%d", uint(minor))

	return path, err
}
