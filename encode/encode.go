package encode

import (
	"io"
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func transformEncoding(rawReader io.Reader, trans transform.Transformer) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(rawReader, trans))
	if err == nil {
		return string(ret), nil
	} else {
		return "", err
	}
}

func ToShiftJIS(str string) (string, error) {
	return transformEncoding(strings.NewReader(str), japanese.ShiftJIS.NewEncoder())
}
