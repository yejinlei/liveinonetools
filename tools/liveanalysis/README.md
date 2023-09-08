# 音视频编解码、容器及协议

## 一、编解码

## 二、音视频容器

| protocols（流协议及其他） | 容器格式                           | codecs（音视频编解码）                      | hwaccels（硬件加速） | 协议/文档            | 工具/库                                                      |
| ------------------------- | ---------------------------------- | ------------------------------------------- | -------------------- | -------------------- | ------------------------------------------------------------ |
| _                         | **rtp**                            | _                                           |                      |                      | 1. [rtptools版本库](https://gitee.com/yejinlei-mirror/rtptools) |
| _                         | **flv**(Flash Video)               |                                             |                      |                      | [flv工具](https://share.weiyun.com/5h54Qqp)                  |
| _                         | **mp4**(MPEG-4 part 14)            | h264视频编码+aac音频编码                    |                      | H5 video标签原生支持 | 1. [mp4工具](https://share.weiyun.com/5pM2dwF)<br>2. [GPAC's github](https://github.com/gpac/gpac)<br>3. [mp4box.js's github](https://github.com/gpac/mp4box.js) |
| _                         | **fmp4(fragment mp4)**             | h264视频编码+aac音频编码                    |                      |                      |                                                              |
| _                         | **m4v**(raw MPEG-4 video)          |                                             |                      |                      |                                                              |
| _                         | **webm**                           | vp8视频编码+vorbis音频编码                  |                      | H5 video标签原生支持 |                                                              |
| _                         | **ogg**                            | theora视频编码+vorbis音频编码               |                      | H5 video标签原生支持 |                                                              |
| _                         | **hds**(HDS Muxer)                 |                                             |                      |                      |                                                              |
| _                         | **mpegts**(MPEG2 Transport Stream) | mpeg1video/mpeg2video视频编码+mp1/2音频编码 |                      |                      | 1. [雷霄骅 / TS_Analysis](https://gitee.com/leixiaohua/TS_Analysis)<br>2. [EasyICE](https://share.weiyun.com/LPCJIjEG) |
| _                         | **aac**(Advanced Audio Coding)     |                                             |                      |                      | [雷霄骅 / AAC_Analysis](https://gitee.com/leixiaohua/AAC_Analysis) |
| _                         | **adts**(ADTS AAC)                 |                                             |                      |                      |                                                              |

### 2.1、MP4 容器

​	**mp4** 或称 **MPEG-4 Part 14**，是一种多媒体容器格式，扩展名为`.mp4`。

#### 2.2.1、MP4 标准

#### 2.1.2、工具

##### [bento4](https://www.bento4.com/documentation/mp4info/)

##### [MP4Box](https://gpac.wp.imt.fr/)

#####  [mp4box.js](https://gpac.github.io/mp4box.js/) 一个在线解析mp4的工具

## 三、开源项目

- RTSPtoWebRTC，[源码](https://github1s.com/deepch/RTSPtoWebRTC)

- deepch/vdk，[源码](https://github1s.com/deepch/vdk)

## 四、参考资料

1. [hybridgroup/gocv](https://github.com/hybridgroup/gocv)
2. [abhiTronix/vidgear](https://github.com/abhiTronix/vidgear)
3. [adrg/libvlc-go](https://github.com/adrg/libvlc-go)
4. [andrewekhalel/sewar](https://github.com/andrewekhalel/sewar)
5. [jrottenberg/ffmpeg](https://github1s.com/jrottenberg/ffmpeg)
6. [q191201771/lal](https://github1s.com/q191201771/lal)
7. [dusong7/rtspserver](https://gitee.com/dusong7/rtspserver/tree/master/src/RTSP)
8. [korandiz/v4l](https://github.com/korandiz/v4l)
---
1. [基于OpenCV实现图像质量评价](https://blog.csdn.net/luohenyj/article/details/108984313)
2. [图像增强算法效果评价指标及实现](https://www.coder55.com/article/51157)
3. [FFmpeg原理](https://ffmpeg.xianwaizhiyin.net/cover.html)