set nocompatible

set history=1000

" 设置行号
set nu
"
" 统一缩进为4空格
set expandtab
set tabstop=4
set softtabstop=4
set shiftwidth=4

" 设置语法高亮
filetype on
syntax on
syntax enable

" 设置鼠标可用
set mouse=a

" 为C程序设置自动缩进
" set cindent
set autoindent
set smartindent

" 查找时自动跳转
set incsearch

" 查找结果高亮
set hls

" 高亮显示匹配的括号
set showmatch

" 同步主选区"*寄存器与匿名寄存器""
" set clipboard=unnamed

" 同步剪切板寄存器"+与匿名寄存器""
set clipboard=unnamedplus

" fencview
set encoding=utf8
set langmenu=zh_CN.UTF-8
language message zh_CN.UTF-8
set fileencodings=ucs-bom,utf-8,cp936,gb18030,big5,euc-jp,euc-kr,latin1

" set background=dark

" 打开文件跳转至上次退出位置
if has("autocmd")
  au BufReadPost * if line("'\"") > 1 && line("'\"") <= line("$") | exe "normal! g'\"" | endif
endif

" makefile tab键不发生转换
autocmd FileType make set noexpandtab

" 修改vim的颜色 
set background=dark
set t_Co=256

" 光标所在行列高亮
set cursorcolumn
set cursorline

" <leader>
let mapleader=","
" nnoremap
" <CR>: 回车键
" <Esc>: Esc键
nnoremap <F4> :w<cr>:!python %<cr>
nnoremap <CR> i<CR><Esc>$
nnoremap <leader>h ^
nnoremap <leader>e $
nnoremap <leader>wq :wq<CR>
nnoremap <leader>qo :tabonly<CR>
nnoremap <leader>w :w<CR>
nnoremap <leader>q :q<CR>
nnoremap x "_x
nnoremap d "_d
nnoremap D "_D

" vnoremap
vnoremap d "_d
nnoremap <leader>d ""d
nnoremap <leader>D ""D
vnoremap <leader>d ""d


set t_Co=256
set laststatus=2
python3 from powerline.vim import setup as powerline_setup
python3 powerline_setup()
python3 del powerline_setup

autocmd VimEnter * silent !tmux set status off
autocmd VimLeave * silent !tmux set status on

" custom EOF.
" set the runtime path to include Vundle and initialize
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()
" alternatively, pass a path where Vundle should install plugins
"call vundle#begin('~/some/path/here')

" let Vundle manage Vundle, required
Plugin 'VundleVim/Vundle.vim'

" The following are examples of different formats supported.
" Keep Plugin commands between vundle#begin/end.
" plugin on GitHub repo
Plugin 'tpope/vim-fugitive'
Plugin 'davidhalter/jedi-vim'
" plugin from http://vim-scripts.org/vim/scripts.html
" Plugin 'L9'
" Git plugin not hosted on GitHub
Plugin 'git://git.wincent.com/command-t.git'
" git repos on your local machine (i.e. when working on your own plugin)
Plugin 'file:///home/gmarik/path/to/plugin'
" The sparkup vim script is in a subdirectory of this repo called vim.
" Pass the path to set the runtimepath properly.
Plugin 'rstacruz/sparkup', {'rtp': 'vim/'}
" Install L9 and avoid a Naming conflict if you've already installed a
" different version somewhere else.
" Plugin 'ascenator/L9', {'name': 'newL9'}

" All of your Plugins must be added before the following line
call vundle#end()            " required
filetype plugin indent on    " required
" To ignore plugin indent changes, instead use:
"filetype plugin on
"
" Brief help
" :PluginList       - lists configured plugins
" :PluginInstall    - installs plugins; append `!` to update or just :PluginUpdate
" :PluginSearch foo - searches for foo; append `!` to refresh local cache
" :PluginClean      - confirms removal of unused plugins; append `!` to auto-approve removal
"
" see :h vundle for more details or wiki for FAQ
" Put your non-Plugin stuff after this line

