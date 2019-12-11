package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ieee0824/zip-line/option"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	log.SetFlags(log.Llongfile)
	opt := new(option.Option)

	flag.StringVar(opt.Target.Pointer(), "t", "", opt.Target.Usage())
	flag.StringVar(opt.Output.Pointer(), "o", "", opt.Output.Usage())
	flag.BoolVar(opt.ForWin.Pointer(), "w", false, opt.ForWin.Usage())
	flag.Parse()

	stat, err := os.Stat(opt.Target.String())
	if err != nil {
		log.Fatalln(err)
	}
	archiveTargets := make([]string, 0, 1024)
	if !stat.IsDir() {
		archiveTargets = []string{opt.Target.String()}
	} else {
		archiveTargets, err = find(opt.Target.String())
		if err != nil {
			log.Fatalln(err)
		}
		for i, v := range archiveTargets {
			archiveTargets[i] = filepath.Clean(opt.Target.String() + "/" + v)
		}
	}

	if opt.ForWin {
		for i, v := range archiveTargets {
			sjis, err := toShiftJIS(v)
			if err != nil {
				log.Println(err)
				continue
			}
			archiveTargets[i] = sjis
		}
	}

	if err := archive(opt.Output.String(), archiveTargets); err != nil {
		log.Fatalln(err)
	}
}

func find(targetDir string) ([]string, error) {
	var paths []string
	err := filepath.Walk(targetDir,
		func(path string, info os.FileInfo, err error) error {
			rel, err := filepath.Rel(targetDir, path)
			if err != nil {
				return err
			}

			if info.IsDir() {
				paths = append(paths, fmt.Sprintf("%s/", rel))
				return nil
			}

			paths = append(paths, rel)

			return nil
		})

	if err != nil {
		return nil, err
	}

	return paths, nil
}

func archive(output string, paths []string) error {
	log.Println(paths)
	var compressedFile *os.File
	var err error
	if compressedFile, err = os.Create(output); err != nil {
		return err
	}
	defer compressedFile.Close()

	if err := compress(compressedFile, ".", paths); err != nil {
		return err
	}

	return nil
}

func compress(compressedFile io.Writer, targetDir string, files []string) error {
	w := zip.NewWriter(compressedFile)

	for _, filename := range files {
		filepath := fmt.Sprintf("%s/%s", targetDir, filename)
		info, err := os.Stat(filepath)
		if err != nil {
			return err
		}

		if info.IsDir() {
			continue
		}

		file, err := os.Open(filepath)
		if err != nil {
			return err
		}
		defer file.Close()

		hdr, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		hdr.Name = filename
		f, err := w.CreateHeader(hdr)
		if err != nil {
			return err
		}

		contents, _ := ioutil.ReadFile(filepath)
		_, err = f.Write(contents)
		if err != nil {
			return err
		}
	}
	if err := w.Close(); err != nil {
		return err
	}
	return nil
}

func transformEncoding(rawReader io.Reader, trans transform.Transformer) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(rawReader, trans))
	if err == nil {
		return string(ret), nil
	} else {
		return "", err
	}
}

func toShiftJIS(str string) (string, error) {
	return transformEncoding(strings.NewReader(str), japanese.ShiftJIS.NewEncoder())
}
