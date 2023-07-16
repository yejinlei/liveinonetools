FROM registry.cn-hangzhou.aliyuncs.com/yejinlei/multi.ubuntu.gcc:11 as build

ARG nasm_version=2.15.05
ARG x264_version=master
ARG x265_version=3.5
ARG zimg_version=release-3.0.4
ARG fdk_aac_version=v2.0.2
ARG opus_version=v1.3.1
ARG libvpx_version=v1.13.0
ARG libaom_version=v3.6.0
ARG svt_av1_version=v1.4.0
ARG vmaf_version=v2.3.1
ARG ffmpeg_version=6.0

ENV DEBIAN_FRONTEND noninteractive

# update, and install basic packages
RUN apt-get update -qq && \
    apt-get upgrade -y && \
    apt-get -y install --no-install-recommends \
    autoconf \
    automake \
    build-essential \
    nasm \
    yasm \
    python3 \
    ninja-build \
    meson \
    cmake \
    git-core \
    libfreetype6-dev \
    libtool \
    openssl \
    libssl-dev \
    pkg-config \
    texinfo \
    wget \
    zlib1g-dev \
    curl \
    git \
    mercurial \
    libnuma-dev \
    ca-certificates


ENV TZ=UTC
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN mkdir -p /opt/ffmpeg/bin
ENV PATH="/opt/ffmpeg/bin:$PATH"
ENV PKG_CONFIG_PATH="/opt/ffmpeg/lib/pkgconfig"

# WORKDIR /tmp
# # RUN curl -sS -O https://www.nasm.us/pub/nasm/releasebuilds/${nasm_version}/nasm-${nasm_version}.tar.xz
# RUN curl -sS -O https://ftp.osuosl.org/pub/blfs/conglomeration/nasm/nasm-${nasm_version}.tar.xz
# RUN tar xf nasm-${nasm_version}.tar.xz
# WORKDIR /tmp/nasm-${nasm_version}
# RUN ./autogen.sh && ./configure --prefix="/opt/ffmpeg" --bindir="/opt/ffmpeg/bin"
# RUN make -j$(nproc)
# RUN make install

#https://github.com/sekrit-twc/zimg
WORKDIR /tmp/zimg
RUN git clone --branch ${zimg_version} --depth 1 https://gitee.com/yejinlei/zimg.git .
RUN ./autogen.sh
RUN ./configure --enable-static  --prefix=/opt/ffmpeg --disable-shared
RUN make -j $(nproc)
RUN make install

#https://github.com/corecodec/x264.git
WORKDIR /tmp/x264
RUN git clone --branch ${x264_version} --depth 1 https://gitee.com/yejinlei/x264.git .
RUN ./configure --prefix="/opt/ffmpeg" --bindir="/opt/ffmpeg/bin" --enable-static --enable-pic
RUN make -j$(nproc)
RUN make install

WORKDIR /tmp/x265
RUN git clone --branch ${x265_version} --depth 1 https://bitbucket.org/multicoreware/x265_git.git .
WORKDIR build/linux
RUN cmake -G "Unix Makefiles" -DCMAKE_INSTALL_PREFIX="/opt/ffmpeg" -DENABLE_SHARED=off ../../source
RUN make -j$(nproc)
RUN make install

#https://github.com/webmproject/libvpx.git
WORKDIR /tmp/libvpx
RUN git clone --branch ${libvpx_version} --depth 1 https://gitee.com/yejinlei/libvpx.git .
RUN ./configure --prefix="/opt/ffmpeg" --disable-examples --disable-unit-tests --enable-vp9-highbitdepth --as=yasm
RUN make -j$(nproc)
RUN make install

#https://github.com/mstorsjo/fdk-aac
WORKDIR /tmp/fdk-aac
RUN git clone --branch ${fdk_aac_version} --depth 1 https://gitee.com/yejinlei/fdk-aac.git .
RUN autoreconf -fiv && ./configure --prefix="/opt/ffmpeg" --disable-shared
RUN make -j$(nproc)
RUN make install

