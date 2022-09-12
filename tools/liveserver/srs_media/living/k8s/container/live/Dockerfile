#about-cloud-native
FROM registry.cn-hangzhou.aliyuncs.com/yejinlei/ubuntu.build:1.0 AS builder
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/' /etc/apk/repositories && apk update
#RUN apk add perl
WORKDIR /tmp
COPY . .
# for srs
RUN apt-get -qq install unzip tcl cmake
RUN cd /tmp/srs/trunk && chmod a+x ./configure && ./configure --static=on && make -j8
FROM alpine:3.14 AS runner
COPY --from=builder /tmp/srs/trunk/objs/srs /opt/objs/srs
COPY --from=builder /tmp/srs/trunk/conf /opt
WORKDIR /opt
CMD objs/srs -c conf/srs.conf

