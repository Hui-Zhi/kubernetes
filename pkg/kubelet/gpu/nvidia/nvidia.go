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
	Vendor          string = "NVIDIA"
	NvidiaDeviceCtl string = "/dev/nvidia/nvidiactl"
	NvidiaDeviceUVM string = "/dev/nvidia/nvidia-uvm"
)

type NvidiaGPU struct {
	gpuInfo []gpuTypes.GPUDevice
}

func ProbePlugin() (gpuTypes.GPUPlugin, error) {
	var nvidiaGPU NvidiaGPU

	err := nvidiaGPU.InitPlugin()

	if err != nil {
		return nil, err
	}

	return &nvidiaGPU, nil
}

func (nvidiaGPU *NvidiaGPU) Vendor() string {
	return Vendor
}

func (nvidiaGPU *NvidiaGPU) InitPlugin() error {
	err := gpuUtil.NVMLInit()

	if err != nil {
		return nil
	}

	allPaths, err := nvidiaGPU.discovery()

	if err != nil {
		return nil
	}

	for _, path := range allPaths {
		nvidiaGPU.gpuInfo = append(nvidiaGPU.gpuInfo, gpuTypes.GPUDevice{Path: path, Status: gpuTypes.GPUFree})
	}

	return nil
}

func (nvidiaGPU *NvidiaGPU) ReleasePlugin() error {
	err := gpuUtil.NVMLShutdown()

	if err != nil {
		return nil
	}

	nvidiaGPU.gpuInfo = nil

	return gpuUtil.NVMLShutdown()
}

func (nvidiaGPU *NvidiaGPU) discovery() ([]string, error) {
	gpuCount, err := gpuUtil.GetDeviceCount()

	if err != nil {
		return nil, err
	}

	if gpuCount <= 0 {
		return nil, nil
	}

	var allPaths []string
	var i uint
	for i = 0; i < gpuCount; i++ {
		path, err := gpuUtil.GetDevicePath(i)
		if err != nil {
			return nil, err
		}
		allPaths = append(allPaths, path)
	}

	return allPaths, nil
}

func (nvidiaGPU *NvidiaGPU) Capacity() (int, error) {
	gpuCount, err := gpuUtil.GetDeviceCount()
	return int(gpuCount), err
}

func (nvidiaGPU *NvidiaGPU) AvailableGPUs() int {
	freeGPUs := 0
	for _, gpuItem := range nvidiaGPU.gpuInfo {
		if gpuItem.Status == gpuTypes.GPUFree {
			freeGPUs += 1
		}
	}

	return freeGPUs
}

func (nvidiaGPU *NvidiaGPU) AllocateGPU(number int) (allocPaths []string) {
	avaliableGPUs := nvidiaGPU.AvailableGPUs()
	if avaliableGPUs < number {
		return
	}

	for _, gpuItem := range nvidiaGPU.gpuInfo {
		if gpuItem.Status == gpuTypes.GPUFree {
			allocPaths = append(allocPaths, gpuItem.Path)
			gpuItem.Status = gpuTypes.GPUInUsing
		}
	}

	return
}

func (nvidiaGPU *NvidiaGPU) FreeGPU() {
}
