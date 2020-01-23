ffmpeg -f video4linux2 -i /dev/video0 -metadata mimetype=application/x-flv -r 25 -vcodec libx264 -f flv rtmp://localhost:1935/live/1

