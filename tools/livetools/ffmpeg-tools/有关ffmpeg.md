# 有关FFmpeg

## 安装与编译

### 一、安装



### 二、编译

```
./configure --enable-gpl --enable-nonfree --enable-libfdk-aac --enable-libx264 --enable-libx265 --enable-filter=delogo --enable-debug --disable-optimizations --enable-libspeex --enable-videotoolbox --enable-static --enable-pthreads --enable-version3 --enable-hardcoded-tables 

#
sudo apt-get install libfdk-aac-dev libspeex-dev libx264-dev libx265-dev
```

### 二、硬解

#### 源码中支持哪些硬解

```cpp
// libavutil/hwcontext.h
enum AVHWDeviceType {
    AV_HWDEVICE_TYPE_NONE,
    AV_HWDEVICE_TYPE_VDPAU,
    AV_HWDEVICE_TYPE_CUDA,
    AV_HWDEVICE_TYPE_VAAPI,
    AV_HWDEVICE_TYPE_DXVA2,
    AV_HWDEVICE_TYPE_QSV,
    AV_HWDEVICE_TYPE_VIDEOTOOLBOX,
    AV_HWDEVICE_TYPE_D3D11VA,
    AV_HWDEVICE_TYPE_DRM,
    AV_HWDEVICE_TYPE_OPENCL,
    AV_HWDEVICE_TYPE_MEDIACODEC,
};
```

 

## 开源项目

- [zimbatm/ffmpeg-static](https://github.com/zimbatm/ffmpeg-static)

## 其他资料

