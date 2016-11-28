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

package util

import (
	"os"
	"path/filepath"
	"regexp"
)

const (
	Vendor string = "NVIDIA"
	// All NVIDIA GPUs cards should be mounted with nvidiactl and nvidia-uvm
	NvidiaDeviceCtl string = "/dev/nvidiactl"
	NvidiaDeviceUVM string = "/dev/nvidia-uvm"
)

var (
	gpuPaths []string
)

// Get all the paths of NVIDIA GPU card from /dev/
func discovery() error {
	var err error
	if gpuPaths == nil {
		err = filepath.Walk("/dev", func(path string, f os.FileInfo, err error) error {
			reg := regexp.MustCompile(`^nvidia[0-9]*$`)
			gpupath := reg.FindAllString(f.Name(), -1)
			if gpupath != nil && gpupath[0] != "" {
				gpuPaths = append(gpuPaths, "/dev/"+gpupath[0])
			}

			return nil
		})
	}

	return err
}

func GetGPUPaths() []string {
	return gpuPaths
}

func Valid(path string) bool {
	reg := regexp.MustCompile(`^/dev/nvidia[0-9]*$`)
	check := reg.FindAllString(path, -1)

	return check != nil && check[0] != ""
}

func Init() error {
	if _, err := os.Stat(NvidiaDeviceCtl); err != nil {
		return err
	}
	if _, err := os.Stat(NvidiaDeviceUVM); err != nil {
		return err
	}

	err := discovery()
	return err
}
