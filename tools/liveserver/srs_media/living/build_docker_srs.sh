sudo cp -f ../../ffmpeg/ffmpeg objs/ffmpeg
sudo cp -f ../../srs/trunk/objs/srs objs/srs
sudo docker build -t yejinlei/srsv2.0-r6_centos7.6:1.2 .
