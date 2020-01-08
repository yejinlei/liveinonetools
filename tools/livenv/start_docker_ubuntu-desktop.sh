#docker run -d -p 6080:80 -p 5900:5900 -e VNC_PASSWORD=eye -v $(pwd)/sources.list:/etc/apt/sources.list:ro yejinlei/ubuntu-desktop-lxde-vnc
docker run -d -p 6080:80 -p 5900:5900 -e VNC_PASSWORD=eye yejinlei/ubuntu-desktop-lxde-vnc
firefox  http://127.0.0.1:6080
