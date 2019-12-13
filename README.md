# zip-line

linux/mac からwindows向けのzipを文字化けさせずに生成する

## install

```
$ go get -u github.com/ieee0824/zip-line/cmd/zipl
```

## オプション

```
  -o string
        出力先を指定する
  -p string
        暗号化zipのパスワード(オプション)
  -t string
        圧縮したいファイルまたはディレクトリのパスを指定する
  -w    windows向けのzipを生成する (オプション)
```