#!/bin/bash

picpath=${HOME}/Pictures/$(date +%s)-background.jpeg

wget https://source.unsplash.com/random/4096x2160 -O ${picpath} > /dev/null 2>&1

gsettings set org.gnome.desktop.background picture-uri "file:${picpath}"
# gsettings set org.gnome.desktop.background picture-options 'none'
# gsettings set org.gnome.desktop.background picture-options 'wallpaper'
# gsettings set org.gnome.desktop.background picture-options 'centered'
# gsettings set org.gnome.desktop.background picture-options 'scaled'
# gsettings set org.gnome.desktop.background picture-options 'stretched'
# gsettings set org.gnome.desktop.background picture-options 'zoom'
gsettings set org.gnome.desktop.background picture-options 'spanned'
