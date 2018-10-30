package dot

type neovim struct {
	application string
	filename string
	path string
	template string
}

func (self Neovim) String() string { return Render(self.template, self.settings) }
func (self Neovim) Path() string { return (path + filename) }

func Init(values Settings) Configuration {
	// TODO: Load config templates from file 
	return  Configuration{
	application: "neovim",
	configFiles: []ConfigFile{
		ConfigFile{
			filename: "init.vim",
			path: "~/.config/nvim/",
			template: initvimTemplate(),
		},
	}
}

func initvimTemplate() string {
		return `" Vim Configuration
"==============================================================================
" Neovim: ~/.local/share/nvim/plugged
call plug#begin()
  " General Vim Settings
  Plug 'tpope/vim-sensible'
  " Theme
  Plug 'tyrannicaltoucan/vim-deep-space'
  " Autocompletion
  Plug 'prabirshrestha/asyncomplete-file.vim'
  Plug 'prabirshrestha/async.vim'
  Plug 'prabirshrestha/asyncomplete.vim'
  Plug 'prabirshrestha/asyncomplete-buffer.vim'
  Plug 'prabirshrestha/asyncomplete-gocode.vim'
  " Snippets
  Plug 'honza/vim-snippets'
  " Alignment
  Plug 'junegunn/vim-easy-align'
  " File Manager
  Plug 'scrooloose/nerdtree', { 'on':  'NERDTreeToggle' }
  " Fuzzy
  Plug 'junegunn/fzf', { 'dir': '~/.fzf', 'do': './install --all' }
  " Specific Programming Language Plugins
  "== Markdown Plugins
  Plug 'junegunn/goyo.vim', { 'for': 'markdown' }
  "== C++
  Plug 'octol/vim-cpp-enhanced-highlight', { 'for': 'cpp' }
  "== Go Related Plugins
  Plug 'fatih/vim-go', { 'do': ':GoUpdateBinaries' }
  Plug 'nsf/gocode', { 'tag': 'v.20150303', 'rtp': 'vim' }
call plug#end()

" Convert tabs to spaces, use 2 spaces for tabs
set expandtab
set tabstop=2
set shiftwidth=2

" Jump to the last position when reopening a file
if has("autocmd")
  au BufReadPost * if line("'\"") > 1 && line("'\"") <= line("$") | exe "normal! g'\"" | endif
endif

" Nerdtree Settings - Key Assignment
nnoremap <C-a> :tabprevious<CR>
nnoremap <C-d> :tabnext<CR>
nnoremap <C-s> :tabnew<CR>:NERDTree<CR>

" Theme
set background=dark
set termguicolors
colorscheme deep-space

function! g:FuckThatMatchParen ()
    if exists(":NoMatchParen")
        :NoMatchParen
    endif
endfunction

augroup plugin_initialize
    autocmd!
    autocmd VimEnter * call FuckThatMatchParen()
augroup END

" Aliasing
command Q q
command Qw wq
command Wq wq


let g:go_fmt_command = 'goimports'

let g:gofmt_formatters = [
\   { 'cmd': 'gofmtrlx', 'args': ['-s', '-w'] },
\   { 'cmd': 'goimports', 'args': ['-w'] },
\   { 'cmd': 'gotypeconv', 'args': ['-w'] },
\ ]


imap <c-space> <Plug>(asyncomplete_force_refresh)
set completeopt-=preview

let g:asyncomplete_smart_completion = 1
let g:asyncomplete_auto_popup = 1
let g:asyncomplete_remove_duplicates = 1

autocmd! CompleteDone * if pumvisible() == 0 | pclose | endif


call asyncomplete#register_source(asyncomplete#sources#gocode#get_source_options({
    \ 'name': 'gocode',
    \ 'whitelist': ['go'],
    \ 'completor': function('asyncomplete#sources#gocode#completor'),
    \ 'config': {
    \    'gocode_path': expand('~/Development/bin/gocode')
    \  },
    \ }))

au User asyncomplete_setup call asyncomplete#register_source(asyncomplete#sources#file#get_source_options({
    \ 'name': 'file',
    \ 'whitelist': ['*'],
    \ 'priority': 10,
    \ 'completor': function('asyncomplete#sources#file#completor')
    \ }))

call asyncomplete#register_source(asyncomplete#sources#buffer#get_source_options({
    \ 'name': 'buffer',
    \ 'whitelist': ['*'],
    \ 'blacklist': ['go'],
    \ 'completor': function('asyncomplete#sources#buffer#completor'),
    \ }))


" Use 'tab' key to select completions.  Default is arrow keys.
"inoremap <expr> <Tab> pumvisible() ? "\<C-n>" : "\<Tab>"
"inoremap <expr> <S-Tab> pumvisible() ? "\<C-p>" : "\<S-Tab>"

" Use tab to trigger auto completion.  Default suggests completions as you type.
"let g:completor_auto_trigger = 0

let g:completor_gocode_binary = '/usr/bin/gocode'
set number


set textwidth=80
au BufRead,BufNewFile *.md setlocal textwidth=80

" Return to previous line if file previously opened
if has("autocmd")
  au BufReadPost * if line("'\"") > 0 && line("'\"") <= line("$")
    \| exe "normal! g'\"" | endif
endif
`}
}
