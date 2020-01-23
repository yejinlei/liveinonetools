apt-get install pkg-config autoconf automake libtool
apt-get build-dep vlc

apt-get install qtdeclarative5-dev

apt-get install lua5.1
apt-get install libxcb-shm0-dev
apt-get install libavcodec-dev
apt-get install libavformat-dev
apt-get install libswscale-dev
apt-get install libpostproc-dev
apt-get install liba52-dev
apt-get install libxcb-xv0-dev

git clone https://github.com/videolan/vlc.git
cd vlc
./bootstrap
./configure --enable-x11 --enable-xvideo --enable-sdl --enable-avcodec --enable-avformat \
       	--enable-swscale --enable-mad --enable-libdvbpsi --enable-a52 --enable-libmpeg2 --enable-dvdnav \
       	--enable-faad --enable-vorbis --enable-ogg --enable-theora --enable-faac --enable-mkv --enable-freetype \
       	--enable-fribidi --enable-speex --enable-flac --enable-live555 --with-live555-tree=/usr/lib/live \
       	--enable-caca --enable-skins --enable-skins2 --enable-alsa --enable-qt4 --enable-ncurses
make -j8
make
