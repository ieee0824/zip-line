package option

import (
	"github.com/ieee0824/getenv"
)

type target string

func (t *target) String() string {
	return string(*t)
}

func (_ target) Usage() string {
	switch getenv.String("LANG") {
	case "ja_JP.UTF-8":
		return "圧縮したいファイルまたはディレクトリのパスを指定する"
	default:
		return "Specify the path of the file or directory you want to compress"
	}
}

func (t *target) Pointer() *string {
	return (*string)(t)
}

type output string

func (_ output) Usage() string {
	switch getenv.String("LANG") {
	case "ja_JP.UTF-8":
		return "出力先を指定する"
	default:
		return "Specify the output destination"
	}
}

func (o *output) Pointer() *string {
	return (*string)(o)
}

func (o *output) String() string {
	return string(*o)
}

type forWin bool

func (_ forWin) Usage() string {
	switch getenv.String("LANG") {
	case "ja_JP.UTF-8":
		return "windows向けのzipを生成する (オプション)"
	default:
		return "Generate zip for windows (optional)"
	}
}

func (f *forWin) Pointer() *bool {
	return (*bool)(f)
}

type password string

func (_ password) Usage() string {
	switch getenv.String("LANG") {
	case "ja_JP.UTF-8":
		return "暗号化zipのパスワード(オプション)"
	default:
		return "Encryption zip password (optional)"
	}
}

func (p *password) Pointer() *string {
	return (*string)(p)
}

func (p *password) String() string {
	return string(*p)
}

type Option struct {
	Target   target
	Output   output
	ForWin   forWin
	Password password
}
