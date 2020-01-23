apt-get update && sudo apt-get install -y apt-transport-https curl
curl -s https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | sudo apt-key add -
cp kubernets.list /etc/apt/source.list.d/
apt-get update
apt-get install -y kubelet=1.14.0-00 kubeadm=1.14.0-00 kubectl=1.14.0-00
apt-mark hold kubelet=1.14.0-00 kubeadm=1.14.0-00 kubectl=1.14.0-00
systemctl enable kubelet && systemctl start kubelet

#!/bin/bash
#images=(kube-proxy:v1.17.0 kube-scheduler:v1.17.0 kube-controller-manager:v1.17.0 kube-apiserver:v1.17.0 etcd:3.2.24 pause:3.1 coredns:1.2.2)
#for imageName in ${images[@]} ; do
#  docker pull anjia0532/google-containers.$imageName
#  docker tag anjia0532/google-containers.$imageName k8s.gcr.io/$imageName
#  docker rmi anjia0532/google-containers.$imageName
#done
