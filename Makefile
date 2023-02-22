
init:
	ssh -p 2231 root@localhost 'bash -s' < ./script/init_server.sh

docker:
	ssh -p 2231 root@localhost 'bash -s' < ./script/install_docker.sh

kubeadm:
	ssh -p 2231 root@localhost 'bash -s' < ./script/install_docker.sh
