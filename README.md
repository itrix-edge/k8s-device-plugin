## k8s-k8s-device-plugin for arm64 architecture

Allows k8s-k8s-device-plugin on arm64 architecture to allocate  NVIDIA GPUs resources.



#### Allocate gpus by k8s yaml
```
    resources:
      limits:
        nvidia.com/gpu: 1
```
