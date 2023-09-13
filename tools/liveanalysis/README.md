| protocols（流协议及其他）                       | formats(封装容器,通常为本地音视文件格式(如mp4)及适用于网络传输的传输流及其本地文件形式(如fmp4,ts/ps流)) | codecs（音视频编解码）                                       | hwaccels（硬件加速） | 协议/文档                                                    | 工具/库                                                      |
| ----------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | -------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| **rtmp**                                        | **flv tag** 与 flv容器率略有差别                             |                                                              |                      |                                                              | 1. [librtmp's github](https://github.com/yixia/librtmp/tree/master/librtmp)<br>2. [nginx-rtmp-module(支持h265)代码](https://github.com/adwpc/nginx-rtmp-module) |
| **rtmps**                                       | **flv tag **与 flv容器率略有差别                             |                                                              |                      |                                                              |                                                              |
| **rtmpt**                                       | **flv tag** 与 flv容器率略有差别                             |                                                              |                      |                                                              |                                                              |
| rtmpts****                                      | **flv tag** 与 flv容器率略有差别                             |                                                              |                      |                                                              |                                                              |
| **udt（可靠udp协议）**                          |                                                              |                                                              |                      |                                                              |                                                              |
| **srt（基于udt）**                              | **mpegts/TS**                                                |                                                              |                      |                                                              | 1. [srt's github](https://github.com/Haivision/srt)<br>2. [srt-live-server's github](https://github.com/Edward-Wu/srt-live-server) |
| **hls**(Apple HTTP Live Streaming)              | **m3u8+mpegts/TS部分**                                       |                                                              |                      | 1. H5 video标签原生支持 or H5 MSE扩展<br>2. 协议 - [HTTP Live Streaming draft-pantos-http-live-streaming-12](https://tools.ietf.org/html/draft-pantos-http-live-streaming-12)<br>3. HLS协议编码格式要求：<br>- 视频的编码格式：H264<br>音频的编码格式：AAC、MP3、AC-3<br>- 视频的封装格式：ts<br>- 保存 ts 索引的 m3u8 文件 |                                                              |
| **hls+**                                        |                                                              |                                                              |                      | 非标准                                                       | [PingOS](https://pingos.io/)                                 |
| **lhls(Low Latency HLS)**                       |                                                              |                                                              |                      | 非标准                                                       |                                                              |
| **llhls**                                       |                                                              |                                                              |                      | 非标准                                                       |                                                              |
| **dash** (Dynamic adaptive streaming over HTTP) | **mpd+m4v/m4a\|webm\|3gp**                                   |                                                              |                      |                                                              | 1. [video.js dash插件](https://github.com/videojs/videojs-contrib-dash.git)<br>2. [dash测试页](http://reference.dashif.org/dash.js/nightly/samples/dash-if-reference-player/index.html)<br>3. [Bento4's github](https://github.com/axiomatic-systems/Bento4)<br>4. [Bento4's Bin](https://share.weiyun.com/5w4fPrC)<br>5. [DASHEncoder2's github](https://github.com/zhanghuicuc/DASHEncoder2/tree/dev-v2)<br>6. [DASH website demo's github（类似bitcodin网站）](https://github.com/zhanghuicuc/dash_website_demo) |
| **dash-ll**                                     |                                                              |                                                              |                      |                                                              |                                                              |
| **http-flv** (flv over http)                    | **flv**                                                      |                                                              |                      | H5 MSE扩展                                                   |                                                              |
| **websocket-flv** (flv over ws)                 | **flv**                                                      |                                                              |                      | H5 MSE扩展                                                   |                                                              |
| **fmp4** (fragment mp4)                         | **mp4 tag** 与mp4容器率略有差别                              |                                                              |                      | H5 MSE扩展                                                   |                                                              |
| **http-fmp4 ** (fmp4 over http)                 | **fmp4**                                                     |                                                              |                      | H5 MSE扩展                                                   |                                                              |
| **websocket-fmp4** (fmp4 over ws)               | **fmp4**                                                     |                                                              |                      | H5 MSE扩展                                                   |                                                              |
| **rtmp over quic（基于udp）**                   |                                                              |                                                              |                      |                                                              |                                                              |
| **http-flv over quic（基于udp）**               |                                                              |                                                              |                      |                                                              |                                                              |
| **hls over quic（基于udp）**                    | **m3u8+mpegts/TS部分**                                       |                                                              |                      |                                                              |                                                              |
| **kmp（私有协议，基于可靠udp协议kcp）**         |                                                              |                                                              |                      |                                                              | [媒体服务器NodeMediaServer](https://github.com/NodeMedia/NodeMediaServer) |
| **rtsp**                                        | **mpegts/TS**                                                |                                                              |                      |                                                              |                                                              |
| **rtcp**                                        |                                                              |                                                              |                      |                                                              |                                                              |
| **mpeg/mpg** 非流协议                           | **mpegts/TS**<br />**mpegps/PS**<br />                       |                                                              |                      |                                                              | ffmpeg **-f mpegts output.mpg**<br />ffmpeg **-f mpeg output.mpg** |
| --------------------                            |                                                              |                                                              |                      |                                                              |                                                              |
|                                                 | **mpegts/TS** (Transport Stream，传输流)，全称MPEG2-TS       | 视频编码: **mpeg1video**(MPEG-2 video)<br />音频编码:  mp1, mp2, mp3, 16-bit pcm_dvd, pcm_s16be, ac3 or dts |                      |                                                              | ffmpeg -f v4l2 -i /dev/video0 **-c:v** **mpeg1video** -f alsa -i default -c:a mp2 -r 24 **-f mpegts** mpgv1.mpg |
|                                                 |                                                              | 视频编码: **mpeg2video**(MPEG-2 video)<br />音频编码:  mp1, mp2, mp3, 16-bit pcm_dvd, pcm_s16be, ac3 or dts |                      |                                                              | ffmpeg -f v4l2 -i /dev/video0 -c:v **mpeg2video** -f alsa -i default -c:a mp2 -r 24 **-f mpegts** mpgv2.mpg <br />或者<br />ffmpeg -i ./mpgv1.mpg **-c:v mpeg2video** -c:a copy **-f mpegts** mpgv2.mpg |
|                                                 | **mpegps/PS**                                                |                                                              |                      |                                                              | ffmpeg **-f mpg output.mpg**                                 |
|                                                 | **mp4** (MPEG-4 part 14)                                     | h264视频编码+aac音频编码                                     |                      | H5 video标签原生支持                                         | 1. [mp4工具](https://share.weiyun.com/5pM2dwF)<br>2. [GPAC's github](https://github.com/gpac/gpac)<br>3. [mp4box.js's github](https://github.com/gpac/mp4box.js)<br />ffmpeg xxxx **-f mp4** |
|                                                 | **m4v **(raw MPEG-4 video)                                   |                                                              |                      |                                                              |                                                              |
|                                                 | **rtp**                                                      | _                                                            |                      |                                                              | 1. [rtptools版本库](https://gitee.com/yejinlei-mirror/rtptools) |
|                                                 | **flv** (Flash Video)                                        |                                                              |                      |                                                              | [flv工具](https://share.weiyun.com/5h54Qqp)                  |
|                                                 | **webm**                                                     | vp8视频编码+vorbis音频编码                                   |                      | H5 video标签原生支持                                         |                                                              |
|                                                 | **ogg**                                                      | theora视频编码+vorbis音频编码                                |                      | H5 video标签原生支持                                         |                                                              |
|                                                 | **hds** (HDS Muxer)                                          |                                                              |                      |                                                              |                                                              |
|                                                 | **mpegts** (MPEG2 Transport Stream)                          | mpeg1video/mpeg2video视频编码+mp1/2音频编码                  |                      |                                                              | 1. [雷霄骅 / TS_Analysis](https://gitee.com/leixiaohua/TS_Analysis)<br>2. [EasyICE](https://share.weiyun.com/LPCJIjEG) |
|                                                 | **aac** (Advanced Audio Coding)                              |                                                              |                      |                                                              | [雷霄骅 / AAC_Analysis](https://gitee.com/leixiaohua/AAC_Analysis) |
|                                                 | **adts** (ADTS AAC)                                          |                                                              |                      |                                                              |                                                              |
| --------------------                            |                                                              |                                                              |                      |                                                              |                                                              |
|                                                 |                                                              | **mpeg1video**(MPEG-2 video)<br />                           |                      |                                                              |                                                              |
|                                                 |                                                              | **mpeg2video**(MPEG-2 video)                                 |                      |                                                              |                                                              |
|                                                 |                                                              | **mpeg4**(MPEG-4 part 2)                                     |                      |                                                              |                                                              |
|                                                 |                                                              | **h264**(H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10)          |                      |                                                              | 1. [x264 github](https://code.videolan.org/videolan/x264.git)<br>2. [openh264 github](https://github.com/cisco/openh264)<br>3. [雷霄骅 / H264_Analysis](https://gitee.com/leixiaohua/H264_Analysis) |
|                                                 |                                                              | **avc1**(H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10)          |                      |                                                              |                                                              |
|                                                 |                                                              | **hevc**(H.265 / HEVC (High Efficiency Video Coding))        |                      |                                                              |                                                              |
|                                                 |                                                              | **av1**(Alliance for Open Media Video 1)                     |                      | H5 video标签原生支持(av1/mp4, 新版浏览器webrtc同样支持)      |                                                              |
|                                                 |                                                              | **theora**                                                   |                      |                                                              |                                                              |
|                                                 |                                                              | **vp8**(On2 VP8))                                            |                      |                                                              |                                                              |
|                                                 |                                                              | **vorbis**                                                   |                      |                                                              |                                                              |
|                                                 |                                                              | _                                                            | **qsv**              |                                                              |                                                              |

---

# 音视频编解码、容器及协议

## 一、编解码

## 二、音视频容器

| protocols（流协议及其他） | 容器格式                            | 支持的视频编码        | 支持的音频编码 | hwaccels（硬件加速） | 协议/文档            | 工具/库                                                      |
| ------------------------- | ----------------------------------- | --------------------- | -------------- | -------------------- | -------------------- | ------------------------------------------------------------ |
| _                         | **rtp**                             |                       | _              |                      |                      | 1. [rtptools版本库](https://gitee.com/yejinlei-mirror/rtptools) |
| **rtmp**                  | **flv** (Flash Video)               |                       |                |                      |                      | [flv工具](https://share.weiyun.com/5h54Qqp)                  |
| **http-flv**              | **flv** (Flash Video)               |                       |                |                      |                      |                                                              |
| **fmp4**                  | **mp4** (MPEG-4 part 14)            | h264                  | aac            |                      | H5 video标签原生支持 | 1. [mp4工具](https://share.weiyun.com/5pM2dwF)<br>2. [GPAC's github](https://github.com/gpac/gpac)<br>3. [mp4box.js's github](https://github.com/gpac/mp4box.js) |
| _                         | **m4v** (raw MPEG-4 video)          |                       |                |                      |                      |                                                              |
| _                         | **webm**                            | vp8                   | vorbis         |                      | H5 video标签原生支持 |                                                              |
| _                         | **ogg**                             | theora                | vorbis         |                      | H5 video标签原生支持 |                                                              |
| _                         | **hds** (HDS Muxer)                 |                       |                |                      |                      |                                                              |
| _                         | **mpegts **(MPEG2 Transport Stream) | mpeg1video/mpeg2video | mp1/2          |                      |                      | 1. [雷霄骅 / TS_Analysis](https://gitee.com/leixiaohua/TS_Analysis)<br>2. [EasyICE](https://share.weiyun.com/LPCJIjEG) |
| _                         | **aac** (Advanced Audio Coding)     |                       |                |                      |                      | [雷霄骅 / AAC_Analysis](https://gitee.com/leixiaohua/AAC_Analysis) |
| _                         | **adts** (ADTS AAC)                 |                       |                |                      |                      |                                                              |

### 2.1、MP4 容器

​	**mp4** 或称 **MPEG-4 Part 14**，是一种多媒体容器格式，扩展名为`.mp4`。

#### 2.2.1、MP4 标准

	mp4info ./1.mp4

#### 2.1.2、工具

| 工具                                                    | 在线 / 本地 | 备注                                   |
| ------------------------------------------------------- | ----------- | -------------------------------------- |
| [bento4](https://www.bento4.com/documentation/mp4info/) | 本地        | 包含mp4dump、mp4edit、mp4encrypt等工具 |
| [MP4Box](https://gpac.wp.imt.fr/)                       | 本地        |                                        |
| [mp4box.js](https://gpac.github.io/mp4box.js/)          | 在线        |                                        |

- bento4
  - mp4info
    - 
  - mp4dump

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