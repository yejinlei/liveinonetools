cd vlc
emconfigure ./configure --cc="emcc" --cxx="em++" --ar="emar" --disable-xcb --disable-qt --disable-ncurses --disable-vdpau 
make -j8
