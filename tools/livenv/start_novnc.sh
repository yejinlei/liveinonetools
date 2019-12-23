x11vnc -storepasswd
nohup x11vnc&  -forever -shared -rfbauth ~/.vnc/passwd
/usr/share/novnc/utils/launch.sh  --vnc localhost:5900
