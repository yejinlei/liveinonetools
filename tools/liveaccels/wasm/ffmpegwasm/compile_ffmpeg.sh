echo "Beginning Build:"
rm -r $(pwd)/dist
mkdir -p $(pwd)/dist
make clean
emconfigure ./configure --cc="emcc" --cxx="em++" --ar="emar" --prefix=$(pwd)/dist --enable-cross-compile --target-os=none --arch=x86_32 --cpu=generic \
	   	--enable-gpl --enable-version3 --disable-avdevice --disable-avformat --disable-swresample --disable-postproc --disable-avfilter \
	        --disable-programs --disable-logging --disable-everything --enable-decoder=hevc --enable-decoder=h264 \
		--disable-ffprobe --disable-ffserver --disable-doc --disable-devices --disable-network \
		--disable-hwaccels --disable-parsers --disable-bsfs --disable-debug --disable-protocols --disable-indevs --disable-outdevs --disable-asm \
make -j4
make install
cd ..

export TOTAL_MEMORY=268435456
export EXPORTED_FUNCTIONS="[ \
    		'_avcodec_register_all', \
	        '_avcodec_find_decoder', \
		'_avcodec_alloc_context3', \
		'_avcodec_open2', \
		'_av_free', \
		'_av_frame_alloc', \
		'_avcodec_close', \
		'_avcodec_decode_video2_js', \
		'_avcodec_get_image_width_js', \
		'_avcodec_get_image_height_js', \
		'_avcodec_get_chroma_format_js', \
		'_avcodec_get_image_plane_js', \
		'_avcodec_get_image_pitch_js', \
		'_avcodec_get_image_bit_depth_js', \
		'_avcodec_close_AVCodecContext_js', \
		'_avcodec_flush_buffers']"

emcc dist/lib/libavcodec.a dist/lib/libavutil.a \
	 -s OUTLINING_LIMIT=100000 \
	 -s TOTAL_MEMORY=${TOTAL_MEMORY} \
	 -s EXPORTED_FUNCTIONS="${EXPORTED_FUNCTIONS}" \
	 -O2 \
	 --pre-js libffmpeg_pre.js \
	 --post-js libffmpeg_post.js \
	 -o libffmpeg.js 
