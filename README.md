# MaaS工具集
Media as a Service, from codec to server architecture...✍ 

![](doc/音视频视图.png)
---

### liveanalysis，多媒体、信令、消息等配套协议

- ❀多媒体及相关协议视图❀

  protocols（流协议及其他） | formats(封装容器) | codecs（音视频编解码） | hwaccels（硬件加速） | 协议/文档 | 工具/库
  ---|---|---|---|---|---
  **直播↷** | **----** | **---** | **---** | **---**
  **rtmp** | flv tag | | | | 1. [librtmp's github](https://github.com/yixia/librtmp/tree/master/librtmp)<br>2. [nginx-rtmp-module(支持h265)代码](https://github.com/adwpc/nginx-rtmp-module)
  **rtmps** | | | |
  **rtmpt** | | | |
  **rtmpts** | | | |
  **udt（可靠udp协议）** | | | |
  **srt（基于udt）** | | | | | 1. [srt's github](https://github.com/Haivision/srt)<br>2. [srt-live-server's github](https://github.com/Edward-Wu/srt-live-server)
  **hls**(Apple HTTP Live Streaming) | m3u8+mpegts部分 | | | 1. H5 video标签原生支持 or H5 MSE扩展<br>2. 协议 - [HTTP Live Streaming draft-pantos-http-live-streaming-12](https://tools.ietf.org/html/draft-pantos-http-live-streaming-12) |
  **dash**(Dynamic adaptive streaming over HTTP) | mpd+m4v/m4a\|webm\|3gp | | | | 1. [video.js dash插件](https://github.com/videojs/videojs-contrib-dash.git)<br>2. [dash测试页](http://reference.dashif.org/dash.js/nightly/samples/dash-if-reference-player/index.html)<br>3. [Bento4's github](https://github.com/axiomatic-systems/Bento4)<br>4. [Bento4's Bin](https://share.weiyun.com/5w4fPrC)<br>5. [DASHEncoder2's github](https://github.com/zhanghuicuc/DASHEncoder2/tree/dev-v2)<br>6. [DASH website demo's github（类似bitcodin网站）](https://github.com/zhanghuicuc/dash_website_demo)
  **http-flv** | | | | H5 MSE扩展
  **http-fmp4** | | | | H5 MSE扩展
  **websocket-flv** | | | | H5 MSE扩展
  **websocket-fmp4** | | | | H5 MSE扩展
  **rtmp over Qquic（基于udp）** | | | |
  **http-flv over quic（基于udp）** | | | |
  **hls over quic（基于udp）** | | | |
  **kmp（私有协议，基于可靠udp协议kcp）** | | | | | [媒体服务器NodeMediaServer](https://github.com/NodeMedia/NodeMediaServer)
  _ |**flv**(Flash Video) |    | | | [flv工具](https://share.weiyun.com/5h54Qqp)
  _ |**mp4**(MPEG-4 part 14) | h264视频编码+aac音频编码 | | H5 video标签原生支持 | 1. [mp4工具](https://share.weiyun.com/5pM2dwF)<br>2. [GPAC's github](https://github.com/gpac/gpac)<br>3. [mp4box.js's github](https://github.com/gpac/mp4box.js)
  _ |**fmp4(fragment mp4)** | h264视频编码+aac音频编码 | | | 
  _ |**m4v**(raw MPEG-4 video) | | | |
  _ |**webm** | vp8视频编码+vorbis音频编码 | | H5 video标签原生支持 |
  _ |**ogg** | theora视频编码+vorbis音频编码 | | H5 video标签原生支持 |
  _ |**hds**(HDS Muxer) |    | | |
  _ |**mpeg1video**(raw MPEG-1 video)
  _ |**mpeg2video**(raw MPEG-2 video)
  _ |**mpegts**(MPEG-2 Transport Stream) | | | | [雷霄骅 / TS_Analysis](https://gitee.com/leixiaohua/TS_Analysis)
  _ |**aac**(Advanced Audio Coding) | | | | [雷霄骅 / AAC_Analysis](https://gitee.com/leixiaohua/AAC_Analysis)
  _ |**adts**(ADTS AAC)
  _ |_ |**mpeg1video**(MPEG-1 video)    | | |
  _ |_ |**mpeg2video**(MPEG-2 video)    | | |
  _ |_ |**mpeg4**(MPEG-4 part 2)    | | |
  _ |_ |**h264**(H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10)    | | | 1. [x264 github](https://code.videolan.org/videolan/x264.git)<br>2. [openh264 github](https://github.com/cisco/openh264)<br>3. [雷霄骅 / H264_Analysis](https://gitee.com/leixiaohua/H264_Analysis)
  _ |_ |**avc1**(H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10)    | | |
  _ |_ |**hevc**(H.265 / HEVC (High Efficiency Video Coding))    | | | 
  _ |_ |**theora** | | |
  _ |_ |**vp8**(On2 VP8)) | | |
  _ |_ |**vorbis** | | |
  _ |_ |_|**qsv** | |
  
- ❀ONVIF❀
  
  _ | _ |  _
  ---|---|---

- ❀GB28181(音视频部分)❀

  _ | _ |  协议及参考实现
  ---|---|---
  **rtsp** | | 1. 协议-[Real Time Streaming Protocol (RTSP)](https://www.ietf.org/rfc/rfc2326.txt)，实时流传输协议<br>2. 参考实现：[rtspdump](https://bisqwit.iki.fi/source/ms-rtsp-dump/)
  **rtcp** | | 1. 协议 - [RTP: A Transport Protocol for Real-Time Applications](https://tools.ietf.org/html/rfc3550)，RTP控制协议
  **rtp**(rtp_mepg_ts) | | 1. 协议 - [ RTP Profile for Audio and Video Conferences with Minimal Control](https://tools.ietf.org/html/rfc3551)，实时传输协议<br>2. 参考实现： [jrtplib's github](https://github.com/j0r1/JRTPLIB)、[GoRTP's github](https://github.com/whtang/GoRTP)、[linephone ortp's github](https://github.com/dmonakhov/ortp)、[live555's github](https://github.com/cmberryau/live555) 
 
- ❀信令控制、IM消息协议❀
  
  协议 | _ | _
  ---|---|---
  **h.323** | | ![](tools/liveanalysis/doc/h323-stack.png)<br>1. [H.323 Version 6 - Overview](https://www.packetizer.com/ipmc/h323/whatsnew_v6.html)<br>2. [h323's wiki](https://wiki.wireshark.org/H323)<br>3. 参考实现：[h323plus's github](https://github.com/willamowius/h323plus)
  **sip/simple** | | 1. [SIP: Session Initiation Protocol](https://www.ietf.org/rfc/rfc3261.txt)<br>2. 参考实现：[开源sip server & sip client 和开发库 一览](https://www.cnblogs.com/lidabo/p/7587709.html)、[baresip's github](https://github.com/alfredh/baresip)、[librc's github](https://github.com/creytiv/re)
  **xmpp/jingle** | | 1. [rfc3921:Extensible Messaging and Presence Protocol (XMPP):Instant Messaging and Presence](https://tools.ietf.org/html/rfc3921), [CN](http://wiki.jabbercn.org/RFC3921)<br>3. 参考实现：[jxmpp's github](https://github.com/igniterealtime/jxmpp)
  **webrtc** | | [mozilla's webrtc api](https://developer.mozilla.org/zh-CN/docs/Web/API)

- ❀NAT/ICE/STUN/TURN（中转）❀

  NAT/ICE/STUN/TURN | 协议、标准 | 参考实现
  ---|---|---
  _ | [RFC 5389: STUN](https://tools.ietf.org/html/rfc5389)<br>[RFC 5780: NAT Behavior Discovery Using STUN](https://tools.ietf.org/html/rfc5780)<br>[RFC 7064: URI Scheme for STUN](https://tools.ietf.org/html/rfc7064)<br>[RFC 5766: TURN: Relay Extensions to STUN](https://tools.ietf.org/html/rfc5766)<br>[RFC 5245: ICE: A Protocol for NAT for Offer/Answer Protocols](https://tools.ietf.org/html/rfc5245)<br>[RFC 6062: TURN Extensions for TCP Allocations](https://tools.ietf.org/html/rfc6062)<br>[RFC 7065: TURN URI](https://tools.ietf.org/html/rfc7065)<br>[RFC 6544: TCP Candidates with ICE](https://tools.ietf.org/html/rfc6544)|

- ❀协议转换/网关❀

  协议 | 源码、项目 | 例子
  ---|---|---
  webrtc2sip | [github](https://github.com/DoubangoTelecom/webrtc2sip) | webrtc sip proxy
  OverSIP | [github](https://github.com/versatica/OverSIP) | sip over websocket 
  kurento-rtmp | [github](https://github.com/godka/kurento-rtmp) |
  rtmp-to-webrtc | [github](https://github.com/notedit/rtmp-to-webrtc) |
  kurento-rtsp2webrtc | [github](https://github.com/lulop-k/kurento-rtsp2webrtc) |
  RTSPtoWebRTC | [github](https://github.com/deepch/RTSPtoWebRTC) |
  media-server-go| [github](https://github.com/notedit/media-server-go) | webrtc媒体服务器，同时支持rtmp to webrtc、webrtc to rtmp、webrct to hls

- WebRTC技术方向，视频会议
  - **WebRTC技术**
  *WebRTC (Web Real-Time Communications) 是一项实时通讯技术，它允许网络应用或者站点，在不借助中间媒介的情况下，建立浏览器之间点对点（Peer-to-Peer）的连接，实现视频流和（或）音频流或者其他任意数据的传输。WebRTC包含的这些标准使用户在无需安装任何插件或者第三方的软件的情况下，创建点对点（Peer-to-Peer）的数据分享和电话会议成为可能。WebRTC包含了若干相互关联的API和协议以达到这个目标。*
  - **对等媒体之应用**
  - **数据通道之应用**
  - **[ICE/STUN/TURN/SDP依赖协议](https://developer.mozilla.org/zh-CN/docs/Web/API/WebRTC_API/Protocols)**
    - **ICE**，交互式连接设施Interactive Connectivity Establishment (ICE) 是一个允许你的浏览器和对端浏览器建立连接的协议框架。在实际的网络当中，有很多原因能导致简单的从A端到B端直连不能如愿完成。这需要绕过阻止建立连接的防火墙，给你的设备分配一个唯一可见的地址（通常情况下我们的大部分设备没有一个固定的公网地址），如果路由器不允许主机直连，还得通过一台服务器转发数据。ICE通过使用以下几种技术完成上述工作：
    - **STUN**，NAT的会话穿越功能Session Traversal Utilities for NAT (STUN) (缩略语的最后一个字母是NAT的首字母)是一个允许位于NAT后的客户端找出自己的公网地址，判断出路由器阻止直连的限制方法的协议。
    - **NAT**，网络地址转换协议Network Address Translation (NAT) 用来给你的（私网）设备映射一个公网的IP地址的协议。一般情况下，路由器的WAN口有一个公网IP，所有连接这个路由器LAN口的设备会分配一个私有网段的IP地址（例如192.168.1.3）。私网设备的IP被映射成路由器的公网IP和唯一的端口，通过这种方式不需要为每一个私网设备分配不同的公网IP，但是依然能被外网设备发现。一些路由器严格地限定了部分私网设备的对外连接。这种情况下，即使STUN服务器识别了该私网设备的公网IP和端口的映射，依然无法和这个私网设备建立连接。这种情况下就需要转向TURN协议。
    - **TURN**，一些路由器使用一种“对称型NAT”的NAT模型。这意味着路由器只接受和对端先前建立的连接（就是下一次请求建立新的连接映射）。NAT的中继穿越方式Traversal Using Relays around NAT (TURN) 通过TURN服务器中继所有数据的方式来绕过“对称型NAT”。你需要在TURN服务器上创建一个连接，然后告诉所有对端设备发包到服务器上，TURN服务器再把包转发给你。很显然这种方式是开销很大的，所以只有在没得选择的情况下采用。客户端通过给公网的STUN服务器发送请求获得自己的公网地址信息，以及是否能够被（穿过路由器）访问。
    - **SDP**，会话描述协议Session Description Protocol (SDP) 是一个描述多媒体连接内容的协议，例如分辨率，格式，编码，加密算法等。所以在数据传输时两端都能够理解彼此的数据。本质上，这些描述内容的元数据并不是媒体流本身。

- 实时通信技术RTC方向，视频会议及移动单兵
  - [Voice over IP protocol families](https://wiki.wireshark.org/VOIPProtocolFamily)
  - [Real-Time Communications Quick Start Guide](http://rtcquickstart.org/guide/multi/index.html)

- 互联网直播技术，视频监控
  - H5音视频
    - [各浏览器MSE情况](https://caniuse.com/#feat=mediasource)
    - [各浏览器对dash、hls、smooth支持情况](https://bitmovin.com/docs/player/articles/device-and-cross-browser-support)
    - [H5音视频兼容性测试页](http://ossrs.net/mse/html5.html)
    - [H5 MSE测试页](http://ossrs.net/mse/mse.html)

- GB28181视频互联，视频监控
    - [《GBT 28181-2011 安全防范视频监控联网系统信息传输、交换、控制技术要求》](https://share.weiyun.com/5FGvp0D)
    - [《GB/T 35114-2017 公共安全视频监控联网信息安全技术要求》](https://share.weiyun.com/59txAC8)
    - [《GB/T 25724 公共安全视频监控数字视音频编解码技术要求》](https://share.weiyun.com/5xVEjAl)
	- RFC 2030 简单网络时间协议（SNTP）第四版
	- RFC 2326 RTSP：实时流协议
	- RFC 3261 SIP: 会话初始协议
	- RFC 3263 会话初始协议(SIP)：定位SIP 服务器
	- RFC 3428 会话初始协议(SIP) 即时消息扩展
	- RFC 3550 实时传输协议
	- RFC 3725 会话初始协议(SIP)中第三方呼叫控制(3PCC)的当前最佳实现
	- RFC 4566 SDP: 会话描述协议
	- RFC 3219 TRIP（Telephony Routing over IP）IP 电话路由协议
	- ITU-T Rec. G.711-1988 数字传输系统的通用部分：音频的脉冲编码调制
	- ITU-T Rec. G.723.1-1996 数字传输系统的通用部分：5.3和6.3kbit的多媒体通信传输中的双速率语音编码器
	- ITU-T Rec. G.729-1996 数字传输系统的通用部分：利用共轭结构代数码激励线形预测的8kbit/s语音编码
	- ITU-T Rec. H.264-2005 H 系列：音视频和多媒体系统，音视频服务基础－活动视频编码：通用音视频服务的先进视频编码
	- ITU-T Rec. H.264.1-2005 H 系列：音视频和多媒体系统，音视频服务基础－活动视频编码：H.264先进视频编码的一致性说明
	- GB/T GB/T 25724-2010 安全防范监控数字视音频编解码技术要求

- ONVIF视频互联，视频监控
	- ONVIF 架构[ONVIF 架构]
	- ONVIF 分析服务WSDL [ONVIF 分析服务WSDL]
	- ONVIF 解析设备服务[ONVIF 解析设备服务WSDL]
	- ONVIF 设备服务WSDL [ONVIF 设备WSDL]
	- ONVIF 设备IO 服务WSDL [ONVIF 设备IO WSDL]
	- ONVIF 显示服务WSDL [ONVIF 显示WSDL]
	- ONVIF 事件服务WSDL [ONVIF 事件WSDL]
	- ONVIF 图像服务WSDL [ONVIF 图像WSDL]
	- ONVIF 媒体服务WSDL [ONVIF 媒体WSDL]
	- ONVIF PTZ 服务WSDL [ONVIF PTZ WSDL]
	- ONVIF 记录服务WSDL [ONVIF 记录WSDL]
	- ONVIF 远程发现WSDL [ONVIF 远程发现WSDL]
	- ONVIF 重放服务WSDL [ONVIF 重放WSDL]
	- ONVIF 检索服务WSDL [ONVIF 检索WSDL]
	- ONVIF 主题XML 命名空间[ONVIF 主题命名空间]

### livesniff， 多媒体嗅探、压测
- livesniff，多功能多媒体嗅探，压测工具
### liveaccels，多媒体加速
- wasm，各种webassembly播放器、服务器等
    - livegowasm，livego wasm形式服务器
- 硬件编码
  - NVENC
  - Quick Sync Video(介绍：http://www.elecfans.com/d/789787.html)
    - lsmod|grep 'i915'
    - vainfo
    - http://trac.ffmpeg.org/wiki/Hardware/QuickSync
  - MMX/SSE/SSE2/SSE3/SSE4
  - GPU/CUDA
### liveplayer，客户端（web、移动、pc）
- ❀播放器项目汇总❀

  播放器/库 | 源码及文档 | 分析及使用 
  --- | --- | ----
  **WEB端直播↷** | **---** | **---**
  video.js | 包含了大量扩展插件，[github](https://github.com/videojs)<br>1. [videojs/videojs-flash](https://github.com/videojs/videojs-flash)，支持rtmp<br>2. [videojs/http-streaming](https://github.com/videojs/http-streaming)，支持hls、dash等over http流协议 |
  mister-ben/videojs-flvjs | Video.js tech to use flv.js for FLV playback using MSE instead of Flash.[测试页](https://mister-ben.github.io/videojs-flvjs/)，[github](https://github.com/mister-ben/videojs-flvjs) |
  winlinvip/videojs-flow | [github](https://github.com/winlinvip/videojs-flow)，支持websocket-flv |
  playease.js | 支持http/websocket-flv、http/websocket-fmp4、dash、rtmp、hls、Original HTML5 Media Resources (eg. Ogg， Mpeg4， WebM)， [github](https://github.com/studease/playease/blob/master/test/index.html) |
  LivePlayer | [npmjs](https://www.npmjs.com/package/@liveqing/liveplayer)，[介绍文档](https://www.liveqing.com/docs/download/LivePlayer.html#%E7%9B%B8%E5%85%B3%E4%BB%8B%E7%BB%8D) |
  video-dev/hls.js | [测试页](https://hls-js.netlify.com/demo/)，[github](https://github.com/video-dev/hls.js/) |
  bilibili/flv.js | 开源h5 flv前端播放库，[github](https://github.com/bilibili/flv.js) | 
  jwplayer6 | flash播放器 |
  ckplayer | flash播放器 |
  CyberPlayer | 基于jwplayer二次开发的flash播放器，支持rtmp、flv、hls，[source](https://share.weiyun.com/5rVCKPD) |
  rx-player | 支持dash、smooth，[测试页](https://developers.canal-plus.com/rx-player/)，[github](https://github.com/canalplus/rx-player) |
  libde265 & libde265.js| 开源h265解码库，js库用到了wasm，[libde265.js's github](https://github.com/strukturag/libde265.js), [libde265's github](https://github.com/strukturag/libde265) |
  NodePlayer.js | 商业h265解码库，js库用到了wasm，[测试页](http://www.nodemedia.cn/products/node-media-player/) |
  **WebRTC生态↷**|**---** |**---**
  linephone | [官网](http://www.linphone.org/technical-corner/linphone?qt-technical_corner=0) | |
  jsSip | [官网](https://jssip.net/) | 支持sip(over websocket)和WebRTC的简易javascript库
  sipjs | [github](https://github.com/onsip/SIP.js), [exmaple](https://github.com/onsip/sipjs-examples) |
  PeerJS | [github](https://github.com/peers/peerjs) | PeerJS提供了一个完整可配置易于使用基于WebRTC的点对点的数据通信和媒体流のAPI

- call_run_nginx.sh，nginx启动脚本，会拷贝player到/
- call_run_livego.sh，livego启动脚本
- pub_rtmp.sh，推送rtmp流
- pub_rtsp.sh，推送rtsp流
- player，播放器
### livevision，多媒体视觉
- env_mxnet.sh，mxnet cpu环境
- env_mxnet_gpu.sh，mxnet gpu环境
- env_imageai_sh，imageai计算机视觉库环境
- env_imageai_gpu.sh，imageai gpu计算机视觉库环境
- obdetect，目标检测代码
### liveserver，多媒体服务器
- ❀服务器汇总❀

  项目 | 源码 |  分析
  --- | --- | --- |
  **RTMP直播服务器及组件↷** | **---** | **---**
  red5 | [github](https://github.com/Red5/red5-server) | java版，支持rtmp， rtmpt， rtmps， rtmpe
  srs | [github](https://github.com/ossrs/srs) | 开源直播多媒体服务器，支持rtmp、hls等协议
  nginx-rtmp-module | [github](https://github.com/arut/nginx-rtmp-module) | nginx的rtmp模块
  nginx-http-flv-module | [github](https://github.com/winshining/nginx-http-flv-module) | 支持http-flv/hls/dash
  nginx-rtmp-win32(windows版支持h265)| [github](https://github.com/illuspas/nginx-rtmp-win32) | [librtmp h265](https://blog.csdn.net/qq_33795447/article/details/89457581)
  im-pingo/nginx-rtmp-module | [github](https://github.com/im-pingo/nginx-rtmp-module) | 支持hls/http-ts/http-flv/rtmp/hls/dash/h265/h264/aac/mp3/notify/pull/push 
  im-pingo/CDN-Union_H265 | [github](https://github.com/im-pingo/CDN-Union_H265) |金山云FFmpeg扩展，支持h265
  **RTSP服务器及组件↷** |**---** |**---**
  live555 | [github](https://github.com/cmberryau/live555) | 开源rtsp多媒体服务器
  easydarwin | [github](https://github.com/EasyDarwin/EasyDarwin) | 开源rtsp多媒体服务器
  livego | [github](https://github.com/gwuhaolin/livego) | 开源rtsp多媒体服务器
  **WebRTC生态↷**|**---** |**---**
  / | ***Voip*** | /
  openmcu | [github](https://github.com/muggot/openmcu) | 全功能的H.323，SIP和RTSP的开源多点控制单元，具有以下特性：支持通过端口1420(http://localhost:1420)上的Web界面进行配置和控制;支持多国语言;支持的协议:H.323，SIP;支持的视频编解码器：H.261，H.263，H.263 +，H.264，VP8;支持的音频编解码器:G.711，G.722，G723.1，G.726，G.729，iLBC，Speex，SILK，OPUS;支持多个不同的会议可以同时进行，使用不同的“房间”;支持实时显示呼叫统计;支持从MCU发起呼叫到远程终端;支持会议翻译为网络流。
  asterisk | [github](https://github.com/asterisk/asterisk) |
  asterisk-freepbx | [github](https://github.com/Brownster/asterisk-freepbx) |
  freeswitch | [github](https://github.com/signalwire/freeswitch) |
  kamailio | [官网](https://www.kamailio.org) | 开源sip代理服务器<br>![](tools/liveserver/doc/kamailio-sip-ucp.png)
  oversip | [github](https://github.com/versatica/OverSIP) | OverSIP is a powerful and flexible SIP proxy & server by the authors of RFC 7118 (The WebSocket Protocol as a Transport for SIP)
  / | ***Signaling*** | /
  opensips | [官网](https://www.opensips.org/) | 开源sip代理服务器<br>![]()
  gnugk | [官网](https://www.gnugk.org) | 基于openh323开源voip服务器 
  / | ***Media Servers*** | /
  Medooze | [github](https://github.com/medooze/media-server) | ![](tools/liveserver/doc/Medooze.png)
  SimpleWebRTC | [github](https://github.com/simplewebrtc/SimpleWebRTC) | 
  EasyRTC | [github](https://github.com/priologic/easyrtc) |
  OpenWebRTC | [github](https://github.com/EricssonResearch/openwebrtc) |
  MediaSoup | [github](https://github.com/versatica/mediasoup) |
  ant media server | [github](https://github.com/ant-media) | 
  licode | [github](https://github.com/lynckia/licode) |
  jitsi | [github](https://github.com/jitsi) | Jitsi不仅是WebRTC媒体服务器，而且还有一个完整的平台。 Jitsi系列产品包括Jitsi Videobridge（媒体中继，SFU），Jitsi Meet（会议网络客户端），Jicofo（Jitsi Conference Focus），Jigasi（Jitsi Gateway to SIP）和Jitsi SIP Phone。 它还使用Jingle（XMPP）和功能齐全的Web界面实现自己的信令控制。
  kurento | [github](https://github.com/Kurento) | 基于WebRTC开源MCU框架，架构图<br>![架构图](tools/liveserver/doc/kurento.png)<br>Networked streaming protocols, including HTTP, RTP and WebRTC. Group communications (MCUs( Multipoint Conferencing Unit) and SFUs(Selective Forwarding Unit.) functionality) supporting bothmedia mixing and media routing/dispatching. Generic support for computational vision（计算机视觉） and augmented reality（现实增强） filters. Media storage supporting writing operations for WebM andMP4and playing in all formats supported by GStreamer. Automatic media transcodification between any of the codecs supported by GStreamer including VP8, H.264, H.263, AMR, OPUS, Speex, G.711, etc
  Kurento/OpenVidu | [github](https://github.com/OpenVidu/openvidu) | |
  kite | [github](https://github.com/webrtc/KITE) | webrtc测试引擎
  / | ***Gateways*** | /
  janus-gateway | [github](https://github.com/meetecho/janus-gateway) |
  / | ***NAT traversal*** | /
  coturn | [github](https://github.com/coturn/coturn) | 开源stun/turn穿透服务器
  restund | [github](https://github.com/otalk/restund) |
  turnserver | [github](https://github.com/pannzh/TurnServer) |
  reTurnServer | |
   / | ***Plugins and SDKs*** | /
  Pion | [github](https://github.com/pion/webrtc) | go实现的WebRTC API，包罗万象 
  Temasys | [github](https://github.com/Temasys) | 

- srstools，srs直播工具相关工具
  - doc
    - [分发RTMP流](https://github.com/ossrs/srs/wiki/v1_CN_DeliveryRTMP)
    - [分发HLS（m3u8）2.0](https://github.com/ossrs/srs/wiki/v2_CN_DeliveryHLS)、[3.0](https://github.com/ossrs/srs/wiki/v3_CN_DeliveryHLS)
    - [分发HLS（只有音频）](https://github.com/ossrs/srs/wiki/v2_CN_DeliveryHLS#hlsaudioonly)、[3.0](https://github.com/ossrs/srs/wiki/v3_CN_DeliveryHLS#hlsaudioonly)
    - [不中断配置重载](https://github.com/ossrs/srs/wiki/v1_CN_Reload)
    - [Gop缓存 2.0](https://github.com/ossrs/srs/wiki/v2_CN_LowLatency#gop-cache)、[3.0](https://github.com/ossrs/srs/wiki/v3_CN_LowLatency#gop-cache)
    - [Forward模式 2.0](https://github.com/ossrs/srs/wiki/v1_CN_Forward)、[3.0](https://github.com/ossrs/srs/wiki/v3_CN_Forward)
    - [对推到SRS中的RTMP进行转码（via FFmpeg）2.0](https://github.com/ossrs/srs/wiki/v2_CN_FFMPEG)、[3.0](https://github.com/ossrs/srs/wiki/v3_CN_FFMPEG)
    - [HTTP回调 2.0](https://github.com/ossrs/srs/wiki/v2_CN_HTTPCallback)、[3.0](https://github.com/ossrs/srs/wiki/v3_CN_HTTPCallback)
    - [带宽测试](https://github.com/ossrs/srs/wiki/v1_CN_BandwidthTestTool)
    - [HTTP API 2.0](https://github.com/ossrs/srs/wiki/v2_CN_HTTPApi)、[3.0](https://github.com/ossrs/srs/wiki/v3_CN_HTTPApi)
    - [采集并推流（via FFmpeg）](https://github.com/ossrs/srs/wiki/v1_CN_Ingest)
    - [DVR功能（SRS支持将RTMP流录制成flv文件）2.0](https://github.com/ossrs/srs/wiki/v2_CN_DVR)、[3.0](https://github.com/ossrs/srs/wiki/v3_CN_DVR)
    - [SRS日志功能（输出到控制台或文件）](https://github.com/ossrs/srs/wiki/v1_CN_SrsLog)
    - [内嵌HTTP服务器 2.0](https://github.com/ossrs/srs/wiki/v2_CN_SampleHTTP)、[3.0](https://github.com/ossrs/srs/wiki/v3_CN_SampleHTTP)
    - [SRS支持点播FLV流](https://github.com/ossrs/srs/wiki/v2_CN_FlvVodStream)
  - 源码分析
    - SRS边缘(edge)，所谓边缘edge服务器，就是边缘直播缓存服务器，配置时指定为remote模式和origin*（指定一个或多个源站IP），这个边缘edge服务器就是源站的缓存了
      - edge边缘节点，播放（SrsEdgeStatePlay标记）
      - edge边缘节点，从源站拉流和采集（SrsEdgeStateIngestConnected标记）
        - SrsEdgeUpstream::connect
      - edge边缘节点，推流到源站（SrsEdgeStatePublish标记）
        - SrsEdgeForwarder::start
 - 实战living目录
   - Forward模式
     - start_srs_compose.sh，SRS Forward模式集群docker-compose编排，支持rtmp(11936)、HLS(11937)、HTTP-FLV(11938)、HDS(11939)
     - start_srs_stack.sh，SRS Forward模式集群docker swarm/stack编排，支持rtmp(11936)、HLS(11937)、HTTP-FLV(11938)、HDS(11939)
     - start_srs_k8s.sh
    - edge模式
- WebRTC生态
   - [WebRTC-Developer-Tools-Landscape](https://bloggeek.me/download/11768/)
   - [Comparative study of WebRTC SFUs for Video Conferencing](https://uploads-ssl.webflow.com/5c6853c495409838d874a0d2/5cbed66cae2b88609106befa_IPTComm_2018_LoadTesting-12%5B23229%5D.pdf)
   - 媒体服务器
     - 媒体中转架构
     ![](tools/liveserver/doc/Mesh-MCU-SFU.png)
- pull_srs.sh，拉取srs版本库脚本
- build_srs，编译srs release脚本，需放入srs版本库中
- build_srs_gprof.sh，编译加入gprof性能工具
- centos6.5_dev.sh，容器开发环境，已经集成了gcc、gdb等工具
- gperf，google性能工具相关脚本
- clangtools，Apple clang/llvm相关工具
- valgrind，linux下valgring内存、cpu分析工具
- install_livego.sh，纯go实现的支持rtmp、hls、httpflv协议的轻量级直播服务器
- install_easydarwin，纯go实现的支持rtsp的直播服务器
- install_live555.sh，c++实现的流媒体项目
- ipc_sipserver，sip服务器

### livesafe，多媒体安全

- ❀全方位保护❀

  层级|方案|实战
  ---|---|---
  协议级|信息安全: 防盗链|
  协议级|信息安全: 软件授权机制|
  协议级|软件保护：定制协议|
  协议级|软件保护：嵌入水印|
  协议级|软件保护：隐写术|
  二进制级|软件保护：二进制乱序、混淆|
  二进制级|软件保护：二进制加密|
  二进制级|软件保护：防调试|
  二进制级|软件保护：防篡改|
  二进制级|软件保护：基于虚拟机运行时|
  代码级|软件保护：软件授权协议<br>*开源软件协议*![](tools/livesafe/doc/how-choice-license-1.png)<br>*如何选择license？*<br>![](tools/livesafe/doc/how-choice-license-2.png)<br>*开源许可证和一般著作权的区别*<br>![](tools/livesafe/doc/how-choice-license-3.png)|
  容器级|防止攻击：容器加固|
  操作系统级|防止攻击：防ddos攻击|
  硬件级|软件保护：加密狗和锁|
  硬件级|软件保护：可信启动|
  硬件级|软件保护：定制硬件|

- [开源软件协议介绍](https://gitee.com/oschina/git-osc/wikis/License)

### livenv， 环境安装
- k8环境及工具
  - install_minikube.sh
  - install_kubeadm.sh
  - install_kubectl.sh
  - install_minikube.sh
- web assembly工具链
  - install_emscripten.sh
  - install_emsdk.sh
- X11 GUI in Docker
  - start_docker_ubuntu-desktop.sh
- 国产CPU
  - 龙芯CPU
    - [龙芯开源社区-项目快速导航](http://www.loongnix.org/index.php/%E9%A1%B9%E7%9B%AE)
    - [龙芯全系 qemu 模拟器](https://gitee.com/loongsonlab/qemu)
    - install_qemu_loogson.sh, 龙芯全系qemu模拟器
- 国产操作系统
  - [uos操作系统](https://www.chinauos.com/applyfor)

### 测试地址
- rtmp，rtmp://202.69.69.180:443/webcast/bshdlive-pc
- rtmp，rtmp://1011.lssplay.aodianyun.com/demo/game
- rtmp，rtmp://play.nodemedia.io/live/ljq
- kmp，kmp://play.nodemedia.io/live/ljq
- http-flv(h265)，http://alplay.nodemedia.cn/live/cam_178.flv
- http-flv(h265)，http://alplay.nodemedia.cn/live/cam_180.flv
- http-flv(h265)，http://alplay.nodemedia.cn/live/cam_112.flv
- rtsp，rtsp://184.72.239.149/vod/mp4:BigBuckBunny_115k.mov

### 其他
- [wireshark's Captures](https://wiki.wireshark.org/SampleCaptures)