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
