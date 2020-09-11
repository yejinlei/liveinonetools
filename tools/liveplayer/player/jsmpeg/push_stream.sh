ffmpeg -i "rtsp://" -q 2 -f mpegts -codec:v mpeg1video -g 1 -bf 0 http://127.0.0.1:8081/supersecret/live1
