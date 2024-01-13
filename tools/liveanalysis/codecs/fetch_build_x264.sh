wget http://download.videolan.org/pub/videolan/x264/snapshots/x264-snapshot-20191217-2245.tar.bz2
tar jxvf x264-snapshot-20191217-2245.tar.bz2 & mkdir build & cd x264-snapshot-20191217-2245
./configure --enable-static --enable-shared --enable-debug --disable-asm --prefix=../build
make j4 & make install

