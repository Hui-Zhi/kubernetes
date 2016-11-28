/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nvidia

import (
	"sync"

	gpuutil "k8s.io/kubernetes/pkg/kubelet/gpu/nvidia/util"
	gputypes "k8s.io/kubernetes/pkg/kubelet/gpu/types"
)

// TODO: If use NVML in the future, the implementation could be more complex,
// and more powerful!

type NvidiaGPU struct {
	gpuInfo  []gputypes.GPUDevice
	gpuMutex sync.Mutex
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
	return gpuutil.Vendor
}

func (nvidiaGPU *NvidiaGPU) InitPlugin() error {
	nvidiaGPU.gpuMutex.Lock()
	defer nvidiaGPU.gpuMutex.Unlock()
	if err := gpuutil.Init(); err != nil {
		return err
	}

	allPaths := gpuutil.GetGPUPaths()

	// Initialize the Nvidia device information.
	for _, path := range allPaths {
		nvidiaGPU.gpuInfo = append(nvidiaGPU.gpuInfo, gputypes.GPUDevice{Path: path, Status: gputypes.GPUFree, ContainerName: ""})
	}

	return nil
}

// Release the Plugin, could be useful once we use NVML.
func (nvidiaGPU *NvidiaGPU) ReleasePlugin() error {
	nvidiaGPU.gpuMutex.Lock()
	defer nvidiaGPU.gpuMutex.Unlock()

	nvidiaGPU.gpuInfo = nil

	return nil
}

func (nvidiaGPU *NvidiaGPU) Capacity() int {
	nvidiaGPU.gpuMutex.Lock()
	defer nvidiaGPU.gpuMutex.Unlock()

	return len(nvidiaGPU.gpuInfo)
}

func (nvidiaGPU *NvidiaGPU) AvailableGPUs() int {
	nvidiaGPU.gpuMutex.Lock()
	defer nvidiaGPU.gpuMutex.Unlock()

	freeGPUs := 0
	for _, gpuItem := range nvidiaGPU.gpuInfo {
		if gpuItem.Status == gputypes.GPUFree {
			freeGPUs += 1
		}
	}

	return freeGPUs
}

func (nvidiaGPU *NvidiaGPU) AllocateGPU(containerName string, number int) (allocPaths []string) {
	availableGPUs := nvidiaGPU.AvailableGPUs()

	nvidiaGPU.gpuMutex.Lock()
	defer nvidiaGPU.gpuMutex.Unlock()

	if availableGPUs < number {
		return
	}

	allocatedGPU := 0

	for _, gpuItem := range nvidiaGPU.gpuInfo {
		if allocatedGPU >= number {
			return
		}

		if gpuItem.Status == gputypes.GPUFree {
			allocPaths = append(allocPaths, gpuItem.Path)
			gpuItem.Status = gputypes.GPUInUse
			gpuItem.ContainerName = containerName

			allocatedGPU++
		}
	}

	return
}

func (nvidiaGPU *NvidiaGPU) Valid(path string) bool {
	return gpuutil.Valid(path)
}

func (nvidiaGPU *NvidiaGPU) UpdateGPUByContainerName(name string, Paths []string) {
	nvidiaGPU.gpuMutex.Lock()
	defer nvidiaGPU.gpuMutex.Unlock()

	for _, path := range Paths {
		for _, gpuItem := range nvidiaGPU.gpuInfo {
			if path == gpuItem.Path {
				gpuItem.ContainerName = name
			}
		}
	}
}

func (nvidiaGPU *NvidiaGPU) FreeGPUByContainerName(name string) {
	nvidiaGPU.gpuMutex.Lock()
	defer nvidiaGPU.gpuMutex.Unlock()

	for _, gpuItem := range nvidiaGPU.gpuInfo {
		if gpuItem.ContainerName == name {
			gpuItem.Status = gputypes.GPUFree
		}
	}
}

func (nvidiaGPU *NvidiaGPU) FreeGPUByPaths(paths []string) {
	nvidiaGPU.gpuMutex.Lock()
	defer nvidiaGPU.gpuMutex.Unlock()

	for _, path := range paths {
		for _, gpuItem := range nvidiaGPU.gpuInfo {
			if gpuItem.Path == path {
				gpuItem.Status = gputypes.GPUFree
			}
		}
	}
}
