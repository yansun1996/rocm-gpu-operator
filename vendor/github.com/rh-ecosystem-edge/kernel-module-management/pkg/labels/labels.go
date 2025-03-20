package labels

import (
	"github.com/rh-ecosystem-edge/kernel-module-management/internal/utils"
)

func GetKernelModuleReadyNodeLabel(namespace, moduleName string) string {
	return utils.GetKernelModuleReadyNodeLabel(namespace, moduleName)
}

func GetDevicePluginNodeLabel(namespace, moduleName string) string {
	return utils.GetDevicePluginNodeLabel(namespace, moduleName)
}
