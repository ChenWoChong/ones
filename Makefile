GIT_SHA = $(shell git rev-list HEAD | awk 'NR==1')
GIT_SHORT_SHA = $(shell git describe --always --abbrev=7 --dirty)
GIT_DATE = $(shell git log --pretty=format:%cd ${GIT_SHA} -1)
GIT_BRANCH = $(shell git branch | sed -n '/\* /s///p')
GIT_TAG_LABLE = $(shell git tag --sort=-taggerdate|head -n '1')
GIT_SVR_PATH = $(shell git remote -v | awk 'NR==1' | sed 's/[()]//g' | sed 's/\t/ /g' |cut -d " " -f2)

init:
	ssh -p 2231 root@localhost 'bash -s' < ./script/init_server.sh

docker:
	ssh -p 2231 root@localhost 'bash -s' < ./script/install_docker.sh

kubeadm:
	ssh -p 2231 root@localhost 'bash -s' < ./script/install_kubeadm.sh

test:
	@echo ${GIT_SHA}
	@echo ${GIT_SHORT_SHA}
	@echo ${GIT_DATE}
	@echo ${GIT_BRANCH}
	@echo ${GIT_TAG_LABLE}
	@echo ${GIT_SVR_PATH}