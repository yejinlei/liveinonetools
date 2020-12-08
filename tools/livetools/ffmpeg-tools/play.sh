./FFmpeg-QuQi-H265-FLV-RTMP/ffmpeg -f video4linux2 -i /dev/video0 -metadata mimetype=application/x-flv -r 25 -t 10 -vcodec libx265 -acodec aac -f flv -y ./test.flv
./FFmpeg-QuQi-H265-FLV-RTMP/ffplay ./test.flv
