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

#==================#
# Plugins          #
#==================#
include -f "${HOME}/.zcomet/bin/zcomet.zsh"

zcomet load ohmyzsh lib {completion,clipboard}.zsh
zcomet load ohmyzsh plugins/autojump
zcomet load ohmyzsh plugins/history
zcomet load ohmyzsh plugins/history-substring-search
zcomet load tj/git-extras etc git-extras-completion.zsh
zcomet load hlissner/zsh-autopair

zcomet compinit

zcomet load Aloxaf/fzf-tab
# zstyle ':fzf-tab:*' fzf-bindings 'tab:accept'
zstyle ':fzf-tab:complete:kill:argument-rest' fzf-preview 'ps --pid=$word -o cmd --no-headers -w -w'
zstyle ':fzf-tab:complete:kill:argument-rest' fzf-flags '--preview-window=down:3:wrap'
zstyle ':fzf-tab:complete:kill:*' popup-pad 0 3

# zcomet load zsh-users/zsh-autosuggestions
# ZSH_AUTOSUGGEST_MANUAL_REBIND=true
# ZSH_AUTOSUGGEST_CLEAR_WIDGETS+=(
#     qc-accept-line
# )
# ZSH_AUTOSUGGEST_PARTIAL_ACCEPT_WIDGETS+=(
#     qc-forward-subword
#     qc-forward-shellword
# )

zcomet load zdharma-continuum/fast-syntax-highlighting
# unset 'FAST_HIGHLIGHT[chroma-man]'  # chroma-man will stuck history browsing

zcomet load romkatv/powerlevel10k

#=========#
# Configs #
#=========#

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

apps=(kubectl helm)
for app in ${apps}; do
    if [ $(type ${app} >/dev/null 2>&1; echo $?) -eq 0 ]; then
        source <(${app} completion zsh)
    fi
done

autoload -Uz colors && colors  # provide color variables (see `which colors`)

# fix echo <fe0f> chars when rime input a emoji.
setopt COMBINING_CHARS
include -f "${ZDOTDIR:-${HOME}}/.p10k.zsh"