# 有关FFmpeg

## 一、安装与编译

### 1.1、安装

### 1.2、编译

#### 静态编译


#### configure文件

    echo "1. 查看所有可用的编码器"
    echo "2. 查看所有可用的解码器"
    echo "3. 查看所有可用的硬件加速器"
    echo "4. 查看所有可用的解封装器"
    echo "5. 查看所有可用的封装器"
    echo "6. 查看所有可用的解析器"
    echo "7. 查看所有可用的协议"
    echo "8. 查看所有可用的比特流过滤器"
    echo "9. 查看所有可用的输入设备"
    echo "10. 查看所有可用的输出设备"
    echo "11. 查看所有可用的过滤器"
    1) ./configure --list-encoders;;
    2) ./configure --list-decoders;;
    3) ./configure --list-hwaccels;;
    4) ./configure --list-demuxers;;
    5) ./configure --list-muxers;;
    6) ./configure --list-parsers;;
    7) ./configure --list-protocols;;
    8) ./configure --list-bsfs;;
    9) ./configure --list-indevs;;
    10) ./configure --list-outdevs;;
    11) ./configure --list-filters;;

```
./configure --enable-gpl --enable-nonfree --enable-libfdk-aac --enable-libx264 --enable-libx265 --enable-filter=delogo --enable-debug --disable-optimizations --enable-libspeex --enable-videotoolbox --enable-static --enable-pthreads --enable-version3 --enable-hardcoded-tables 

#
sudo apt-get install libfdk-aac-dev libspeex-dev libx264-dev libx265-dev
```

## 二、源码分析

### 2.1、整体结构
- libavcodec：音视频编解码模块，实现了多种音视频编解码算法，例如H.264、AAC、MP3等。
- libavdevice：设备输入输出模块，实现了多种音视频设备的输入输出，例如摄像头、麦克风等。
- libavfilter：滤波器模块，实现了多种音视频滤波器，例如降噪、去除水印等。
- libavformat：封装格式模块，实现了多种音视频封装格式的解析和生成，例如MP4、AVI、FLV等。
- libavutil：工具函数模块，实现了多种常用的工具函数，例如内存管理、时间戳计算等。

### 2.2、libavcodec: 音视频编解码模块

> libavcodec是FFmpeg的核心库之一，实现了多种音视频编解码算法。它包含了多个编解码器和解码器，可以对多种音视频格式进行编解码。
>
> libavcodec模块的主要作用是对音视频进行编解码，以实现音视频的播放、转码、编辑等功能。它可以在FFmpeg的多个模块中使用，例如ffmpeg、ffplay、ffprobe等。
>
> libavcodec模块的源代码位于libavcodec目录下，包含多个头文件和源文件。在编译FFmpeg时，libavcodec模块是默认启用的，无需使用额外的编译选项。
>
> libavcodec模块支持多种音视频编解码格式，例如H.264、MPEG-4、AAC、MP3等。它还支持多种编解码器和解码器，例如libx264、libvpx、libmp3lame等。
>
> libavcodec模块的使用非常灵活，可以通过命令行参数或API接口进行调用。以下是一个使用libavcodec模块的示例命令：
>
> ```sh
> ffmpeg -i input.mp4 -c:v libx264 -preset medium -tune film -c:a copy output.mp4
> ```
>
> 该命令将输入视频文件input.mp4进行编码，并使用libx264编码器、medium预设和film调整参数。同时，它还将音频流直接复制到输出文件中，以保持音频的原始格式。

#### 2.2.1、源码中支持哪些硬解

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

#### 2.2.2、参考资料

