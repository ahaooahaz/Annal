all: env

env: OPT = all
env: _envbuild

FFmpeg: OPT = FFmpeg
FFmpeg: _envbuild

OpenCV: OPT = OpenCV
OpenCV: _envbuild

vscode: OPT = vscode
vscode: _envbuild
	
_envbuild:
	@bash utils/envbuild/linux/envbuild.sh --$(OPT)
	
.PHONY: all env FFmpeg OpenCV vscode
