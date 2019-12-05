ffmpeg -i rtsp://127.0.0.1:554/live/1.sdp -q 0 -f mpegts -codec:v mpeg1video -s 1366x768 http://127.0.0.1:8081/supersecret
