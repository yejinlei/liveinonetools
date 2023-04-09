#docker buildx，编译多平台镜像
sudo docker run --privileged --rm tonistiigi/binfmt --install all
#x86架构下运行其他平台容器
sudo docker run --rm --privileged multiarch/qemu-user-static:register
