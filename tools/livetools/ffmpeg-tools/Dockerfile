FROM ubuntu:18.04 as build
ENV DEBIAN_FRONTEND noninteractive

RUN echo 'nameserver 114.114.114.114' >> /etc/resolv.conf && echo 'nameserver 8.8.4.4' >> /etc/resolv.conf \
	&& sed -i s/archive.ubuntu.com/mirrors.aliyun.com/g /etc/apt/sources.list \
    && sed -i s/security.ubuntu.com/mirrors.aliyun.com/g /etc/apt/sources.list \
    #&& apt-get update && apt-get upgrade
    && apt-get update \
	&& apt-get -y install git build-essential \
    && apt-get -y install yasm librtmp-dev libx264-dev libx265-dev libfdk-aac-dev 

ADD build.sh build.sh
RUN ./build.sh 
#RUN rm -rf FFmpeg-QuQi-H265-FLV-RTMP &&  apt-get clean

FROM ubuntu:18.04
COPY --from=build ./FFmpeg-QuQi-H265-FLV-RTMP/ffmpeg /usr/bin/ffmpeg
COPY --from=build /usr/lib/x86_64-linux-gnu/ /usr/lib/x86_64-linux-gnu/
