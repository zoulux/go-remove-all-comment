# go-remove-all-comment

可以一键去除项目 go 文件中的注释

```bash
$ go install github.com/zoulux/go-remove-all-comment

$ go-remove-all-comment -h                                                                                                                                                                                                                                   (1s)[18:44:53]
Usage of go-remove-all-comment:
-file-dir string
support dir, each file
-file-name string
support one file
-out-dir string
export dir, default empty is origin file
```

- -file-name 指定一个 .go 文件，优先级比 file-dir 高
- -file-dir  指定一个文件夹，会递归下面所有的 .go 文件
- -out-dir   指定一个导出的文件夹，如果不指定，则会覆盖原来的文件

例子ð：
```bash
$ go-remove-all-comment --file-dir /Users/jake/workspace/book/ 
```