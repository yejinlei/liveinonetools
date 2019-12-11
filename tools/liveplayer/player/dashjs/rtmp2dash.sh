ffmpeg -i "rtmp://127.0.0.1:11936/live/5021" -window_size 5  -min_seg_duration 30000 -single_file 0 -c copy -f dash 5021.mpd
