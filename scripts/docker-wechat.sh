#!/bin/bash

xhost +

docker run -d --name wechat --device /dev/snd --ipc=host \
    -v /tmp/.X11-unix:/tmp/.X11-unix \
    -v $HOME/.WeChatFiles:/WeChatFiles \
    -e DISPLAY=unix$DISPLAY \
    -e XMODIFIERS=@im=fcitx \
    -e QT_IM_MODULE=fcitx \
    -e GTK_IM_MODULE=fcitx \
    -e AUDIO_GID=`getent group audio | cut -d: -f3` \
    -e GID=`id -g` \
    -e UID=`id -u` \
    bestwu/wechat
