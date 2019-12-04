ffmpeg -f video4linux2 -i /dev/video0 -qscale 10 -r 25 -s 640x480 -vcodec libx264 -f rtsp rtsp://127.0.0.1:554/live/1.sdp
