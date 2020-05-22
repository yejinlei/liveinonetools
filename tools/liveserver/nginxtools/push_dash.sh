#!/bin/bash
for((;;)); do \
    ffmpeg -i 'rtsp'  \
    -vcodec copy -an \
    -f flv -y rtmp://127.0.0.1:21935/dash/livestream; \
    sleep 1; \
  done
