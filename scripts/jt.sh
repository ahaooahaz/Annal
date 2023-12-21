#!/bin/bash

set -x

old_details=${HOME}/.config/jt.csv
details=${HOME}/.config/annal/jt/jt.csv
gauth_config=${HOME}/.config/gauth.csv

function usage() {
cat << USAGE
usage: jt [OPTION] [PARAMS]
        e    jump to remote machine with ssh, default option.
        s    save remote machine details.
            -i|--ip         remote ip.
            -u|--user       remote user.
            -p|--password   remote password.
            -P|--port       remote sshd service binding port.
            --2FA           2FA secret, base on gauth.
            --2FA-tag       2FA secret tag.
            -f|--focus      overwrite already exist detail.
        l    show exist detail ips.

        -h|--help   show help.
USAGE
}

function alert() {
    echo -e "\033[31m$1\033[0m"
}

function warn() {
    echo -e "\033[33m$1\033[0m"
}

function info() {
    echo -e "\033[32m$1\033[0m"
}

function s() {
    while [ $# -ne 0 ]
    do
        key=$1
        case ${key} in
            -i|--ip)
                ip=$2
                shift
                shift
                ;;
            -u|--user)
                user=$2
                shift
                shift
                ;;
            -p|--password)
                password=$2
                shift
                shift
                ;;
            -P|--port)
                port=$2
                shift
                shift
                ;;
            -f|--focus)
                focus=true
                shift
                ;;
            --2FA)
                fa2_secret=$2
                shift
                shift
                ;;
            --2FA-tag)
                fa2_tag=$2
                shift
                shift
                ;;
            *)
                usage
                return 1
        esac
    done

    if [ -z ${ip} ]; then
        echo -n "ip: "
        read ip
    else
        echo "ip: ${ip}"
    fi

    if [ -z ${user} ]; then
        echo -n "user: "
        read user
    else
        echo "user: ${user}"
    fi

    if [ -z ${port} ]; then
        echo -n "port: "
        read port
    else
        echo "port: ${port}"
    fi

    if [ -z ${password} ]; then
        echo -n "password: "
        read password
    else
        echo "password: ${password}"
    fi
    
    if [ -z ${fa2_tag} ]; then
        if [ $(type gauth >/dev/null 2>&1; echo $?) -ne 0 ]; then
            alert "2FA install gauth first"
            return
        fi
        echo -n "2FA tag: "
        read fa2_tag
    else
        echo "2FA tag: ${fa2_tag}"
    fi

    if [ -z ${fa2_secret} ]; then
        if [ $(type gauth >/dev/null 2>&1; echo $?) -ne 0 ]; then
            alert "2FA install gauth first"
            return
        fi
        echo -n "2FA secret: "
        read fa2_secret
    else
        echo "2FA secret: ${fa2_secret}"
    fi

    if [ ! -z "$(grep -E "^[^:]*{1}:${ip}:[^:]*{1}:[0-9]{1,5}{1}.*$" ${details} 2>/dev/null)" -a -z "${focus}" ]; then
        warn "ip already exist"
        return
    fi
    
    if [ -z ${fa2_tag} ]; then
        if [ $(sshpass -p ${password} ssh ${user}@${ip} -p ${port} -q -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null 'exit' 2>/dev/null; echo $?) -ne 0 ]; then
            alert "ssh check remote machine failed"
            return
        fi
    fi

    if [ ! -f ${details} ]; then
        mkdir -p $(dirname ${details})
    fi

    crypted=$(base64 <<< ${password})
    echo "${user}:${ip}:${crypted}:${port}:${fa2_tag}" >> ${details}
    if [ -z ${fa2_secret} -a -z ${fa2_tag} ]; then
        fa2 ${fa2_tag} ${fa2_secret}
    fi
    info "OK"
}

function e() {
    if [ -z "$1" ]; then
        alert "you need input remote machine ip suffix"
        usage
        return
    fi
    
    if [ ! -f "${details}" ]; then
        alert "you need save remote machine details first, use 'jt s'"
        usage
        return
    fi
    detail=$(tac ${details} 2>/dev/null | grep -E "^[^:]*{1}:[^:]*${1}{1}:[^:]*{1}:[0-9]{1,5}{1}.*$" 2>/dev/null)
    if [ "${detail}" == "" ]; then
        warn "${1} not match"
        return
    fi

    read user ip crypted port fa2_tag <<< $(echo ${detail} | awk -F ':' '{print $1,$2,$3,$4,$5}' 2>/dev/null)
    password=$(base64 -d <<< ${crypted})
    if [ -z ${fa2_tag} ]; then
        exec sshpass -p ${password} ssh ${user}@${ip} -p ${port} -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null
    else
        fa2_auth=$(gauth | grep ${fa2_tag} 2>/dev/null | awk -F ' ' '{print $2}')
        expect -c "
            log_user 0
            set timeout 60
            spawn ssh ${user}@${ip} -p ${port} -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null
            log_user 1
            expect {
                -re \".*(p|P)assword:\" {send \"${password}\r\";exp_continue}
                -re \".*FA auth]:\" {send \"${fa2_auth}\r\"}
                -re \".*>\" {}
            }
            set timeout -1
            interact
        "
    exit 0
    fi

    alert "not match details. insert remote machine into ${details}"
}

function l() {
    while read line
    do
        info $(echo ${line} | awk -F ':' '{print $2}')
    done < ${details}
}

function upgrade() {
    # move config.
    if [ -f ${old_details} ]; then
        mv ${old_details} ${details}
    fi
}

# $1: tag $2: secret
function fa2() {
    tag=$1
    secret=$2
    if [ -z $tag -o -z $secret ]; then
        alert "tag or secret invalid"
        return
    fi

    if [ $(grep ${tag} ${gauth_config} >/dev/null 2>&1; echo $?) -eq 0 ]; then
        return
    fi

    echo -n "$tag:$secret" >> ${gauth_config}
}

function main() {
    upgrade
    if [ $# -eq 0 ]; then
        usage
        exit 0
    fi

    case $1 in
        e)
            $@
            ;;
        s)
            $@
            ;;
        l)
            l
            ;;
        -h|--help)
            usage
            ;;
        *)
            e $@
    esac
}


main $@
