package main

import (
	"fmt"
	"os"

	pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1beta1"
)

// func check(err error) {
// 	if err != nil {
// 		log.Panicln("Fatal:", err)
// 	}
// }

func getDevices(number int) []*pluginapi.Device {
	hostname, _ := os.Hostname()
	devs := []*pluginapi.Device{}
	for i := 0; i < number; i++ {
		devs = append(devs, &pluginapi.Device{
			ID:     fmt.Sprintf("fuse-%s-%d", hostname, i),
			Health: pluginapi.Healthy,
		})
	}
	return devs
}

func deviceExists(devs []*pluginapi.Device, id string) bool {
	for _, d := range devs {
		if d.ID == id {
			return true
		}
	}
	return false
}
