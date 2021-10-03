[English](README_EN.md) | 中文

# fuse device plugin

> Inspired by @JasonChenY's [fuse-device-plugin](https://github.com/JasonChenY/fuse-device-plugin)

## 环境要求

[Kubernetes](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/resource-management/device-plugin.md) version >= 1.8.

## 背景

使用 `sshfs` 或者 `s3fs` 等时， 需要在容器中使用 `/dev/fuse` 的话需要使用特权模式，这会带来许多的问题，比如GPU数量无法屏蔽，容器内可以看到宿主机上所有的GPU卡数。基于此，我们可以仿照 `nvidia-device-plugin` 的方式实现 `fuse-device-plugin`,通过注入的方式来使用 `/dev/fuse`

## 使用要求

使用前请确保 `--feature-gates=DevicePlugins=true` 已开启.

```bash
kubelet -h | grep "DevicePlugins"
```

## 部署:

* kubernete version < 1.16

```bash
kubectl create -f fuse-device-plugin.yml
```

* kubernete version > 1.16

```
kubectl create -f fuse-device-plugin-k8s-1.16.yml
```

## 使用

参照 [fuse-test.yml](fuse-test.yml)

```yaml
spec: 
  containers:
  - ...
    resources:
      limits:
        github.com/fuse: 1
```

# 特别感谢

![Goland](https://blog.jetbrains.com/wp-content/uploads/2019/01/goland_icon.svg)
