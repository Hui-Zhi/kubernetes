package types

import (
	//	"k8s.io/kubernetes/pkg/api"
	//	"k8s.io/kubernetes/pkg/types"
	gpuTypes "k8s.io/kubernetes/pkg/kubelet/gpu/types"
)

const (
	GPUInUsing = "Using"
	GPUFree    = "Free"
	GPUUnknow  = "Unknow"
)

type GPUDevice struct {
	Path   string
	Status string
}

type GPUPlugin interface {
	ProbePlugin() (types.GPUPlugin, error)
	InitPlugin() error
	ReleasePlugin() error
	Discovery() ([]GPUDevice, error)
	Capacity() (int, error)
	AvailableGPUs() (int, error)
	Name() string
	AllocateGPU(int) error
	FreeGPU()
}