- [ffmpeg HWAccelIntro](https://trac.ffmpeg.org/wiki/HWAccelIntro)
- [w3 webcodecs](https://w3c.github.io/webcodecs)

### 2.3、libavdevice：设备输入输出模块

> libavdevice是FFmpeg的一个可选模块，用于音视频设备的输入和输出。它提供了多种音视频设备的输入和输出接口，例如摄像头、麦克风、扬声器等。
>
> libavdevice模块的主要作用是对音视频设备进行输入和输出，以实现音视频的采集、录制、播放等功能。它可以在FFmpeg的多个模块中使用，例如ffmpeg、ffplay、ffprobe等。
>
> libavdevice模块的源代码位于libavdevice目录下，包含多个头文件和源文件。在编译FFmpeg时，可以使用--enable-avdevice选项启用libavdevice模块。
>
> libavdevice模块支持多种音视频设备的输入和输出，例如V4L2、ALSA、PulseAudio等。它还支持多种音视频设备的输入和输出接口，例如Video4Linux2、ALSA、DirectSound等。
>
> libavdevice模块的使用非常灵活，可以通过命令行参数或API接口进行调用。以下是一个使用libavdevice模块的示例命令：
>
> ```sh
> ffmpeg -f v4l2 -i /dev/video0 -f alsa -i hw:0 -c:v libx264 -preset ultrafast -c:a aac output.mp4
> ```
>
> 该命令将从Video4Linux2设备/dev/video0和ALSA设备hw:0中读取音视频数据，并使用libx264编码器和AAC编码器进行编码。同时，它还将音视频数据写入输出文件output.mp4中。

### 2.4、libavfilter：滤波器模块

> libavfilter是FFmpeg的一个可选模块，用于音视频滤镜和效果处理。它提供了多种音视频滤镜和效果，例如裁剪、缩放、旋转、模糊、去噪等。
>
> libavfilter模块的主要作用是对音视频进行滤镜和效果处理，以实现音视频的特效和处理。它可以在编码和解码过程中使用，也可以在视频播放过程中使用。
>
> libavfilter模块的源代码位于libavfilter目录下，包含多个头文件和源文件。在编译FFmpeg时，可以使用--enable-avfilter选项启用libavfilter模块。
>
> libavfilter模块的使用非常灵活，可以通过命令行参数或API接口进行调用。以下是一个使用libavfilter模块的示例命令：
>
> ```sh
> ffmpeg -i input.mp4 -vf "crop=640:360:0:0" output.mp4
> ```
>
> 该命令将输入视频文件input.mp4进行裁剪，并将其分辨率调整为640x360。同时，它还使用了libavfilter模块中的crop滤镜，对视频进行裁剪处理。

### 2.5、libavformat：封装格式模块

> libavformat是FFmpeg的一个核心库之一，用于音视频格式的封装和解封装。它包含了多个音视频封装器和解封装器，可以对多种音视频格式进行封装和解封装。
>
> libavformat模块的主要作用是对音视频进行封装和解封装，以实现音视频的播放、转码、编辑等功能。它可以在FFmpeg的多个模块中使用，例如ffmpeg、ffplay、ffprobe等。
>
> libavformat模块的源代码位于libavformat目录下，包含多个头文件和源文件。在编译FFmpeg时，libavformat模块是默认启用的，无需使用额外的编译选项。
>
> libavformat模块支持多种音视频格式的封装和解封装，例如MP4、AVI、FLV、MKV等。它还支持多种音视频封装器和解封装器，例如libx264、libvpx、libmp3lame等。
>
> libavformat模块的使用非常灵活，可以通过命令行参数或API接口进行调用。以下是一个使用libavformat模块的示例命令
>
> ```sh
> ffmpeg -i input.mp4 -c:v libx264 -preset medium -tune film -c:a copy output.mkv
> ```
>
> 该命令将输入视频文件input.mp4进行重新封装，并使用libx264编码器、medium预设和film调整参数。同时，它还将音频流直接复制到输出文件中，以保持音频的原始格式。

### 2.6、libpostproc: 视频后处理模块
> libpostproc是FFmpeg的一个可选模块，用于视频后处理。它提供了多种视频滤镜和效果，例如去隔行、去噪、锐化、模糊等。
>
> libpostproc模块的主要作用是对视频进行后处理，以提高视频的质量和清晰度。它可以在编码和解码过程中使用，也可以在视频播放过程中使用。
>
> libpostproc模块的源代码位于libpostproc目录下，包含多个头文件和源文件。在编译FFmpeg时，可以使用--enable-postproc选项启用libpostproc模块。
>
> libpostproc模块的使用非常灵活，可以通过命令行参数或API接口进行调用。以下是一个使用libpostproc模块的示例命令：
>
> ```sh
> ffmpeg -i input.mp4 -c:v libx264 -preset medium -tune film -c:a copy -vf "pp=ac" output.mp4
> ```
>
> 该命令将输入视频文件input.mp4进行编码，并使用libx264编码器、medium预设和film调整参数。同时，它还使用了libpostproc模块中的pp滤镜，对视频进行去隔行处理。

### 2.7、libswscale: 视频图像转换模块
> libswscale是FFmpeg的一个可选模块，用于视频图像缩放和色彩空间转换。它提供了多种图像缩放算法和色彩空间转换算法，可以将视频图像从一种分辨率和色彩空间转换为另一种分辨率和色彩空间。
>
> libswscale模块的主要作用是对视频图像进行缩放和色彩空间转换，以适应不同的显示设备和播放环境。它可以在编码和解码过程中使用，也可以在视频播放过程中使用。
>
> libswscale模块的源代码位于libswscale目录下，包含多个头文件和源文件。在编译FFmpeg时，可以使用--enable-swscale选项启用libswscale模块。
>
> libswscale模块的使用非常灵活，可以通过命令行参数或API接口进行调用。以下是一个使用libswscale模块的示例命令：
>
> ```sh
> ffmpeg -i input.mp4 -vf scale=640:360 output.mp4
> ```
>
> 该命令将输入视频文件input.mp4进行缩放，并将其分辨率调整为640x360。同时，它还使用了libswscale模块中的scale滤镜，对视频进行缩放处理。

### 2.8、libswresample: 音频重采样和格式转换模块

> libswresample是FFmpeg的一个可选模块，用于音频重采样和格式转换。它提供了多种音频重采样算法和格式转换算法，可以将音频从一种采样率、声道数和格式转换为另一种采样率、声道数和格式。
>
> libswresample模块的主要作用是对音频进行重采样和格式转换，以适应不同的播放设备和播放环境。它可以在编码和解码过程中使用，也可以在音频播放过程中使用。
>
> libswresample模块的源代码位于libswresample目录下，包含多个头文件和源文件。在编译FFmpeg时，可以使用--enable-swresample选项启用libswresample模块。
>
> libswresample模块的使用非常灵活，可以通过命令行参数或API接口进行调用。以下是一个使用libswresample模块的示例命令：
>
> ```sh
> ffmpeg -i input.mp3 -c:a pcm_s16le -ar 44100 -ac 2 output.wav
> ```
>
> 该命令将输入音频文件input.mp3进行格式转换和重采样，并将其转换为PCM_S16LE格式、采样率为44100Hz、声道数为2。同时，它还使用了libswresample模块中的音频重采样算法和格式转换算法，对音频进行处理。

### 2.9、libavutil：工具函数模块

> 根据您提供的信息，libavutil是FFmpeg的一个核心库之一，提供了多种常用的工具函数和数据结构。它包含了多个头文件和源文件，可以在FFmpeg的多个模块中使用。
>
> libavutil模块的主要作用是提供常用的工具函数和数据结构，以便其他模块进行调用。它包含了多种常用的工具函数和数据结构，例如内存管理、字符串处理、时间处理、数据结构等。
>
> libavutil模块的源代码位于libavutil目录下，包含多个头文件和源文件。在编译FFmpeg时，libavutil模块是默认启用的，无需使用额外的编译选项。
>
> libavutil模块支持多种常用的工具函数和数据结构，例如AVBuffer、AVFrame、AVDictionary等。它还支持多种常用的工具函数，例如av_malloc()、av_free()、av_gettime()等。
>
> libavutil模块的使用非常灵活，可以通过命令行参数或API接口进行调用。以下是一个使用libavutil模块的示例命令：
>
> ```sh
> ffmpeg -i input.mp4 -t 10 -c copy -bsf:v h264_mp4toannexb -f h264 output.h264
> ```
>
> 该命令将输入视频文件input.mp4进行重新封装，并将其转换为H.264裸流格式。同时，它还使用了libavutil模块中的av_malloc()和av_free()函数，对内存进行动态分配和释放。

## 二、开源项目

- [jrottenberg/ffmpeg/](https://github.com/jrottenberg/ffmpeg/)
- [BtbN/FFmpeg-Builds](https://github.com/BtbN/FFmpeg-Builds)
- [zimbatm/ffmpeg-static](https://github.com/zimbatm/ffmpeg-static)
- [FFmpeg Static Builds](https://www.johnvansickle.com/ffmpeg/)
- [cosmin/docker-ffmpeg](https://github.com/cosmin/docker-ffmpeg/)

## 三、其他资料

- [FFmpeg官方网站](https://www.ffmpeg.org/)
- [FFmpeg文档](https://www.ffmpeg.org/documentation.html)
- [FFmpeg API参考](https://www.ffmpeg.org/doxygen/trunk/index.html)
- [FFmpeg示例代码](https://github.com/FFmpeg/FFmpeg/tree/master/doc/examples)
- [FFmpeg开发者社区](https://trac.ffmpeg.org/wiki/Community)
- [FFmpeg wiki]()
  - [Capture/Webcam](https://trac.ffmpeg.org/wiki/Capture/Webcam)
- [流媒体视频基础 MSE 入门 & FFmpeg 制作视频预览缩略图和 fmp4](https://juejin.cn/post/6953777965838630926)
