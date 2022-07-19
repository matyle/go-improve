#golang #vim 
[[golang 中的string和byte数组]]
[[coc-vim配置golang环境]]
[[vim-readme]]


# 配置 git 相关插件

### diffview
```vim
" autocmd BufWritePost * GitGutter
nnoremap <LEADER>gf :GitGutterFold<CR>
nnoremap <LEADER>gh :GitGutterPreviewHunk<CR>
nnoremap <LEADER>g- :GitGutterPrevHunk<CR>
nnoremap <LEADER>g= :GitGutterNextHunk<CR>
"Diff
nnoremap <LEADER>gd :DiffviewOpen<CR> 
"cloSe
nnoremap <LEADER>gs :DiffviewClose<CR>
"hisTory
nnoremap <LEADER>gt :DiffviewFileHistory<CR>
"togglefile
nnoremap <LEADER>go :DiffviewToggleFiles<CR>
"gitlens
```

# 书籍《go 语言精进之路》
## golang 第二部分 项目结构，代码风格与标识符命名


### 第5条 使用公认且广泛使用的项目结构
官方 golang 的项目结构：
```bash
❯ tree -LF 1
./
├── crypto/
├── exp/
├── mod/
├── net/
├── oauth2/
├── sync/
├── sys/
├── text/
├── tools/
├── tour/
└── xerrors/
```

```bash
❯ tree -LF 1
./
├── AUTHORS
├── CONTRIBUTING.md
├── CONTRIBUTORS
├── LICENSE
├── PATENTS
├── README.md
├── benchmark/
├── blog/
├── cmd/
├── codereview.cfg
├── container/
├── copyright/
├── cover/
├── go/
├── go.mod
├── go.sum
├── godoc/
├── gopls/
├── imports/
├── internal/
├── playground/
├── present/
├── refactor/
└── txtar/
```
- 构建二进制可执行文件为目的的 Go 项目结构



