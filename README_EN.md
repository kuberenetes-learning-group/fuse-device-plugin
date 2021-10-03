English | [中文](README.md)

# fuse device plugin for Kubernetes

This repository contains implementation of the [Kubernetes device plugin](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/resource-management/device-plugin.md) alpha feature from version 1.8.

## Background
When pod need access /dev/fuse, e.g to mount sshfs or ceph s3fs, containter need to be running in privileged mode to access host's /dev/fuse device, this is not safe.

What make it worse is when the pod need access nvidia GPU devices on the host node. 

For example totally 4 GPU devices available on the host node, but only one of them is requested and allocated for the pod, in privileged mode, all of the 4 GPUs are accessible by the application running in the pod container, application like tensorflow will use all the visible GPUs by default, if the other GPUs already allocated for other pod, will cause errors
  
With this device plugin, pod dont need to run in privileged mode to access /dev/fuse by injecting the fuse device to pod directly, avoiding the potential confilct for GPU.

## Usage
Please make sure that the Kubelet has been started with the `--feature-gates=DevicePlugins=true`
before running the device plugin.

#### Deploy as Daemon Set:

* kubernete version < 1.16

```bash
kubectl create -f fuse-device-plugin.yml
```

* kubernete version > 1.16

```
kubectl create -f fuse-device-plugin-k8s-1.16.yml
```

#### Deploy

Refer to [fuse-test.yml](fuse-test.yml)

```yaml
spec: 
  containers:
  - ...
    resources:
      limits:
        github.com/fuse: 1
```


# Thanks

![Goland](https://blog.jetbrains.com/wp-content/uploads/2019/01/goland_icon.svg)
