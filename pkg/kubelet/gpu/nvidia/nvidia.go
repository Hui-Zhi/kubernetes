package nvidia

import (
	gpuutil "k8s.io/kubernetes/pkg/kubelet/gpu/nvidia/util"
	gputypes "k8s.io/kubernetes/pkg/kubelet/gpu/types"
)

const (
	Vendor string = "NVIDIA"
	// All NVIDIA GPUs cards should be mounted with nvidiactl and nvidia-uvm
	NvidiaDeviceCtl string = "/dev/nvidiactl"
	NvidiaDeviceUVM string = "/dev/nvidia-uvm"
)

type NvidiaGPU struct {
	gpuInfo []gputypes.GPUDevice
}

func ProbePlugin() (gputypes.GPUPlugin, error) {
	var nvidiaGPU NvidiaGPU

	err := nvidiaGPU.InitPlugin()

	if err != nil {
		return nil, err
	}

	return &nvidiaGPU, nil
}

// Get the vendor information.
func (nvidiaGPU *NvidiaGPU) Vendor() string {
	return Vendor
}

func (nvidiaGPU *NvidiaGPU) InitPlugin() error {
	err := gpuutil.NVMLInit()

	if err != nil {
		return err
	}

	allPaths, err := nvidiaGPU.discovery()

	if err != nil {
		return err
	}

	// Initialize the Nvidia device information.
	for _, path := range allPaths {
		nvidiaGPU.gpuInfo = append(nvidiaGPU.gpuInfo, gputypes.GPUDevice{Path: path, Status: gputypes.GPUFree, ContainerID: ""})
	}

	return nil
}

func (nvidiaGPU *NvidiaGPU) ReleasePlugin() error {
	err := gpuutil.NVMLShutdown()

	if err != nil {
		return err
	}

	nvidiaGPU.gpuInfo = nil

	return gpuutil.NVMLShutdown()
}

// Get all the NVIDIA GPU cards' path from /dev/
func (nvidiaGPU *NvidiaGPU) discovery() ([]string, error) {
	gpuCount, err := gpuutil.GetDeviceCount()

	if err != nil {
		return nil, err
	}

	if gpuCount <= 0 {
		return nil, nil
	}

	var allPaths []string
	var i uint
	for i = 0; i < gpuCount; i++ {
		path, err := gpuutil.GetDevicePath(i)
		if err != nil {
			return nil, err
		}
		allPaths = append(allPaths, path)
	}

	return allPaths, nil
}

func (nvidiaGPU *NvidiaGPU) Capacity() (int, error) {
	gpuCount, err := gpuutil.GetDeviceCount()
	return int(gpuCount), err
}

func (nvidiaGPU *NvidiaGPU) AvailableGPUs() int {
	freeGPUs := 0
	for _, gpuItem := range nvidiaGPU.gpuInfo {
		if gpuItem.Status == gputypes.GPUFree {
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
		if gpuItem.Status == gputypes.GPUFree {
			allocPaths = append(allocPaths, gpuItem.Path)
			gpuItem.Status = gputypes.GPUInUsing
		}
	}

	return
}

func (nvidiaGPU *NvidiaGPU) UpdateContainerID(containerID string, Paths []string) {
	for _, path := range Paths {
		for _, gpuItem := range nvidiaGPU.gpuInfo {
			if path == gpuItem.Path {
				gpuItem.ContainerID = containerID
			}
		}
	}
}

func (nvidiaGPU *NvidiaGPU) FreeGPUByContainer(containerID string) {
	for _, gpuItem := range nvidiaGPU.gpuInfo {
		if gpuItem.ContainerID == containerID {
			gpuItem.Status = gputypes.GPUFree
		}
	}
}

func (nvidiaGPU *NvidiaGPU) FreeGPUByPaths(paths []string) {
	for _, path := range paths {
		for _, gpuItem := range nvidiaGPU.gpuInfo {
			if gpuItem.Path == path {
				gpuItem.Status = gputypes.GPUFree
				continue
			}
		}
	}
}
