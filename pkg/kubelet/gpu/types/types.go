package types

import (
//	"k8s.io/kubernetes/pkg/api"
//	"k8s.io/kubernetes/pkg/types"
)

type GPUDevice struct {
	Path string
}

type GPUPlugin interface {
	InitPlugin() error
	ReleasePlugin() error
	Capacity() (int, error)
	AvailableGPUs() ([]int, error)
	Name() string
	AllocateGPU(int) ([]GPUDevice, error)
}


