FROM ubuntu:18.04
MAINTAINER yejinlei 4010896@qq.com

ENV DEBIAN_FRONTEND noninteractive

ADD nginx-1.15.2.tar.gz /usr/local/src/
ADD v1.2.1.tar.gz /usr/local/src/
ADD pcre-8.40.tar.gz /usr/local/src/
ADD openssl_1.0.1f.orig.tar.gz /usr/local/src/
ADD zlib-1.2.11.tar.gz /usr/local/src/
ADD sources.list /etc/apt/
RUN apt-get update
RUN apt-get install build-essential libtool gcc g++ make -y


#ADD cd /usr/local/src/pcre-8.40 && ./configure && make && make install
#ADD cd /usr/local/src/zlib-1.2.11 && ./configure && make && make install

WORKDIR /usr/local/src/nginx-1.15.2/
RUN ./configure --prefix=/usr/local/nginx --add-module=../nginx-rtmp-module-1.2.1 --with-pcre=../pcre-8.40 --with-openssl=../openssl-1.0.1f --with-http_ssl_module --with-zlib=../zlib-1.2.11 && make && make install
#RUN ./configure --prefix=/usr/local/nginx --add-module=../nginx-rtmp-module-1.2.1 && make --with-zlib=../zlib-1.2.11 && make install
RUN rm -rf /usr/local/src

ENV PATH /usr/local/nginx/sbin:$PATH
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]


