apt install -y qemu qemu-kvm libvirt-bin bridge-utils virt-manager
systemctl start libvirtd
systemctl status libvirtd
systemctl enable libvirtd

#启动默认NAT网络
virsh net-autostart default

#磁盘准备
cd /mnt/kvm
qemu-img create -f qcow2 disk-1.img 4G

#下载镜像
#下载镜像
#下载镜像
wget https://stable.release.core-os.net/amd64-usr/current/coreos_production_qemu_image.img.bz2
