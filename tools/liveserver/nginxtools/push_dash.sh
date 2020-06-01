#!/bin/bash
for((;;)); do \
    ffmpeg -i 'rtsp'  \
    -vcodec copy -an \
    -f flv -y rtmp://127.0.0.1:11935/live/1; \
    sleep 1; \
  done
