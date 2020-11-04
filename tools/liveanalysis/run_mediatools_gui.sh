docker stop mediatools_gui
docker rm mediatools_gui
docker run --name='mediatools_gui' --rm --privileged=true -d -p 6080:80 -p 5900:5900 -e VNC_PASSWORD=eye yejinlei/mediatools_gui:0.1
#firefox http://127.0.0.1:6080
