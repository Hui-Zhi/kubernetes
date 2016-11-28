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

package types

const (
	// GPU has been assigned to container.
	GPUInUse = "InUse"
	// No container is using this GPU.
	GPUFree = "Free"
	// Don't know the status of the GPU.
	GPUUnknown = "Unknown"
)

// Basic device struct for all the GPU type.
// Like the path, status, also we need to
// associate it with container, once a container
// be killed/stopped, we need to free it.
type GPUDevice struct {
	ContainerName string
	Path          string
	Status        string
}

// General GPU plugin interface.
// The GPU is dedicated.
type GPUPlugin interface {
	InitPlugin() error
	ReleasePlugin() error
	// Number of GPU cards. different vendor only
	// need to care about their own cards.
	Capacity() int
	AvailableGPUs() int
	Vendor() string
	AllocateGPU(string, int) []string
	Valid(string) bool
	UpdateGPUByContainerName(string, []string)
	FreeGPUByContainerName(string)
	FreeGPUByPaths([]string)
}
