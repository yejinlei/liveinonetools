#!/bin/sh

MIN_PORT=30000;
MAX_PORT=30050;
docker run -p  3000:3000 -p $MIN_PORT-$MAX_PORT:$MIN_PORT-$MAX_PORT/udp -p 3001:3001  -p 8080:8080 -e "MIN_PORT=$MIN_PORT" -e "MAX_PORT=$MAX_PORT" -e "PUBLIC_IP=192.168.1.9" -v /home/yee/src/liveinonetools/tools/liveserver/webrtc/scheme/licode/licode_config.js:/opt/licode/licode_config.js lynckia/licode
