# If you come from bash you might have to change your $PATH.
# export PATH=$HOME/bin:/usr/local/bin:$PATH

# Path to your oh-my-zsh installation.
export ZSH="$HOME/.oh-my-zsh"

# Set name of the theme to load --- if set to "random", it will
# load a random theme each time oh-my-zsh is loaded, in which case,
# to know which specific one was loaded, run: echo $RANDOM_THEME
# See https://github.com/ohmyzsh/ohmyzsh/wiki/Themes
# ZSH_THEME="robbyrussell"
#ZSH_THEME="fino"
ZSH_THEME="custom"


# Set list of themes to pick from when loading at random
# Setting this variable when ZSH_THEME=random will cause zsh to load
# a theme from this variable instead of looking in $ZSH/themes/
# If set to an empty array, this variable will have no effect.
# ZSH_THEME_RANDOM_CANDIDATES=( "robbyrussell" "agnoster" )

# Uncomment the following line to use case-sensitive completion.
# CASE_SENSITIVE="true"

# Uncomment the following line to use hyphen-insensitive completion.
# Case-sensitive completion must be off. _ and - will be interchangeable.
# HYPHEN_INSENSITIVE="true"

# Uncomment the following line to disable bi-weekly auto-update checks.
# DISABLE_AUTO_UPDATE="true"

# Uncomment the following line to automatically update without prompting.
# DISABLE_UPDATE_PROMPT="true"

# Uncomment the following line to change how often to auto-update (in days).
# export UPDATE_ZSH_DAYS=13

# Uncomment the following line if pasting URLs and other text is messed up.
# DISABLE_MAGIC_FUNCTIONS="true"

# Uncomment the following line to disable colors in ls.
# DISABLE_LS_COLORS="true"

# Uncomment the following line to disable auto-setting terminal title.
# DISABLE_AUTO_TITLE="true"

# Uncomment the following line to enable command auto-correction.
# ENABLE_CORRECTION="true"

# Uncomment the following line to display red dots whilst waiting for completion.
# Caution: this setting can cause issues with multiline prompts (zsh 5.7.1 and newer seem to work)
# See https://github.com/ohmyzsh/ohmyzsh/issues/5765
# COMPLETION_WAITING_DOTS="true"

# Uncomment the following line if you want to disable marking untracked files
# under VCS as dirty. This makes repository status check for large repositories
# much, much faster.
# DISABLE_UNTRACKED_FILES_DIRTY="true"

# Uncomment the following line if you want to change the command execution time
# stamp shown in the history command output.
# You can set one of the optional three formats:
# "mm/dd/yyyy"|"dd.mm.yyyy"|"yyyy-mm-dd"
# or set a custom format using the strftime function format specifications,
# see 'man strftime' for details.
# HIST_STAMPS="mm/dd/yyyy"

# Would you like to use another custom folder than $ZSH/custom?
# ZSH_CUSTOM=/path/to/new-custom-folder

# Which plugins would you like to load?
# Standard plugins can be found in $ZSH/plugins/
# Custom plugins may be added to $ZSH_CUSTOM/plugins/
# Example format: plugins=(rails git textmate ruby lighthouse)
# Add wisely, as too many plugins slow down shell startup.
plugins=(
        git
        z
#zsh-autosuggestions
        zsh-syntax-highlighting
        autojump
        )

source $ZSH/oh-my-zsh.sh

# User configuration

# export MANPATH="/usr/local/man:$MANPATH"

# You may need to manually set your language environment
# export LANG=en_US.UTF-8

# Preferred editor for local and remote sessions
# if [[ -n $SSH_CONNECTION ]]; then
#   export EDITOR='vim'
# else
#   export EDITOR='mvim'
# fi

# Compilation flags
# export ARCHFLAGS="-arch x86_64"

# Set personal aliases, overriding those provided by oh-my-zsh libs,
# plugins, and themes. Aliases can be placed here, though oh-my-zsh
# users are encouraged to define aliases within the ZSH_CUSTOM folder.
# For a full list of active aliases, run `alias`.
#
# Example aliases
# alias zshconfig="mate ~/.zshrc"
# alias ohmyzsh="mate ~/.oh-my-zsh"

