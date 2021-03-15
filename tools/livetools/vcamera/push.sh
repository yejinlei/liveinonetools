ffmpeg -framerate 16 -video_size 1024x768 -f x11grab -i $DISPLAY -vf format=yuv420p -threads 0 -f v4l2 /dev/video0
