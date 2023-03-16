#about-cloud-native
FROM registry.cn-hangzhou.aliyuncs.com/yejinlei/ubuntu.build:1.0 AS builder
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/' /etc/apk/repositories && apk update
#RUN apk add perl
WORKDIR /tmp
COPY . .
# srs静态编译
RUN apt-get update && apt-get -qq install unzip tcl cmake --fix-missing
#RUN cd /tmp/srs/trunk && chmod a+x ./configure && ./configure --cxx="g++ -static" && make -j8
RUN cd /tmp/srs/trunk && chmod a+x ./configure && ./configure --hds=on --nginx=off --srtp-nasm=off --ffmpeg-fit=off && make -j8 && make install

FROM ubuntu:focal AS runner
RUN apt-get update && apt-get -qq install python3 python3-pip --fix-missing

COPY --from=builder /tmp/srs/trunk/objs/srs /usr/local/srs/objs/srs
COPY --from=builder /tmp/srs/trunk/conf /usr/local/srs/conf
ADD server.py /usr/local/srs/server.py

WORKDIR /usr/local/srs
#ENV SRS_ENV_ONLY=on 
#ENV SRS_EXPORTER_ENABLED=on
CMD objs/srs -c conf/srs.conf
