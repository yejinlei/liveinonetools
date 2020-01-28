apt install libtool libreadline-dev autoconf
git clone https://github.com/coreemu/core.git quagga 
cd quagga
./bootstrap.sh
./configure --disable-doc --enable-user=root --enable-group=root --with-cflags=-ggdb \
	    --sysconfdir=/usr/local/etc/quagga --enable-vtysh \
	        --localstatedir=/var/run/quagga
make
sudo make installdir core
