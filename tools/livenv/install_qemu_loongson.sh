git clone https://gitee.com/loongsonlab/qemu.git
cd ./qemu
mkdir ./build
cd ./build
../configure --target-list=mipsel-softmmu,mips64el-softmmu --disable-werror
make -j8