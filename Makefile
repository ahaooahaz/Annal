env:
	@bash tools/envbuild/linux/envbuild.sh -a

sdl2:
	

ffmpeg:	PREFIX = $(HOME)/.local/ffmpeg/ffmpeg-4.4
ffmpeg:
	@sudo apt-get install libx264-dev libx265-dev libsdl2-2.0 libsdl2-dev libsdl2-mixer-dev libsdl2-image-dev libsdl2-ttf-dev libsdl2-gfx-dev -y
	cd 3rdparty/FFmpeg ; ./configure --prefix=$(PREFIX) --enable-gpl --enable-nonfree --enable-libfdk-aac --enable-libx264 --enable-libx265 --disable-optimizations --enable-libspeex --enable-shared --enable-pthreads --enable-version3 --enable-hardcoded-tables --cc=gcc --host-cflags= --host-ldflags= --disable-x86asm --enable-ffplay --enable-ffprobe --enable-ffmpeg ; make -j$(shell cat /proc/cpuinfo| grep "processor"| wc -l)


.PHONY: env ffmpeg
