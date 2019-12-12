package main

import (
	"flag"
	"log"

	"github.com/ieee0824/zip-line/archive"
	"github.com/ieee0824/zip-line/option"
)

func main() {
	log.SetFlags(log.Llongfile)
	opt := new(option.Option)

	flag.StringVar(opt.Target.Pointer(), "t", "", opt.Target.Usage())
	flag.StringVar(opt.Output.Pointer(), "o", "", opt.Output.Usage())
	flag.BoolVar(opt.ForWin.Pointer(), "w", false, opt.ForWin.Usage())
	flag.StringVar(opt.Password.Pointer(), "p", "", opt.Password.Usage())
	flag.Parse()

	if err := archive.Archive(opt); err != nil {
		log.Fatalln(err)
	}
}
