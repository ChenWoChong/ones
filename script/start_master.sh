#!/bin/env bash

kubeadm init \
    --image-repository registry.aliyuncs.com/google_containers \
    --kubernetes-version v1.22.3 \
    --apiserver-advertise-address=192.168.59.3 \
    --pod-network-cidr=10.10.0.0/16 \
    --ignore-preflight-errors=all

# 下载flannel
curl -LO https://github.com/flannel-io/flannel/releases/latest/download/kube-flannel.yml && kubectl apply -f kube-flannel.yml


# 修改 node ip
cat <<EOF | sudo tee -a /usr/lib/systemd/system/kubelet.service.d/10-kubeadm.conf
Environment="KUBELET_EXTRA_ARGS=--node-ip=192.168.59.3"
EOF

# 重启
sudo systemctl daemon-reload
sudo systemctl restart kubelet

