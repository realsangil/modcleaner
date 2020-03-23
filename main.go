package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

const cacheDir = "cache"

var verbose = true
var gopath = os.Getenv("GOPATH")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\nmodcleaner pkg0 pkg1 pkg2...\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.BoolVar(&verbose, "v", false, "verbose mode")
	flag.Parse()

	for i := 0; i < flag.NArg(); i++ {
		mod := parseModule(flag.Arg(i))
		log.Printf("%#v\n", mod)
		if err := deleteMod(mod); err != nil {
			log.Fatalln(err)
		}
		log.Printf("deleted cache: %s\n", flag.Arg(i))
	}
	fmt.Println("module caches delete complete")
}

func deleteMod(mod Module) error {
	if err := deletePkg(mod); err != nil {
		return errors.Wrap(err, mod.String())
	}

	if err := deleteCache(mod); err != nil {
		return errors.Wrap(err, mod.String())
	}

	if err := deleteVCS(mod); err != nil {
		return errors.Wrap(err, mod.String())
	}

	return nil
}

func deletePkg(mod Module) error {
	p := path.Join(gopath, "pkg", "mod", mod.String())
	if err := os.RemoveAll(p); err != nil {
		return errors.Wrap(err, "remove go pkg")
	}
	return nil
}

func deleteCache(mod Module) error {
	version := "*"
	if mod.Version != "" {
		version = mod.Version + ".*"
	}
	p := path.Join(gopath, "pkg", "mod", cacheDir, "download", mod.Package, "@v", version)
	files, err := filepath.Glob(p)
	if err != nil {
		return errors.Wrap(err, "file list with pattern")
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			return errors.Wrapf(err, "remove cache: %s", f)
		}
	}
	return nil
}

func deleteVCS(mod Module) error {
	vcsDir := path.Join(gopath, "pkg", "mod", cacheDir, "vcs")
	vcsPath := path.Join(vcsDir, "*.info")
	files, err := filepath.Glob(vcsPath)
	if err != nil {
		return errors.Wrap(err, "vcs list with pattern")
	}

	for _, f := range files {
		buf, err := ioutil.ReadFile(f)
		if err != nil {
			return errors.Wrapf(err, "read file: %s", f)
		}
		if strings.Contains(string(buf), mod.Package) {
			if err := deleteVCSWithPath(f); err != nil {
				return errors.Wrapf(err, "remove vsc: %q", f)
			}
			break
		}
	}

	return nil
}

func deleteVCSWithPath(p string) error {
	filename := path.Base(p)
	dir := path.Dir(p)
	hash := strings.TrimSuffix(filename, ".info")
	hashDir := path.Join(dir, hash)
	if err := os.RemoveAll(hashDir); err != nil {
		return errors.Wrapf(err, "delete hash: %q", hashDir)
	}

	hashInfo := path.Join(dir, hash+".info")
	if err := os.Remove(hashInfo); err != nil {
		return errors.Wrapf(err, "delete hash info: %q", hashInfo)
	}

	hashLock := path.Join(dir, hash+".lock")
	if err := os.Remove(hashLock); err != nil {
		return errors.Wrapf(err, "delete hash lock: %q", hashLock)
	}

	return nil
}
