package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func EncodeMd5Str(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

func EncodeMd5File(filename string) (string, error) {
	file, err := os.Open(filename)
	if !IsEmpty(err) {
		return "", err
	}
	m := md5.New()
	if _, err = io.Copy(m, file); !IsEmpty(err) {
		return "", err
	}
	return hex.EncodeToString(m.Sum(nil)), nil
}
