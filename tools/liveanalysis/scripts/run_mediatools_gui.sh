docker run --rm --privileged=true -d -p 6080:80 -p 5900:5900 -v /var/run/docker.sock:/var/run/docker.sock -e VNC_PASSWORD=eye yejinlei/mediatools_gui:0.1
