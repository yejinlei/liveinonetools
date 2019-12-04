ffmpeg -f video4linux2 -i /dev/video0 -r 25 -vcodec libx264 -f flv rtmp://localhost:1935/live/1