#https://github.com/xiph/opus.git
WORKDIR /tmp/opus
RUN git clone --branch ${opus_version} --depth 1 https://gitee.com/yejinlei/opus.git .
RUN ./autogen.sh
RUN ./configure --prefix="/opt/ffmpeg" --disable-shared
RUN make -j$(nproc)
RUN make install

#https://github.com/jbeich/aom
WORKDIR /tmp/aom
RUN git -C aom pull 2> /dev/null || git clone --branch ${libaom_version} --depth 1 https://gitee.com/yejinlei/aom.git .
WORKDIR /tmp/aom_build
RUN cmake -G "Unix Makefiles" -DCMAKE_INSTALL_PREFIX="/opt/ffmpeg" -DBUILD_SHARED_LIBS=off -DENABLE_TOOLS=off -DENABLE_DOCS=off -DENABLE_EXAMPLES=off -DENABLE_TESTS=off -DENABLE_NASM=on ../aom
RUN make -j$(nproc)
RUN make install

#https://github.com/Netflix/vmaf.git
WORKDIR /tmp/vmaf
RUN git clone --branch ${vmaf_version} --depth 1 https://gitee.com/yejinlei/vmaf.git .
WORKDIR libvmaf/build
RUN meson .. --default-library=static --prefix=/opt/ffmpeg --libdir=/opt/ffmpeg/lib --buildtype=release
RUN ninja -vC . install

WORKDIR /tmp/svt-av1
RUN git clone --branch ${svt_av1_version} --depth 1 https://gitlab.com/AOMediaCodec/SVT-AV1.git .
WORKDIR Build/linux
RUN cmake -G "Unix Makefiles" -DCMAKE_INSTALL_PREFIX="/opt/ffmpeg" -DBUILD_SHARED_LIBS=off -DCMAKE_BUILD_TYPE=Release ../../
RUN make -j $(nproc)
RUN make install

WORKDIR /tmp
RUN curl -sS -O https://ffmpeg.org/releases/ffmpeg-${ffmpeg_version}.tar.bz2
WORKDIR ffmpeg-${ffmpeg_version}
RUN tar xjf ../ffmpeg-${ffmpeg_version}.tar.bz2 --strip-components 1
#RUN curl -sS -o svt-av1-consolidated.patch https://github.com/cosmin/FFmpeg/commit/9a151b0edee710727a588dddd0b331b4dc676227.patch
#RUN patch -p1 < svt-av1-consolidated.patch

RUN    ./configure \
    --prefix="/opt/ffmpeg" \
    --ld="g++" \
    --pkg-config-flags="--static" \
    --pkg-config-flags="--static" \
    --extra-cflags="-I/opt/ffmpeg/include" \
    --extra-ldflags="-L/opt/ffmpeg/lib" \
    --extra-ldexeflags="-Bstatic" \
    --extra-libs="-lpthread -lm" \
    --bindir="/opt/ffmpeg/bin" \
    --disable-shared \
    --enable-static \
    --disable-ffplay \
    --enable-gpl \
    --enable-nonfree \
    --enable-version3 \
    --enable-zlib \
    --enable-libzimg \
    --enable-libaom \
    --enable-libfdk-aac \
    --enable-libfreetype \
    --enable-libopus \
    --enable-libvpx \
    --enable-libx264 \
    --enable-libx265 \
    --enable-libvmaf \
    --enable-libsvtav1 \
    --enable-openssl
RUN make -j$(nproc)
RUN make install

FROM ubuntu:22.04
LABEL maintainer "yejinlei <4010896@qq.com>"
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update -qq && apt-get upgrade -y && \
    apt-get -y install --no-install-recommends \
    libnuma1 \
    libssl3 \
    libfreetype6 \
    && apt-get -y clean && rm -r /var/lib/apt/lists/*

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /opt/ffmpeg
COPY --from=build /opt/ffmpeg .

RUN ln -s /opt/ffmpeg/share/model /usr/local/share/
RUN ldconfig

ENV PATH="/opt/ffmpeg/bin:$PATH"
ENV PKG_CONFIG_PATH="/opt/ffmpeg/lib/pkgconfig"
ENV LD_LIBRARY_PATH="/opt/ffmpeg/lib"

WORKDIR /root
