package gpu

import (
	//	"github.com/golang/glog"
	//	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/kubelet/gpu/nvidia"
	gpuTypes "k8s.io/kubernetes/pkg/kubelet/gpu/types"
)

func ProbeGPUPlugins() []gpuTypes.GPUPlugin {
	nvidiaPlugin, err := nvidia.ProbePlugin()

	if err != nil {
		return nil
	}

	allPlugins := []gpuTypes.GPUPlugin{nvidiaPlugin}

	return allPlugins
}
