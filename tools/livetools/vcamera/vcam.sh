git clone https://github.com/umlaeute/v4l2loopback.git && cd v4l2loopback
make && make install
depmod -a
modprobe v4l2loopback
