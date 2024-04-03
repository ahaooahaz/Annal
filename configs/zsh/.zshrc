#=====================#
# Directory Shortcuts #
#=====================#
hash -d config=$XDG_CONFIG_HOME
hash -d cache=$XDG_CACHE_HOME
hash -d data=$XDG_DATA_HOME

hash -d zdot=$ZDOTDIR

#=====================#
# P10k Instant Prompt #
#=====================#
include -f "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh"

#=========#
# Plugins #
#=========#
include -f "${HOME}/.zcomet/bin/zcomet.zsh"

zcomet load ohmyzsh lib {completion,clipboard}.zsh
zcomet load ohmyzsh plugins/autojump
zcomet load ohmyzsh plugins/history
zcomet load ohmyzsh plugins/history-substring-search
zcomet load ohmyzsh plugins/git
zcomet load tj/git-extras etc git-extras-completion.zsh
zcomet load hlissner/zsh-autopair

zcomet compinit

zcomet load Aloxaf/fzf-tab
# zstyle ':fzf-tab:*' fzf-bindings 'tab:accept'
zstyle ':fzf-tab:complete:kill:argument-rest' fzf-preview 'ps --pid=$word -o cmd --no-headers -w -w'
zstyle ':fzf-tab:complete:kill:argument-rest' fzf-flags '--preview-window=down:3:wrap'
zstyle ':fzf-tab:complete:kill:*' popup-pad 0 3

# zcomet load zsh-users/zsh-autosuggestions
# zcomet load zsh-users/zsh-syntax-highlighting
zcomet load zdharma-continuum/fast-syntax-highlighting

zcomet load romkatv/powerlevel10k
###############
# Key Binding #
###############

bindkey -r '^['  # Unbind [Esc]    (default: vi-cmd-mode)

bindkey '^A'   beginning-of-line   # [Ctrl-A]
bindkey '^E'   end-of-line         # [Ctrl-E]
bindkey '^Z'   undo                # [Ctrl-Z]
bindkey '^Y'   redo                # [Ctrl-Y]
bindkey ' '    magic-space         # [Space]     Trigger history expansion
bindkey '^[^M' self-insert-unmeta  # [Alt-Enter] Insert newline
bindkey '^R' history-incremental-search-backward  # [Ctrl-R] history search

#=========#
# Configs #
#=========#

# zsh misc
setopt auto_cd               # simply type dir name to `cd`
setopt auto_pushd            # make `cd` behave like pushd
setopt pushd_ignore_dups     # don't pushd duplicates
setopt pushd_minus           # exchange the meanings of `+` and `-` in pushd
setopt ksh_option_print      # make `setopt` output all options
setopt extended_glob         # extended globbing

# time (zsh built-in)
TIMEFMT="\
%J   %U  user %S system %P cpu %*E total
avg shared (code):         %X KB
avg unshared (data/stack): %D KB
total (sum):               %K KB
max memory:                %M MB
page faults from disk:     %F
other page faults:         %R"

# fzf
export FZF_DEFAULT_OPTS='--ansi --height=60% --reverse --cycle --bind=tab:accept'

# bat
export BAT_THEME='OneHalfDark'

# man
export MANPAGER='sh -c "col -bx | bat -pl man --theme=Monokai\ Extended"'
export MANROFFOPT='-c'

# zsh history
setopt hist_ignore_all_dups  # no duplicates
setopt hist_save_no_dups     # don't save duplicates
setopt hist_ignore_space     # no commands starting with space
setopt hist_reduce_blanks    # remove all unneccesary spaces
setopt share_history         # share history between sessions
HISTFILE=~/.zsh_history
HISTSIZE=1000000  # number of commands that are loaded
SAVEHIST=1000000  # number of commands that are stored

autoload -Uz colors && colors  # provide color variables (see `which colors`)

# fix echo <fe0f> chars when rime input a emoji
# Ref: https://github.com/rime/rime-emoji/issues/8
setopt COMBINING_CHARS

#=========#
# Scripts #
#=========#
include -f "${ZDOTDIR:-${HOME}}/.p10k.zsh"

apps=(kubectl helm)
for app in ${apps}; do
    if [ $(type ${app} >/dev/null 2>&1; echo $?) -eq 0 ]; then
        source <(${app} completion zsh)
    fi
done

export LS_COLORS="$LS_COLORS:*.tar=1;31:*.tgz=1;31:*.arj=1;31:*.taz=1;31:*.lzh=1;31:*.lzma=1;31:*.tlz=1;31:*.txz=1;31:*.zip=1;31:*.z=1;31:*.Z=1;31:*.dz=1;31:*.gz=1;31:*.lz=1;31:*.xz=1;31:*.bz2=1;31:*.bz=1;31:*.tbz=1;31:*.tbz2=1;31:*.tz=1;31:*.deb=1;31:*.rpm=1;31:*.jar=1;31:*.war=1;31:*.ear=1;31:*.sar=1;31:*.rar=1;31:*.ace=1;31:*.zoo=1;31:*.cpio=1;31:*.7z=1;31:*.rz=1;31"

