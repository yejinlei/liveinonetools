sed 's/[s|S]rs/gcj/g' < srs > srs1
sed 's/SRS/GCJ/g' < srs1 > srs2
sed 's/ZhouGuowen/          /g' < srs2 > srs3
sed 's/winlin,wenjie.zhao/                  /g' < srs3 > srs4
sed 's/http:\/\/osgcj.net/                /g' < srs4 > srs5
sed 's/https:\/\/github.com\/osgcj\/gcj/                            /g' < srs5 > srs6
sed 's/Email: winlin@vip.126.com/                         /g' < srs6 > srs7
sed 's/\--with-hls --with-hds --with-dvr --without-nginx --with-ssl --without-ffmpeg --with-transcode --with-ingest --with-stat --with-http-callback --with-http-server --with-stream-caster --with-http-api --with-librtmp --with-research --without-utest --without-gperf --without-gmc --without-gmp --without-gcp --without-gprof --without-arm-ubuntu12 --without-mips-ubuntu12 --log-trace/                                                                                                                                                                                                                                                                                                                                                                                        /g' < srs7 > srs8
mv srs8 gcj
rm srs*
chmod +x gcj
