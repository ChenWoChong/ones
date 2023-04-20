#!/bin/env bash

cat > /etc/yum.repos.d/kubernetes.repo <<EOF
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64/
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
exclude=kube*
EOF


kubeVersion=1.22.3
yum install -y kubelet-$kubeVersion kubeadm-$kubeVersion kubectl-$kubeVersion --disableexcludes=kubernetes

kubeadm version

### 锁定版本，防止意外升级
yum install -y yum-plugin-versionlock && yum versionlock kubeadm kubelet kubectl

systemctl enable kubelet
systemctl daemon-reload
systemctl restart kubelet