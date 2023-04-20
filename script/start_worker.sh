#!/bin/env bash

kubeadm join 192.168.59.7:6443 --token 3f6dt2.hz8wy38z6v5r8kue \
        --discovery-token-ca-cert-hash sha256:bb50878bda80cd39354ebecf276386f9948415b09861265a423c336b479a33e1


### kubectl 权限
scp root@192.168.59.7:/etc/kubernetes/admin.conf /etc/kubernetes/admin.conf
mkdir -p $HOME/.kube && \
sudo \cp -rf /etc/kubernetes/admin.conf $HOME/.kube/config && \
sudo chown $(id -u):$(id -g) $HOME/.kube/config