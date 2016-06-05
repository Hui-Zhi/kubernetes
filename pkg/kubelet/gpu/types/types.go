package types

import ()

const (
	// GPU has been assigned to container.
	GPUInUsing = "Using"
	// No container is using this GPU.
	GPUFree = "Free"
	// Don't know the status of the GPU.
	GPUUnknow = "Unknow"
)

// Basic device struct for all the GPU type.
// Like the path, status, also we need to
// associate it with container, once a container
// be killed/stopped, we need to free it.
type GPUDevice struct {
	ContainerID string
	Path        string
	Status      string
}

// General GPU plugin interface.
// The GPU is dedicated.
type GPUPlugin interface {
	InitPlugin() error
	ReleasePlugin() error
	// Number of GPU cards. different vendor only
	// need to care about their own cards.
	Capacity() (int, error)
	AvailableGPUs() int
	Vendor() string
	AllocateGPU(int) []string
	UpdateContainerID(string, []string)
	FreeGPUByContainer(string)
	FreeGPUByPaths([]string)
}
