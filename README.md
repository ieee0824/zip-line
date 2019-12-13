# zip-line

linux/mac からwindows向けのzipを文字化けさせずに生成する

## インストール

### バイナリー
OSに合わせたバイナリーをダウンロードしてパスの通った場所にコピーする
https://github.com/ieee0824/zip-line/releases

### ソースコードから
```
$ go get -u github.com/ieee0824/zip-line/cmd/zipl
```

## オプション

```
  -o string
        出力先を指定する
  -p string
        暗号化zipのパスワード(オプション)
        未指定のとき暗号化はされない
  -t string
        圧縮したいファイルまたはディレクトリのパスを指定する
  -w    windows向けのzipを生成する (オプション)
```