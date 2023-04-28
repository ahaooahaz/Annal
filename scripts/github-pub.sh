#!/bin/bash
# github-pub.sh - publish file to github repo.
# to publish image to github imagehosting.
set -e

username=
repo=
path=
token=
message=
filepath=
content=

# _UPLOADTEMPLATE = "https://api.github.com/repos/%s/contents/%s"
# _PROXYTEMPLATE  = `https://cdn.jsdelivr.net/gh/%s@%s/%s`
function show_URL() {
    echo "https://cdn.jsdelivr.net/gh/${username}/${repo}/${path}"
}


# curl put URL.
# $1 username
# $2 repo
# $3 path
# $4 token
# $5 message
# $6 content
function put() {
    curl --location --request PUT 'https://api.github.com/repos/'$1'/'$2'/contents/'$3 \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer '$4 \
    --data '{
    "message": "'$5'",
    "content": "'$6'"
    }' 2>/dev/null
}

function usage() {
cat << EOF
Usage: github-pub.sh -r <repo> -p <path> -t <token> -m <message> [filepath]
    -u <username>   github username, default use GITHUB_IMAGEHOSTING_USERNAME env
    -r <repo>       github repo name default use GITHUB_IMAGEHOSTING_REPO env
    -t <token>      github token, default use GITHUB_IMAGEHOSTING_TOKEN env
    -p <path>       github repo path, default use GITHUB_IMAGEHOSTING_PATH env
    -m <message>    commit message
    
EOF
}

# base64 file content.
function base64_content() {
    content=$(base64 $1)
}

function main() {
    while getopts 'r:p:m:t:f:' opt; do
        case $opt in
            r)
                repo="$OPTARG"
                ;;
            p)
                path="$OPTARG"
                ;;
            t)
                token="$OPTARG"
                ;;
            m)
                message="$OPTARG"
                ;;
            f)
                filepath="$OPTARG"
                ;;
            \?)
                usage
                exit 1
                ;;
        esac
    done
    
    if [ -z "${token}" ]; then
        token=${GITHUB_IMAGEHOSTING_TOKEN}
    fi
    if [ -z "${username}" ]; then
        username=${GITHUB_IMAGEHOSTING_USERNAME}
    fi
    if [ -z "${repo}" ]; then
        repo=${GITHUB_IMAGEHOSTING_REPO}
    fi
    if [ -z "${path}" ]; then
        path=${GITHUB_IMAGEHOSTING_PATH}
    fi

    if [ -z "$repo" ] || [ -z "$path" ] || [ -z "$message" ] || [ -z "$filepath" ]; then
        usage
        exit 1
    fi

    base64_content $filepath
    put $username $repo $path $token $message $content
    show_URL
}

main $@