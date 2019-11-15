apt-get build-dep vlc
git clone https://github.com/videolan/vlc.git
cd vlc
./bootstrap
./configure --disable-xcb --disable-qt --disable-ncurses --disable-vdpau
make -j8
make
