#!/bin/env bash

nodeIP=192.168.59.7
kubeadm init \
    --image-repository registry.aliyuncs.com/google_containers \
    --kubernetes-version v1.22.3 \
    --apiserver-advertise-address=$nodeIP \
    --pod-network-cidr=10.10.0.0/16 \
    --ignore-preflight-errors=all

### kubectl 权限
mkdir -p $HOME/.kube && \
sudo \cp -rf /etc/kubernetes/admin.conf $HOME/.kube/config && \
sudo chown $(id -u):$(id -g) $HOME/.kube/config

### install flannel
bash install_flannel.sh

master::changeNodeIP() {
    ## 修改 node ip
    cat <<EOF | sudo tee -a /usr/lib/systemd/system/kubelet.service.d/10-kubeadm.conf
Environment="KUBELET_EXTRA_ARGS=--node-ip=$nodeIP"
EOF

    # 重启
    sudo systemctl daemon-reload
    sudo systemctl restart kubelet
}
