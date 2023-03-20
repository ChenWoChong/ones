#!/usr/bin/env bash

# 获取脚本所在目录的路径
script_dir="$(cd "$(dirname "$0")" && pwd)"

readonly node_version20=kindest/node:v1.20.15
readonly node_version22=kindest/node:v1.22.15
readonly node_version24=kindest/node:v1.24.6
readonly node_version26=kindest/node:v1.26.0

readonly m=my-cluster-control-plane
readonly w1=my-cluster-worker
readonly w2=my-cluster-worker2

# get cluster
if [[ "$1" == "gc" ]]; then
  kind get clusters
fi

# create cluster test
case $1 in
"cct20")
  kind create cluster --name test-20 --image $node_version20
  ;;
"cct22")
  kind create cluster --name test-20 --image $node_version22
  ;;
"cct24")
  kind create cluster --name test-20 --image $node_version24
  ;;
"cct26")
  kind create cluster --name test-20 --image $node_version26
  ;;
esac

# create cluster
if [[ "$1" == "cc" ]]; then
  # kind创建集群
  kind create cluster --name my-cluster --image $node_version22 --config "$script_dir"/yaml/my-cluster.yaml
  kubectl cluster-info --context kind-my-cluster
fi

# delete cluster
if [[ "$1" == "dc" ]]; then
  # kind删除集群
  kind delete cluster --name my-cluster
fi

# init node network, install vim
node_install_pkg() {
  # add dns server
  docker exec "$1" sh -c "echo 'nameserver 8.8.8.8' >> /etc/resolv.conf"
  docker exec "$1" sh -c "echo 'nameserver 8.8.4.4' >> /etc/resolv.conf"

  # install pkg
  docker exec "$1" sh -c "apt-get update && apt-get install -y vim"
}

node_start_koord_proxy() {
  #  make build-koord-runtime-proxy
  docker exec "$1" sh -c "nohup /koordinator/bin/koord-runtime-proxy --remote-runtime-service-endpoint=/run/containerd/containerd.sock --remote-image-service-endpoint=/run/containerd/containerd.sock > output.log 2>&1 &"
  docker exec "$1" sh -c "echo ' --container-runtime=remote --container-runtime-endpoint=unix:///var/run/koord-runtimeproxy/runtimeproxy.sock' >> /etc/default/kubelet && \
   systemctl daemon-reload && systemctl restart kubelet"
}

node_init() {
  echo "init node $1"
  node_install_pkg "$1"
  node_start_koord_proxy "$1"
}

# node init
if [[ "$1" == "ni" ]]; then
  node_init "$m"
#  node_init "$w1"
#  node_init "$w2"
fi

# helm install koordlet
helm_install_koordlet() {
  # Firstly add koordinator charts repository if you haven't done this.
  helm repo add koordinator-sh https://koordinator-sh.github.io/charts/

  # [Optional]
  helm repo update

  # Install the latest version.
  helm install koordinator koordinator-sh/koordinator --version 1.1.1
}

if [[ "$1" == "hik" ]]; then
  helm_install_koordlet
fi

# inter node
inter_node() {
  docker exec -it "$1" bash
}

# inter node master
case $1 in
"in")
  echo "inter node: $2"
  inter_node "$2"
  ;;
"inm")
  echo "inter node: $m"
  inter_node "$m"
  ;;
"inw1")
  echo "inter node: $w1"
  inter_node "$w1"
  ;;
"inw2")
  echo "inter node: $w2"
  inter_node "$w2"
  ;;
esac
