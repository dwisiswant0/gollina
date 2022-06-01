package main

import "embed"

var (
	//go:embed all:res
	res embed.FS
	opt *options

	pages = map[string]string{
		"/": "res/exploit.gohtml",
	}
	skips = []string{".", "exploit.gohtml"}
)
