package nvidia


import (
//	"fmt"
//	"strings"
//	"github.com/golang/glog"
//	"k8s.io/kubernetes/pkg/api"
	gpuTypes "k8s.io/kubernetes/pkg/kubelet/gpu/types"
	gpuUtil "k8s.io/kubernetes/pkg/kubelet/gpu/nvidia/util"
//	"k8s.io/kubernets/pkg/types"
)


const (
	NvidiaName string = "NVIDIA"
	NvidiaDeviceCtl string = "/dev/nvidia/nvidiactl"
	NvidiaDeviceUVM string = "/dev/nvidia/nvidia-uvm"
)


type NvidiaGPU struct {
	gpuInfo gpuTypes.GPUDevice
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

func (nvidiaGPU *NvidiaGPU) Capacity() int, error {
	gpuCount, err := gpuUtil.GetDeviceCount()
	return int(gpuCount), err
}

func (nvidiaGPU *NvidiaGPU) AvailableGPUs() ([]int, error) {
	return nil, nil
}

func (nvidiaGPU *NvidiaGPU) AllocateGPU(number int) ([]gpuTypes.GPUDevice, error) {
	return nil, nil
}

