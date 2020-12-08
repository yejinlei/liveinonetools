#git clone https://github.com/numberwolf/FFmpeg-QuQi-H265-FLV-RTMP.git
#apt install librtmp-dev libx265-dev libfdk-aac-dev
git clone https://gitee.com/yejinlei-mirror/FFmpeg-QuQi-H265-FLV-RTMP.git
cd FFmpeg-QuQi-H265-FLV-RTMP && ./configure --enable-static --enable-gpl --enable-nonfree --enable-libx264 --enable-libx265 --enable-libfdk-aac --enable-librtmp && make
