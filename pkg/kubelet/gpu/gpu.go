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