# 加载速度测试
# PS4=$'\\\011%D{%s%6.}\011%x\011%I\011%N\011%e\011'
# exec 3>&2 2>/tmp/zshstart.$$.log
# setopt xtrace prompt_subst

export PATH="/opt/hisi-linux/x86-arm/aarch64-himix100-linux/bin:$HOME/.local/bin:$HOME/.local/go/bin:$HOME/dev/go/bin:$HOME/.local/cmake/cmake-3.21.0/bin:$HOME/.local/protobuf/protobuf-3.17.3/bin:$HOME/.local/OpenCV/OpenCV-4.5.4/bin:$HOME/.local/node/node-v14.17.6/bin:$HOME/.local/keshub/2.0.6/keshub:$HOME/.local/FFmpeg/FFmpeg-n4.4/bin:$PATH"
export LD_LIBRARY_PATH="$HOME/.local/FFmpeg/FFmpeg-n4.4/lib:$HOME/.local/protobuf/protobuf-3.17.3/lib:$HOME/.local/OpenCV/OpenCV-4.5.4/lib:$HOME/.local/node/node-v14.17.6/lib:$HOME/.local/cJSON/cJSON-1.7.15/lib:$LD_LIBRARY_PATH"
export C_INCLUDE_PATH="$HOME/.local/OpenCV/OpenCV-4.5.4/include:$HOME/.local/FFmpeg/FFmpeg-n4.4/include:$HOME/.local/protobuf/protobuf-3.17.3/include:$HOME/.local/node/node-v14.17.6/include:$HOME/.local/cJSON/cJSON-1.7.15/include:$C_INCLUDE_PATH"
export CPLUS_INCLUDE_PATH=$C_INCLUDE_PATH
export LIBRARY_PATH=$LD_LIBRARY_PATH
export GOPATH="$HOME/dev/go"
export DOCKER_CLI_EXPERIMENTAL=enabled # 启动docker buildx
export LUA_LOCAL_PATH=$HOME/.local/luapkg
export LUA_PATH="$LUA_LOCAL_PATH/share/lua/5.1/?.lua;$LUA_LOCAL_PATH/share/lua/5.1/?/?.lua;;"

#ZSH_HIGHLIGHT_STYLES[suffix-alias]=fg=blue,underline
#ZSH_HIGHLIGHT_STYLES[precommand]=fg=blue,underline
ZSH_HIGHLIGHT_STYLES[arg0]=fg=green,bold

alias cman="man -M /usr/share/man/zh_CN"
#! ps aux | grep -q fetchmail && fetchmail &

#source /opt/poky/2.6/environment-setup-aarch64-poky-linux

# proxy polipo with sock5
function proxy() {
    case $1 in
    on)
        export no_proxy="localhost,127.0.0.1,localaddress,.localdomain.com"
        export http_proxy="http://127.0.0.1:8123"
        export https_proxy=$http_proxy
        export ftp_proxy=$http_proxy
        export rsync_proxy=$http_proxy
        export HTTP_PROXY=$http_proxy
        export HTTPS_PROXY=$http_proxy
        export FTP_PROXY=$http_proxy
        export RSYNC_PROXY=$http_proxy
    ;;
    off)
        unset no_proxy
        unset http_proxy
        unset https_proxy
        unset ftp_proxy
        unset rsync_proxy
        unset HTTP_PROXY
        unset HTTPS_PROXY
        unset FTP_PROXY
        unset RSYNC_PROXY
    ;;
    -h|--help)
        cat << EOF
Usage: ${0##*/} [-h|--help]
    -h|--help   show help
    on          proxy enable
    off         proxy disable
EOF
    ;;
    *)
        proxy on
        $@
        proxy off
    ;;
    esac
}

function skiplfs() {
    case $1 in
    on)
        export GIT_LFS_SKIP_SMUDGE=1
    ;;
    off)
        unset GIT_LFS_SKIP_SMUDGE
    ;;
    *)
        cat << EOF
Usage: ${0##*/} [-h|--help]
    on  "git modules skip lfs files"
    off "git modules not skip lfs files"
EOF
    ;;
    esac
}

skiplfs on
