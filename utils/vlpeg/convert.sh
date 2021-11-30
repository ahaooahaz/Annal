#!/bin/bash

set -x

function Usage() {
cat << EOF
Usage: ${0##*/} [-h|--help]
    -s|--source     "source video path"
    -d|--dest       "dest video path"
    -bs|--batch_source  "video dir"
EOF
}


# $1 source video path $2 dest video path
function vlc_convert() {
    # vlc -I dummy -vvv $1 --sout="#transcode{vcodec=h264,vb=1024,ab=192,channels=2,deinterlace}:standard{access=file,mux=ts,dst=$2}" vlc://quit
    # vlc -I dummy -vvv $1 --sout="#transcode{vcodec=h264,vb=800,acodec=mpga,ab=128,channels=2,samplerate=44100,scodec=none}:std{access=file{no-overwrite},mux=ts,dst=$2}" vlc://quit
    # 上面两个命令都不可用，对比发现不指定vb数值才好使
    #vlc -I dummy -vvv $1 --sout="#transcode{vcodec=h264,acodec=mpga,ab=128,channels=2,samplerate=44100,scodec=none}:std{access=file{no-overwrite},mux=mp4,dst=$2}" vlc://quit
    vlc -I dummy -vvv $1 --sout="#transcode{vcodec=h264,ab=128,channels=2,samplerate=44100,scodec=none}:std{access=file{no-overwrite},mux=mp4,dst=$2}" vlc://quit
    
    
}

# $1 source mp4 $2 dest mkv
function ffmpeg_convert() {
    ffmpeg -i $1 $2 -y
}

# $1 source $2 dest
function convert() {
    tmp="tmp_convertsh"
    tmp_v_name=$(cat /proc/sys/kernel/random/uuid)
    mkdir -p ./${tmp}
    vlc_convert $1 ${tmp}/${tmp_v_name}.mp4
    if [ $? -ne 0 ]; then
        echo -e "warning vlc convert not complate ok!"
    fi
    ffmpeg_convert ${tmp}/${tmp_v_name}.mp4 $2
    rm ${tmp} -rf
}

# $1 dir path
function batch_convert() {
    if [ ! -d $1 ]; then
        echo -e "dir not exist!"
        exit 2
    fi

    ls $1 | while read video_file
    do
        source_video=$1/${video_file}
        dest_video=$1/${video_file%%.*}.mkv
        echo "source: ${source_video}, dest: ${dest_video}"
        convert ${source_video} ${dest_video}
    done
}

function main() {
    source=""
    dest=""
    while [[ $# -ne 0 ]]
    do
    k=$1
    case $k in
        -s|--source)
            source=$2
            shift
            shift
            ;;
        -d|--dest)
            dest=$2
            shift
            shift
            ;;
        -bs|--batch_source)
            dir=$2
            shift
            shift
            ;;
        *)
            Usage
            exit 1
            ;;
    esac
    done
    
    if [[ "${source}" == "" || "${dest}" == "" ]] && [[ "${dir}" == "" ]]; then
        echo "error"
        Usage
        exit -1
    fi
    
    if [[ "${dir}" != "" ]]; then
        batch_convert ${dir%/}
    else
        convert ${source} ${dest}
    fi
}

main $@
