ufw disable
modprobe br_netfilter
echo 1 > /proc/sys/net/bridge/bridge-nf-call-iptables
setenforce 0
echo "SELINUX=permissive" >> /etc/selinux/config
echo "net.ipv4.ip_forward = 1" >> /etc/sysctl.conf
sysctl -p
swapoff -a
cp k8s.conf /etc/sysctl.d/
sysctl --system
cp daemon.json /etc/docker/
systemctl restart docker
#hostnamectl set-hostname master
