# liveinonetools
音视频工具集

## 工具集（tools）
### capture，截图工具
### srstools，srs相关
- clangtools，各种clang分析工具及脚本
- gperf，google性能分析工具及脚本
- valgrind，valgrind分析脚本
- build_srs.sh，srs构建脚本
- build_srs_gprof.sh，srs gprof编译脚本
- pull_srs.sh，拉取srs代码
- centos6.5_dev.sh，srs centos6.5的容器开发环境
### liveanalysis，多媒体协议分析
### livesniff, 多媒体嗅探、压测
- livesniff，多功能多媒体嗅探，压测工具
### liveaccels，多媒体加速
- wasm，各种webassembly播放器、服务器等
    - install_emsdk.sh，emsdk安装脚本
    - livegowasm，livego wasm形式服务器
### liveplayer，多媒体播放
- call_run_nginx.sh，nginx启动脚本，会拷贝player到/
- call_run_livego.sh，livego启动脚本
- pub_rtmp.sh，推送rtmp流
- pub_rtsp.sh，推送rtsp流
- player，播放器
### livevision，多媒体视觉
### liveserver，多媒体服务器
- srstools，srs直播工具相关工具
  - living
    - start_srs.sh，SRS Forward模式集群docker-compose编排，支持rtmp(11936)、HLS(11937)、HTTP-FLV(11938)、HDS(11939)
    - stop_srs.sh，停止脚本
  - pull_srs.sh，拉取srs版本库脚本
  - build_srs，编译srs release脚本，需放入srs版本库中
  - build_srs_gprof.sh，编译加入gprof性能工具
  - centos6.5_dev.sh，容器开发环境，已经集成了gcc、gdb等工具
  - gperf，google性能工具相关脚本
  - clangtools，Apple clang/llvm相关工具
  - valgrind，linux下algring内存、cpu分析工具
- run_livego，纯go实现的支持rtmp、hls、httpflv协议的轻量级直播服务器
- run_easydarwin，纯go实现的支持rtsp的直播服务器
- run_live555.sh，c++实现的流媒体项目
- ipc_sipserver，sip服务器