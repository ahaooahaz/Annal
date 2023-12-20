# Annal

![icon](./icons/icon.svg)

![built status](https://github.com/ahaooahaz/Annal/actions/workflows/ci.yaml/badge.svg) ![codecov](https://codecov.io/gh/ahaooahaz/Annal/branch/master/graph/badge.svg) [![License](https://img.shields.io/github/license/ahaooahaz/Annal)](https://raw.githubusercontent.com/ahaooahaz/Annal/master/LICENSE)

Synchronize develop environment, reference samples, built-in utils and configuration to build develop environment quickly.

## plugins

- [ohmyzsh](https://github.com/ohmyzsh/ohmyzsh)
- [powerlevel10k](https://github.com/romkatv/powerlevel10k)
- [ohmytmux](https://github.com/gpakosz/.tmux)
- [zsh-syntax-highlighting](https://github.com/zsh-users/zsh-syntax-highlighting)
- [zsh-autosuggestions](https://github.com/zsh-users/zsh-autosuggestions)
- [rime-config](configs/rime)
- [kitty-config](configs/kitty)
- [docker-wechat](scripts/docker-wechat.sh)
- [docker-wechat.work](scripts/docker-chat.work.sh)

## built-in utils

### jt

`jumpto` in order to improve the efficiency of switching between multiple machines.

jump to remote machine by ssh and sshpass, suitable for scenarios where multiple jump to different machines in a safe environment, password and details will save in ${HOME}/.jtremote.local, make sure it will not leak.

### WebRTC Media Publish Client

Publish video stream to server by `WebRTC` protocol through `UDP` or `TCP`.

supported media server:
- [zlmediakit](https://github.com/ZLMediaKit/ZLMediaKit)

### [MIT LICENSE](LICENSE)

