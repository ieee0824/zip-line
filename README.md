# zip-line

linux/mac からwindows向けのzipを文字化けさせずに生成する

## インストール

### バイナリー
OSに合わせたバイナリーをダウンロードしてパスの通った場所にコピーする

https://github.com/ieee0824/zip-line/releases

### ソースコードから
Go 1.26.8 以上が必要です。

```
$ go get -u github.com/ieee0824/zip-line/cmd/zipl
```

## CI

PR と master への push で [mikoto](https://github.com/ieee0824/mikoto) による静的解析、`go vet ./...`、`go test ./...` を実行します。mikoto は関数の行数を 80 行までに制限し、`//mikoto:pure` を付けた関数の副作用も検査します。

参照透過性の検査対象は AES 鍵長の計算とファイルモードの変換（4 関数）です。mikoto の保守的な規則に合わせ、値型の入出力と定数だけで計算する関数に `//mikoto:pure` を付けています。ファイル I/O、乱数生成、ポインタやスライスを扱う処理は対象外です。

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
