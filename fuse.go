package main

import (
	"log"
	"os"
	pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1alpha1"
	"fmt"
)

func check(err error) {
	if err != nil {
		log.Panicln("Fatal:", err)
	}
}

func getDevices() []*pluginapi.Device {
	hostname, _ := os.Hostname()
	devs := []*pluginapi.Device{
		&pluginapi.Device{
			ID:     fmt.Sprintf("fuse-%s", hostname),
			Health: pluginapi.Healthy,
		},
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

