package nvidia

import (
	//	"fmt"
	//	"strings"
	//	"github.com/golang/glog"
	//	"k8s.io/kubernetes/pkg/api"
	gpuUtil "k8s.io/kubernetes/pkg/kubelet/gpu/nvidia/util"
	gpuTypes "k8s.io/kubernetes/pkg/kubelet/gpu/types"
	//	"k8s.io/kubernets/pkg/types"
)

const (
	NvidiaName      string = "NVIDIA"
	NvidiaDeviceCtl string = "/dev/nvidia/nvidiactl"
	NvidiaDeviceUVM string = "/dev/nvidia/nvidia-uvm"
)

type NvidiaGPU struct {
	gpuInfo []gpuTypes.GPUDevice
}

func (nvidiaGPU *NvidiaGPU) ProbePlugin() (gpuTypes.GPUPlugin, error) {
	gpuUtil.NVMLInit()
	gpuDevices, err := nvidiaGPU.Discovery()

	if err != nil {
		return nil, err
	}

	return &NvidiaGPU{gpuInfo: gpuDevices}, nil
}

func (nvidiaGPU *NvidiaGPU) Name() string {
	return NvidiaName
}

func (nvidiaGPU *NvidiaGPU) InitPlugin() error {
	return gpuUtil.NVMLInit()
}

func (nvidiaGPU *NvidiaGPU) ReleasePlugin() error {
	return gpuUtil.NVMLShutdown()
}

func (nvidiaGPU *NvidiaGPU) Discovery() ([]gpuTypes.GPUDevice, error) {
	gpuCount, err := gpuUtil.GetDeviceCount()

	if err != nil {
		return nil, err
	}

	if gpuCount <= 0 {
		return nil, nil
	}

	gpuDevices := []gpuTypes.GPUDevice{}

	var i uint
	for i = 0; i < gpuCount; i++ {
		path, err := gpuUtil.GetDevicePath(i)
		if err != nil {
			return nil, err
		}
		append(gpuDevices, gpuTypes.GPUDevice{Path: path, Status: gpuTypes.GPUUnknow})
	}

	return gpuDevices, nil
}

func (nvidiaGPU *NvidiaGPU) Capacity() (int, error) {
	gpuCount, err := gpuUtil.GetDeviceCount()
	return int(gpuCount), err
}

func (nvidiaGPU *NvidiaGPU) AvailableGPUs() (int, error) {
	return 0, nil
}

func (nvidiaGPU *NvidiaGPU) AllocateGPU(number int) error {
	return nil
}

func (nvidiaGPU *NvidiaGPU) FreeGPU() {
}
