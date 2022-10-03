ffmpeg -i "rtsp://admin:hxpti123@10.9.0.227:554/cam/realmonitor?channel=1&subtype=0" -q 2 -f mpegts -codec:v mpeg1video -g 1 -bf 0 http://127.0.0.1:8081/supersecret/live1
