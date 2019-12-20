apt-get update && sudo apt-get install -y apt-transport-https curl
curl -s https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | sudo apt-key add -
cp kubernets.list /etc/apt/source.list.d/
apt-get update
apt-get install -y kubelet=1.17.0-00 kubeadm=1.17.0-00 kubectl=1.17.0-00
apt-mark hold kubelet=1.17.0-00 kubeadm=1.17.0-00 kubectl=1.17.0-00
systemctl enable kubelet && systemctl start kubelet
