## k8s-device-plugin for arm64 architecture

Allows k8s-k8s-device-plugin to allocate  NVIDIA GPUs resources on arm64 architecture.

Please make sure the Linux system has Cuda for arm64 version, k8s-device-plugin use cuda_runtime_API to get gpu device counts.

#### Build docker image
```
docker build -t kevin7674/k8s-device-plugin:1.11 .
```
#### deploy k8s-device-plugin
```
kubectl create -f k8s-device-plugin/k8s/k8s-device-plugin.yml
```
#### Allocate gpus when create a k8s pod yaml
```
    resources:
      limits:
        nvidia.com/gpu: 1
```
#### Device query

https://github.com/NVIDIA/nvidia-docker/wiki/NVIDIA-Container-Runtime-on-Jetson






