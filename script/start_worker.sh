#!/bin/env bash

kubeadm join 192.168.59.3:6443 --token 2t4v53.s5ggnqoczvzezkgu \
        --discovery-token-ca-cert-hash sha256:ef376f292cde4a429a7492bddd74f905f6a95ac53d558f6be1832783ef810e27


# kubeadm的启动配置中，追加node ip

cat <<EOF | sudo tee -a /usr/lib/systemd/system/kubelet.service.d/10-kubeadm.conf
Environment="KUBELET_EXTRA_ARGS=--node-ip=192.168.59.4"
EOF

# 重启
sudo systemctl daemon-reload
sudo systemctl restart kubelet