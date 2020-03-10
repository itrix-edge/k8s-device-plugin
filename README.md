## k8s-k8s-device-plugin for arm64 architecture

Allows k8s-k8s-device-plugin on arm64 architecture to allocate  NVIDIA GPUs resources.

Please make sure the system has cuda for arm64, k8s-device-plugin use cuda_runtime_API to get gpu device counts.

#### Build docker image
```
docker build -t kevin7674/k8s-device-plugin:1.11 .
```
#### deploy k8s-device-plugin
```
kubectl create -f k8s-device-plugin/k8s/k8s-device-plugin.yml
```
#### Allocate gpus when create a k8s yaml
```
    resources:
      limits:
        nvidia.com/gpu: 1
```
