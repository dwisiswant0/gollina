package main

import (
	"io"
	"os"

	"archive/zip"
	"encoding/base64"
	"path/filepath"

	"golang.org/x/text/encoding/unicode"
)

func shouldSkip(substr string) bool {
	for _, v := range skips {
		if v == substr {
			return true
		}
	}

	return false
}

func convertUtf8ToUtf16LE(message string) (string, error) {
	utf16le := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	utfEncoder := utf16le.NewEncoder()
	ut16LeEncodedMessage, err := utfEncoder.String(message)

	return ut16LeEncodedMessage, err
}

func powershellEncode(message string) (string, error) {
	utf16LEEncodedMessage, err := convertUtf8ToUtf16LE(message)
	if err != nil {
		return "", err
	}

	input := []uint8(utf16LEEncodedMessage)
	return base64.StdEncoding.EncodeToString(input), nil
}

func zipSource(src, dst string) error {
	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Method = zip.Deflate
		header.Name, err = filepath.Rel(filepath.Dir(src), path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			header.Name += "/"
		}

		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(headerWriter, f)
		return err
	})
}
