# liveinonetools
音视频工具集

## 工具集（tools）
### capture，截图工具
### srstools，srs相关
- clangtools，各种clang分析工具及脚本
- gperf，google性能分析工具及脚本
- build_srs.sh，srs构建脚本
- pull_srs.sh，拉取srs代码
- centos6.5_dev.sh，srs centos6.5的容器开发环境
### liveanalysis，多媒体协议分析
### livesniff, 多媒体嗅探、压测
- livesniff，多功能多媒体嗅探，压测工具
### liveaccels，多媒体加速
- wasm，各种webassembly播放器、服务器等
### liveplayer，多媒体播放
- call_run_nginx.sh，nginx启动脚本，会拷贝player到/
- call_run_livego.sh，livego启动脚本
- pub_rtmp.sh，推送rtmp流
- pub_rtsp.sh，推送rtsp流
- player，播放器
### livevision，多媒体视觉
### liveserver，多媒体服务器
- run_livego，纯go实现的支持rtmp、hls、httpflv协议的轻量级直播服务器
- run_easydarwin，纯go实现的支持rtsp的直播服务器
- run_live555.sh，c++实现的流媒体项目
- ipc_sipserver，sip服务器