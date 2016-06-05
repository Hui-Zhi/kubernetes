package gpu

import (
	"k8s.io/kubernetes/pkg/kubelet/gpu/nvidia"
	gputypes "k8s.io/kubernetes/pkg/kubelet/gpu/types"
)

// Initialize all the GPU plugin, current we only support
// NVIDIA GPU.
// TODO: we could support more GPUs from different vendor
// like Intel, AMD.
func ProbeGPUPlugins() []gputypes.GPUPlugin {
	// Initialize NVIDIA GPU, and all kinds of GPU should
	// be implemented as a plugin and initialize here.
	nvidiaPlugin, err := nvidia.ProbePlugin()

	if err != nil {
		// If the GPU initialization failed, we should ignore
		// this plugin, at present we return nil due to we
		// only support NVIDIA GPU plugin.
		return nil
	}

	allPlugins := []gputypes.GPUPlugin{nvidiaPlugin}

	return allPlugins
}
