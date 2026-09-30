package archive

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ieee0824/zip-line/option"
	"github.com/ieee0824/zip-line/zip"
)

func TestArchiveRejectsInputAsOutput(t *testing.T) {
	for _, kind := range []string{"same path", "alternate path", "symlink output", "hardlink output", "symlink input", "directory input"} {
		t.Run(kind, func(t *testing.T) {
			dir := t.TempDir()
			source := filepath.Join(dir, "original.txt")
			const contents = "irreplaceable input"
			if err := os.WriteFile(source, []byte(contents), 0600); err != nil {
				t.Fatal(err)
			}
			input, output := source, source
			alias := filepath.Join(dir, "alias.txt")
			switch kind {
			case "alternate path":
				output = dir + string(os.PathSeparator) + "." + string(os.PathSeparator) + "original.txt"
			case "symlink output", "symlink input":
				if err := os.Symlink(source, alias); err != nil {
					t.Fatal(err)
				}
				if kind == "symlink input" {
					input = alias
				} else {
					output = alias
				}
			case "hardlink output":
				if err := os.Link(source, alias); err != nil {
					t.Fatal(err)
				}
				output = alias
			case "directory input":
				input = dir
			}
			opt := new(option.Option)
			*opt.Target.Pointer(), *opt.Output.Pointer() = input, output
			if err := Archive(opt); err == nil || !strings.Contains(err.Error(), "same file") {
				t.Fatalf("expected same-file error, got %v", err)
			}
			got, err := os.ReadFile(source)
			if err != nil || string(got) != contents {
				t.Fatalf("input changed: contents=%q, error=%v", got, err)
			}
		})
	}
}

func TestArchiveSeparateOutput(t *testing.T) {
	for _, existing := range []bool{false, true} {
		dir := t.TempDir()
		source, output := filepath.Join(dir, "input.txt"), filepath.Join(dir, "out.zip")
		const contents = "hello"
		if err := os.WriteFile(source, []byte(contents), 0600); err != nil {
			t.Fatal(err)
		}
		if existing {
			if err := os.WriteFile(output, []byte(strings.Repeat("old data", 1000)), 0600); err != nil {
				t.Fatal(err)
			}
		}
		opt := new(option.Option)
		*opt.Target.Pointer(), *opt.Output.Pointer() = source, output
		if err := Archive(opt); err != nil {
			t.Fatal(err)
		}
		r, err := zip.OpenReader(output)
		if err != nil {
			t.Fatal(err)
		}
		if len(r.File) != 1 || r.File[0].UncompressedSize64 != uint64(len(contents)) {
			t.Fatal("unexpected archive contents")
		}
		r.Close()
		got, err := os.ReadFile(source)
		if err != nil || string(got) != contents {
			t.Fatalf("input changed: contents=%q, error=%v", got, err)
		}
	}
}
