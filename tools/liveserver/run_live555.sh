#git clone https://github.com/rgaufman/live555 live555_rep
cd live555_rep
./genMakefiles linux-64bit
make -j8
cp mediaServer/live555MediaServer ../
cd ..
./live555MediaServer
