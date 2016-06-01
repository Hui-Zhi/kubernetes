package types

import (
//	"k8s.io/kubernetes/pkg/api"
//	"k8s.io/kubernetes/pkg/types"
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
	InitPlugin() error
	ReleasePlugin() error
	Capacity() (int, error)
	AvailableGPUs() int
	Vendor() string
	AllocateGPU(int) []string
	FreeGPU()
}
