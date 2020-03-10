// Copyright (c) 2017, NVIDIA CORPORATION. All rights reserved.

package main

//#cgo CFLAGS: -I.
//#cgo LDFLAGS: -L/go/src/semaphore-device-plugin -lcuda
//#include <cuda.h>
import "C"
import (
	//"flag"
	"github.com/google/uuid"
	"log"
	"os"
	"syscall"

	"github.com/fsnotify/fsnotify"
	pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1beta1"
)

var (
	semaphore int
	C_DeviceCount C.int
)
/*
func parseFlags() {
	//devicecount := C.DeviceGetCount()
	flag.IntVar(&semaphore, "semaphore", devicecount, "Semaphore count")
	flag.Parse()
}
*/

func main() {

	//parseFlag()
        C_DeviceCount = C.DeviceGetCount()
        semaphore = int(C_DeviceCount)

	log.Printf("Semaphore count is: %d", semaphore)

	log.Println("Starting FS watcher.")
	watcher, err := newFSWatcher(pluginapi.DevicePluginPath)
	if err != nil {
		log.Println("Failed to created FS watcher.")
		os.Exit(1)
	}
	defer watcher.Close()

	log.Println("Starting OS watcher.")
	sigs := newOSWatcher(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	restart := true
	var devicePlugin *SemaphoreDevicePlugin

L:
	for {
		if restart {
			if devicePlugin != nil {
				devicePlugin.Stop()
			}

			devicePlugin = NewSemaphoreDevicePlugin(semaphore)
			if err := devicePlugin.Serve(); err != nil {
				log.Println("Could not contact Kubelet, retrying. Did you enable the device plugin feature gate?")
			} else {
				restart = false
			}
		}

		select {
		case event := <-watcher.Events:
			if event.Name == pluginapi.KubeletSocket && event.Op&fsnotify.Create == fsnotify.Create {
				log.Printf("inotify: %s created, restarting.", pluginapi.KubeletSocket)
				restart = true
			}

		case err := <-watcher.Errors:
			log.Printf("inotify: %s", err)

		case s := <-sigs:
			switch s {
			case syscall.SIGHUP:
				log.Println("Received SIGHUP, restarting.")
				restart = true
			default:
				log.Printf("Received signal \"%v\", shutting down.", s)
				devicePlugin.Stop()
				break L
			}
		}
	}
}

func getDevices(n int) []*pluginapi.Device {

	var devs []*pluginapi.Device
	for i := 0; i < n; i++ {
		devs = append(devs, &pluginapi.Device{
			ID:     uuid.New().String(),
			Health: pluginapi.Healthy,
		})
	}
	return devs
}

