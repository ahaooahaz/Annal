#!/bin/bash

set -e

omz=false
go=false
vim=true
tmux=false
vscode=false
xterm=false
gdb=false
FFmpeg=false
OpenCV=false

SHELL_FOLDER=$(cd "$(dirname "$0")";pwd)

FFMPEG_PATH=${SHELL_FOLDER}/../../../3rdparty/FFmpeg
OPENCV_PATH=${SHELL_FOLDER}/../../../3rdparty/OpenCV
OPENCV_CONTRIB_PATH=${SHELL_FOLDER}/../../../3rdparty/OpenCVContrib

OMZ_CONFIG_PATH=${SHELL_FOLDER}/../../../appc/omz
VIM_CONFIG_PATH=${SHELL_FOLDER}/../../../appc/vim
GDB_CONFIG_PATH=${SHELL_FOLDER}/../../../appc/gdbinit
TMUX_CONFIG_PATH=${SHELL_FOLDER}/../../../appc/tmux
XTERM_CONFIG_PATH=${SHELL_FOLDER}/../../../appc/xterm


function Usage() {
cat << EOF
Usage: ${0##*/} [-h|--help]
    -a|--all            "install all and all config"
    -go|--golang        "install golang env"
    --vim               "install vim and personal config"
    -omz|--oh_my_zsh    "install oh my zsh and personal config"
    --vscode            "install vscode and global config"
    --tmux              "install tmux and personal config"
    --FFmpeg            "install ffmpeg"
    --OpenCV            "install OpenCV"
EOF
}

function log() {
    case $1 in
        info)
        shift
        echo -e "\033[32m$@\033[0m"
        ;;
        warn)
        shift
        echo -e "\033[33m$@\033[0m"
        ;;
        error)
        shift
        echo -e "\033[31m$@\033[0m"
        ;;
        *)
        echo -e "\033[32m$@\033[0m"
        ;;
    esac
}
# $1 process schedule [0,100] 100 is must.
function probe() {
    declare -i progress=0;
    printf "[="
    while [ $progress -lt 100 ];do
    printf "="
    sleep 0.02
    done
    printf "=]100%"
}

# $1 command
function install() {
    if [[ $(which $1 >/dev/null 2>&1; echo $?) -ne 0 ]]; then
        sudo apt install $1 -y
    fi
}

function load_tmux() {
    tmux kill-server
}

function load() {
    if [[ ${tmux} ]]; then
        load_tmux
    fi
}

function check_system() {
    if [[ "$(lsb_release -d | awk -F ' ' '{print $2}')" != "Ubuntu" ]]; then
        log error "error: only support ubuntu."
        exit 1
    fi
}

