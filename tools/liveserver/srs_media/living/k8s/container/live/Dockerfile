#about-cloud-native
FROM registry.cn-hangzhou.aliyuncs.com/yejinlei/ubuntu.build:1.0 AS builder
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/' /etc/apk/repositories && apk update
#RUN apk add perl
WORKDIR /tmp
COPY . .
# srs静态编译
RUN apt-get -qq install unzip tcl cmake --fix-missing
RUN cd /tmp/srs/trunk && chmod a+x ./configure && ./configure --cxx="g++ -static" && make -j8

FROM alpine:3.14 AS runner
COPY --from=builder /tmp/srs/trunk/objs/srs /usr/local/objs/srs
COPY --from=builder /tmp/srs/trunk/conf /usr/local
ADD server /usr/local/bin/server
WORKDIR /usr/local
CMD objs/srs -c conf/srs.conf
