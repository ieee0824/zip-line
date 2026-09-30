package archive

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/ieee0824/zip-line/encode"
	"github.com/ieee0824/zip-line/option"
	"github.com/ieee0824/zip-line/tree"
	"github.com/ieee0824/zip-line/zip"
)

func addFile(zw *zip.Writer, file string, opt *option.Option) error {
	stat, err := os.Stat(file)
	if err != nil {
		return err
	}

	r, err := os.Open(file)
	if err != nil {
		return err
	}
	var w io.Writer
	rootDir := "/" + filepath.Base(opt.Target.String())
	key := rootDir + strings.TrimPrefix(file, opt.Target.String())
	if opt.ForWin {
		sjis, err := encode.ToShiftJIS(key)
		if err != nil {
			return err
		}
		key = sjis
	}
	switch opt.Password.String() {
	case "":
		var err error
		w, err = zw.Create(key, stat)
		if err != nil {
			return err
		}
	default:
		var err error
		w, err = zw.Encrypt(key, stat, opt.Password.String(), zip.StandardEncryption)
		if err != nil {
			return err
		}
	}
	if !stat.IsDir() {
		if _, err := io.Copy(w, r); err != nil {
			return err
		}
	}

	return nil
}

func Archive(opt *option.Option) error {
	files, err := tree.Tree(opt.Target.String())
	if err != nil {
		return err
	}
	zipFile, err := openOutput(opt.Output.String(), files)
	if err != nil {
		return err
	}
	defer zipFile.Close()
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Flush()
	defer zipWriter.Close()

	for _, file := range files {
		if err := addFile(zipWriter, file, opt); err != nil {
			return err
		}
	}

	return nil
}

// openOutput checks the opened file's identity before truncating it so that
// alternate paths, symlinks, and hard links cannot overwrite an input file.
func openOutput(path string, inputs []string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	ok := false
	defer func() {
		if !ok {
			f.Close()
		}
	}()
	outputInfo, err := f.Stat()
	if err != nil {
		return nil, err
	}
	for _, input := range inputs {
		inputInfo, err := os.Stat(input)
		if err != nil {
			return nil, err
		}
		if os.SameFile(inputInfo, outputInfo) {
			return nil, fmt.Errorf("input %q and output %q refer to the same file", input, path)
		}
	}
	if err := f.Truncate(0); err != nil {
		return nil, err
	}
	ok = true
	return f, nil
}
