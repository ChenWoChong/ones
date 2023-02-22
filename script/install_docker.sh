#!/bin/env bash

# 卸载docker
# 1. 查看已经安装的docker
yum list installed |grep docker
# 卸载
yum -y remove docker* containerd*
# 删除配置文件
rm -rf /etc/systemd/system/docker.service.d && \
rm -rf /var/lib/docker && \
rm -rf /var/run/docker

echo "设置docker的目录软连接 /var/lib/docker把docker存储设置到数据盘上"
mkdir -p /data1/docker
ln -s /data1/docker  /var/lib

echo "安装docker 依赖 "
yum install yum-utils device-mapper-persistent-data lvm2 -y

## 新增 Docker 仓库。
 yum-config-manager \
 --add-repo \
 https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

 # yum clean all

## 安装 Docker CE.
# yum update -y && yum install -y \
#   containerd.io-1.2.13 \
#   docker-ce-19.03.11 \
#   docker-ce-cli-19.03.11

## 安装 Docker 最新版
yum install -y docker-ce

## 创建 /etc/docker 目录。
mkdir /etc/docker || true
cat > /etc/docker/daemon.json <<EOF
{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2",
  "storage-opts": [
    "overlay2.override_kernel_check=true"
  ],
  "registry-mirrors": [ "https://1nj0zren.mirror.aliyuncs.com", "https://docker.mirrors.ustc.edu.cn", "http://f1361db2.m.daocloud.io", "https://registry.docker-cn.com" ]
}
EOF

systemctl daemon-reload
systemctl restart docker
sudo systemctl enable docker