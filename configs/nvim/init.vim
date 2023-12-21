" 文件备份, 修改时创建同名~文件作为备份文件
" if has("vms")
"     set nobackup
" else
"     set backup
"     if has("persistent_undo")
"         set undofile
"     endif
" endif

" 与vi的兼容性设置
set nocompatible

" 记录N条命令和匹配记录
set history=1000

" 行号和当前光标位置以及命令显示
set nu
set ruler
set showcmd

" 在状态行上显示补全匹配
set wildmenu

" 使<Esc>键生效更快
set ttimeout
set ttimeoutlen=100

" 如果末行被截短，显示@@@而不是隐藏整行
set display=truncate

" 1.查找时不循环跳转,2.输入部分查找模式时显示相应的匹配点,3.高亮显示匹配字符,4.高亮显示括号匹配
set nowrapscan
set incsearch
set hls
set showmatch

" 不把0开头的字符识别成八进制数
set nrformats-=octal

" 统一缩进为4空格
set expandtab
set tabstop=4
set softtabstop=4
set shiftwidth=4


" 设置鼠标可用
if has("mouse")
    set mouse=a
endif

" 文件探测和语法高亮
filetype plugin indent on
syntax on

" 为C程序设置自动缩进
" set cindent
set autoindent
set smartindent

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
au FileType make set noexpandtab

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
nnoremap <leader>" Bi"<Esc>Ea"<Esc>
nnoremap <F4> :w<cr>:!python %<cr>
nnoremap <leader>h ^
nnoremap <leader>e $
nnoremap <leader>wq :wq<CR>
nnoremap <leader>qo :tabonly<CR>
nnoremap <leader>w :w<CR>
nnoremap <leader>q :q<CR>
nnoremap x "_x
nnoremap d "_d
nnoremap D "_D
nnoremap <CR> o<Esc>

" vnoremap
vnoremap d "_d
nnoremap <leader>d ""d
nnoremap <leader>D ""D
vnoremap <leader>d ""d


set t_Co=256
set laststatus=2

"autocmd VimEnter * silent !tmux set status off
" autocmd VimLeave * silent !tmux set status on
set relativenumber