function install_omz {
    log info  "install oh my zsh ..."
    install zsh
    if [[ ! -d $HOME/.oh-my-zsh ]]; then
        log warn "oh my zsh not exist will install it."
        log info "install oh my zsh ..."
        sh -c "$(curl -fsSL https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"
        log info "install oh my zsh done."
    fi
    
    log info "cp personal config ..."
    cp -r ${OMZ_CONFIG_PATH}/custom/plugins/* ${HOME}/.oh-my-zsh/custom/plugins > /dev/null 2>&1 || true
    cp -r ${OMZ_CONFIG_PATH}/custom/themes/* ${HOME}/.oh-my-zsh/custom/themes > /dev/null 2>&1 || true
    cp ${OMZ_CONFIG_PATH}/.zshrc ${HOME} || true
    log info "cp personal config done."
    log info  "install oh my zsh done."
    
}

function install_go {
    pkg="go.1.17.3.install.tar.gz"
    log info "install go env ..."
    if [[ $(go version > /dev/null 2>&1; echo $?) -ne 0 ]]; then
        wget https://golang.google.cn/dl/go1.17.3.linux-amd64.tar.gz -O ${pkg}
        tar -C $HOME/.local -xzf ${pkg}
    fi
    log info "install go env done."
}

function install_vim {
    log info  "install vim and personal config ..."
    install vim

    cp ${VIM_CONFIG_PATH}/.vimrc $HOME
    log info  "install vim and personal config done"
}

function install_tmux {
    log info "install tmux and personal config ..."
    install tmux

    cp ${TMUX_CONFIG_PATH}/.tmux.conf $HOME
    log info "install tmux and personal config done."

}

function install_vscode {
    log warn "install vscode and personal config TODO ..."
}

function install_xterm {
    log info "install xterm and personal config ..."
    install xterm

    cp ${XTERM_CONFIG_PATH}/.Xresources $HOME
    xrdb $HOME/.Xresources
    log info "install xterm and personal config done."
}

function install_gdb {
    log info  "install gdb and personal config ..."
    install gdb

    cp ${GDB_CONFIG_PATH}/.gdbinit $HOME
    log info  "install gdb and personal config done."
}

function install_FFmpeg {
    PREFIX=$HOME/.local/FFmpeg/FFmpeg-n4.4
    log info "install FFmpeg to ${PREFIX} ..."
    if [[ $(ffmpeg -version > /dev/null 2>&1; echo $?) -ne 0 ]]; then
        set -x
        cd ${FFMPEG_PATH}
        sudo apt-get install libx264-dev libx265-dev libsdl2-2.0 libsdl2-dev libsdl2-mixer-dev libsdl2-image-dev libsdl2-ttf-dev libsdl2-gfx-dev -y
        ./configure --prefix=${PREFIX} --enable-gpl --enable-nonfree --enable-libfdk-aac --enable-libx264 --enable-libx265 --disable-optimizations --enable-libspeex --enable-shared --enable-pthreads --enable-version3 --enable-hardcoded-tables --cc=gcc --host-cflags= --host-ldflags= --disable-x86asm --enable-ffplay --enable-ffprobe --enable-ffmpeg
        make -j$(cat /proc/cpuinfo| grep "processor"| wc -l)
        make install
        cd -

    fi
    log info "install FFmpeg done."
}

function install_OpenCV() {
    PREFIX=$HOME/.local/OpenCV/OpenCV-4.5.4
    log info "install OpenCV to ${PREFIX} ..."
    sudo apt-get install libgtk2.0-dev libjpeg-dev libpng-dev libtiff-dev libva-dev -y
    mkdir -p ${OPENCV_PATH}/build
    cd ${OPENCV_PATH}/build
    cmake -D CMAKE_BUILD_TYPE=RELEASE \
        -D CMAKE_INSTALL_PREFIX=${PREFIX} \
        -D OPENCV_EXTRA_MODULES_PATH=${OPENCV_CONTRIB_PATH}/modules \
        -D BUILD_TIFF=ON \
        -D WITH_FFMPEG=ON \
        -D WITH_TBB=ON \
        -D BUILD_TBB=ON \
        -D WITH_EIGEN=ON \
        -D WITH_V4L=ON \
        -D WITH_LIBV4L=ON \
        -D WITH_VTK=OFF \
        -D WITH_QT=OFF \
        -D WITH_OPENGL=ON \
        -D WITH_CUDA=ON \
        -D OPENCV_ENABLE_NONFREE=ON \
        -D INSTALL_C_EXAMPLES=ON \
        -D INSTALL_PYTHON_EXAMPLES=ON \
        -D BUILD_NEW_PYTHON_SUPPORT=ON \
        -D OPENCV_GENERATE_PKGCONFIG=ON \
        -D BUILD_TESTS=ON \
        -D ENABLE_FAST_MATH=ON \
        -D CUDA_FAST_MATH=ON \
        -D WITH_CUBLAS=ON \
        -D WITH_CUDNN=OFF \ # if CUDNN on should set behind 2 options.
        -D CUDNN_LIBRARY= \
        -D CUDNN_INCLUDE_DIR=/usr/local/cuda/include \
        -D BUILD_EXAMPLES=OFF ..
    make -j$(cat /proc/cpuinfo| grep "processor"| wc -l)
    make install
    cd -
    log info "install OpenCV done."
}

function manager() {
    if ${omz}; then
        install_omz
    fi

    if ${go}; then
        install_go
    fi

    if ${vim}; then
        install_vim
    fi

    if ${tmux}; then
        install_tmux
    fi

    if ${vscode}; then
        install_vscode
    fi

    if ${xterm}; then
        install_xterm
    fi

    if ${gdb}; then
        install_gdb
    fi

    if ${FFmpeg}; then
        install_FFmpeg
    fi

    if ${OpenCV}; then
        install_OpenCV
    fi
}



function main() {
    if [[ $# -eq 0 ]]; then
        Usage
        exit 0
    fi
    while [[ $# -ne 0 ]]
    do
    arg=$1
    case $arg in
        -a|--all)
        omz=true
        go=false
        vim=true
        tmux=true
        vscode=true
        xterm=true
        gdb=true
        FFmpeg=false
        OpenCV=false
        shift
        ;;
        -omz|--oh_my_zsh)
        omz=true
        shift
        ;;
        -go|--golang)
        go=true
        shift
        ;;
        --vim)
        vim=true
        shift
        ;;
        --tmux)
        tmux=true
        shift
        ;;
        --vscode)
        vscode=true
        shift
        ;;
        --xterm)
        xterm=true
        shift
        ;;
        --gdb)
        gdb=true
        shift
        ;;
        --FFmpeg)
        FFmpeg=true
        shift
        ;;
        --OpenCV)
        OpenCV=true
        shift
        ;;
        *)
        Usage
        exit 1
        ;;
    esac
    done
    check_system
    manager
    load
}

main $@
