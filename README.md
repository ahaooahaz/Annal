# Annal

![built status](https://api.travis-ci.com/AHAOAHA/Annal.svg)

Synchronize develop environment, reference samples, utils and configuration to build develop environment quickly.

## plugins

* [ohmyzsh](https://github.com/ohmyzsh/ohmyzsh)
* [ohmytmux](https://github.com/gpakosz/.tmux)
* [zsh-syntax-highlighting](https://github.com/zsh-users/zsh-syntax-highlighting)
* [zsh-autosuggestions](https://github.com/zsh-users/zsh-autosuggestions)
* [powerlevel10k](https://github.com/romkatv/powerlevel10k)

## utils

### jt

autojump to remote machine by ssh and sshpass, suitable for scenarios where multiple jump to different machines in a safe environment, password and details will save in ${HOME}/.jtremote.local, make sure it will not leak.

usage:

``` shell
$ jt -h
usage: jt [OPTION] [PARAMS]
    e    jump to remote machine with ssh, default option.
    s    save remote machine details.
        -i|--ip         remote ip.
        -u|--user       remote user.
        -p|--password   remote password.
        -P|--port       remote sshd service binding port.
        -f|--focus      overwrite already exist detail.
    l   show exist detail ips.

    -h|--help show help.
```

### video

#### gen

generate video from single image, base on [hybridgroup/gocv](https://github.com/hybridgroup/gocv).

usage:

```shell
$ annal video gen -h
gen video from single image, only support image(.jpeg) to video(.avi)

Usage:
    annal video gen [flags]

Flags:
        --N uint           count of video frames (default 1500)
    -f, --fps int          video fps (default 25)
        --height int       video height (default 1080)
    -h, --help             help for gen
    -i, --images strings   source image
    -o, --output string    output video path (default "annal.avi")
        --width int        video width (default 1920)
```

*video tools require gocv installed, if without gocv then without video tools*.

## [MIT LICENSE](LICENSE)
