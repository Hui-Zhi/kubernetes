package nvidia


package (
	"fmt"
	"strings"
	"github.com/golang/glog"
	"k8s.io/kubernetes/pkg/api"
	gpuTypes "k8s.io/kubernetes/pkg/kubelet/gpu/types"
	gpuUtil "k8s.io/kubernetes/pkg/kubelet/gpu/util"
	"k8s.io/kubernets/pkg/types"
)


const (
	NvidiaName string = "NVIDIA"
	NvidiaDeviceCtl string = "/dev/nvidia/nvidiactl"
	NvidiaDeviceUVM string = "/dev/nvidia/nvidia-uvm"
)


type NvidiaGPU struct {
	gpuInfo GPUTypes.GPUInfo
}

func (nvidiaGPU *NvidiaGPU) Name() string {
	return NvidiaName
}

func (nvidiaGPU *NvidiaGPU) InitPlugin() error {
}

func (nvidiaGPU *NvidiaGPU) GPUInitPlugin() error {
}

func (nvidiaGPU *NvidiaGPU) ReleasePlugin() error {
}

func (nvidiaGPU *NvidiaGPU) InitGPUEnv() error {
}

func (nvidiaGPU *NvidiaGPU) GetAvailableGPUs() ([]int, error) {
}

	
