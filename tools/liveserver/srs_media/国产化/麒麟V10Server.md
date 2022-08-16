# SRS媒体服务器 - 国产化麒麟V10 Server版本(Aarch64/Armv8 Centos 7) 

## 一、安装docker环境

cat docker-ce.repo

[docker-ce-stable]
name=Docker CE Stable - $basearch
baseurl=https://mirrors.aliyun.com/docker-ce/linux/centos/7/$basearch/stable
enabled=1
gpgcheck=0
gpgkey=https://mirrors.aliyun.com/docker-ce/linux/centos/gpg

---

yum clean all

yum makecache

yum install docker-ce -y --skip-broken

yum install docker-compose -y --skip-broken

## 二、搭建编译环境(验证可行性)

### 2.1 配置源、拉取镜像、安装编译工具链

- 配置源
```css
#sources.list
deb http://mirrors.ustc.edu.cn/ubuntu-ports/ focal main restricted universe multiverse
deb-src http://mirrors.ustc.edu.cn/ubuntu-ports/ focal main main restricted universe multiverse
deb https://mirrors.ustc.edu.cn/ubuntu-ports/ focal-updates main restricted universe multiverse
deb-src http://mirrors.ustc.edu.cn/ubuntu-ports/ focal-updates main restricted universe multiverse
deb https://mirrors.ustc.edu.cn/ubuntu-ports/ focal-backports main restricted universe multiverse
deb-src http://mirrors.ustc.edu.cn/ubuntu-ports/ focal-backports main restricted universe multiverse
deb http://mirrors.ustc.edu.cn/ubuntu-ports/ focal-security main restricted universe multiverse
```

- 拉取镜像、运行容器
```sh
docker pull arm64v8/ubuntu:20.04 #https://hub.docker.com/r/arm64v8/ubuntu/
docker run --rm -ti  -v `pwd`sources.list:/etc/apt/sources.list arm64v8/ubuntu:20.04 sh
```

- 制作srs容器(需在国产化麒麟V10 Server版本)

```sh
#docker build -f trunk/Dockerfile -t hxdi/arm64v8_ubuntu20.04_srs:v4.0 .
FROM arm64v8/ubuntu:20.04 AS build

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Install depends tools.
ADD sources.list /etc/apt/sources.list
RUN apt-get update
RUN apt-get install -y build-essential cmake unzip pkg-config patch unzip perl-base tcl git

ARG SRS_AUTO_PACKAGER

# Build and install SRS.
COPY . /srs
WORKDIR /srs/trunk
RUN ./configure --srt=on --jobs=2 && make -j2 && make install

# All config files for SRS.
RUN cp -R conf /usr/local/srs/conf && \
    cp research/api-server/static-dir/index.html /usr/local/srs/objs/nginx/html/ && \
    cp research/api-server/static-dir/favicon.ico /usr/local/srs/objs/nginx/html/ && \
    cp research/players/crossdomain.xml /usr/local/srs/objs/nginx/html/ && \
    cp -R research/console /usr/local/srs/objs/nginx/html/ && \
    cp -R research/players /usr/local/srs/objs/nginx/html/ && \
    cp -R 3rdparty/signaling/www/demos /usr/local/srs/objs/nginx/html/

############################################################
# dist
############################################################
FROM arm64v8/ubuntu:20.04 AS dist

# Expose ports for streaming @see https://github.com/ossrs/srs#ports
EXPOSE 1935 1985 8080 8000/udp 10080/udp

# FFMPEG 4.1
#COPY --from=build /usr/local/bin/ffmpeg /usr/local/srs/objs/ffmpeg/bin/ffmpeg
# SRS binary, config files and srs-console.
COPY --from=build /usr/local/srs /usr/local/srs

# Default workdir and command.
WORKDIR /usr/local/srs
CMD ["./objs/srs", "-c", "conf/docker.conf"]
```

```sh
[root@localhost srs.oschina]# sudo docker images 
REPOSITORY                     TAG                 IMAGE ID            CREATED             SIZE
hxdi/arm64v8_ubuntu20.04_srs   v4.0                bb2eec1fc239        6 minutes ago       102MB
```

## 三、Qemu模拟

```sh
qemu-system-aarch64 \
 -m 4096 \
 -cpu cortex-a72 -smp 4,cores=4,threads=1,sockets=1 \
 -M virt \
 -bios QEMU_EFI.fd \
 -net nic,model=pcnet \
 -device nec-usb-xhci \
 -device usb-kbd \
 -device usb-mouse \
 -device VGA \
 -drive if=none,file=./ubuntukylin-20.04.1-enhanced-amd64.iso,id=cdrom,media=cdrom \
 -device virtio-scsi-device \
 -device scsi-cd,drive=cdrom \
 -drive if=none,file=rootfs.qcow2,id=hd0 \
 -device virtio-blk-device,drive=hd0
```

```objectivec
qemu-system-aarch64 \
 -m 4096 \
 -cpu cortex-a72 -smp 4,cores=4,threads=1,sockets=1 \
 -M virt \
 -bios QEMU_EFI.fd \
 -net nic,model=pcnet \
 -device nec-usb-xhci \
 -device usb-kbd \
 -device usb-mouse \
 -device VGA \
 -drive if=none,file=./ubuntukylin-20.04.1-enhanced-amd64.iso,id=cdrom,media=cdrom \
 -device virtio-scsi-device \
 -device scsi-cd,drive=cdrom \
 -drive if=none,file=rootfs.qcow2,id=hd0 \
 -device virtio-blk-device,drive=hd0
```



## 四、参考资料

1. [v4_CN_SrsLinuxArm](https://github.com/ossrs/srs/wiki/v4_CN_SrsLinuxArm#ubuntu-cross-build-srs-armv8aarch64)
2. [Support buildin setjmp/longjmp for ARM](https://github.com/ossrs/state-threads/issues/1)
3. [ARM上部署SRS实例](https://www.wenjiangs.com/doc/srs-samplearm)
4. [qemu kvm 安装kylin server v10系统](https://blog.csdn.net/weixin_42915431/article/details/113887132)
5. [ubuntukylin 20.04 iso](https://mirrors.tuna.tsinghua.edu.cn/ubuntukylin-cdimage/20.04/)
