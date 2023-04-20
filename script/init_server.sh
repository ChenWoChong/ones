#!/bin/env bash

echo "change ps1"
echo 'export PS1="\n\e[1;37m[\e[m\e[1;32m\u\e[m\e[1;33m@\e[m\e[1;35m\H\e[m \e[4m`pwd`\e[m\e[1;37m]\e[m\e[1;36m\e[m\n\$ "' >> /etc/bashrc && source /etc/bashrc

### 修改host
# hostnamectl set-hostname centos1

echo "清理一些无用的yum repos"
yum clean all

### 关闭swap分区
echo "swapoff ## 临时生效"
swapoff -a

## 永久生效可以编辑/etc/fstab文件 注释掉swap那行
sed -ri '/\sswap\s/s/^#?/#/' /etc/fstab

echo "关闭防火墙"
systemctl stop firewalld
systemctl disable --now firewalld

echo "关闭selinux"
setenforce 0
sed -i 's/enforcing/disabled/' /etc/selinux/config

# 内核开启IPv4转发需要开启下面的模块
modprobe br_netfilter

cat >/etc/sysctl.d/k8s.conf <<EOF
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
EOF

### 安装基本软件
yum install -y net-tools
