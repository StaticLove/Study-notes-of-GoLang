package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"os"
	"path"
)

const (
	BLOCK_BITS = 22 // Indicate that the blocksize is 4M
	BLOCK_SIZE = 1 << BLOCK_BITS
)

var (
	currentPath, _ = os.Getwd()
	filePath       = path.Clean(currentPath + "../../../resource/picture/Ulquiorra_8.jpg")
)

func BlockCount(fsize int64) int {
	return int((fsize + (BLOCK_SIZE - 1)) >> BLOCK_BITS)
}

func CalSha1(b []byte, r io.Reader) ([]byte, error) {

	h := sha1.New()
	_, err := io.Copy(h, r)
	if err != nil {
		return nil, err
	}
	return h.Sum(b), nil
}

// 对文件计算 qetag
func GetEtag(s string) (etag string, err error) {
	file, err := os.Open(s)
	if err != nil {
		return
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return
	}

	fsize := fi.Size()
	blockCnt := BlockCount(fsize)
	sha1Buf := make([]byte, 0, 21)

	if blockCnt <= 1 { // file size <= 4M
		sha1Buf = append(sha1Buf, 0x16)
		sha1Buf, err = CalSha1(sha1Buf, file)
		if err != nil {
			return
		}
	} else { // file size > 4M
		sha1Buf = append(sha1Buf, 0x96)
		sha1BlockBuf := make([]byte, 0, blockCnt*20)
		for i := 0; i < blockCnt; i++ {
			body := io.LimitReader(file, BLOCK_SIZE)
			sha1BlockBuf, err = CalSha1(sha1BlockBuf, body)
			if err != nil {
				return
			}
		}
		sha1Buf, _ = CalSha1(sha1Buf, bytes.NewReader(sha1BlockBuf))
	}
	etag = base64.URLEncoding.EncodeToString(sha1Buf)

	return
}

// 对字符串计算 sha1
func MakeSha1OfString(s string) string {
	h := sha1.New()
	h.Write([]byte(s))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// 对文件计算 sha1
func MakeSha1OfFile(s string) (string, error) {
	file, err := os.Open(s)
	if err != nil {
		return "", err
	}

	defer file.Close()

	h := sha1.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

// 对字符串计算 sha256
func MakeSha256OfString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// 对文件计算 sha256
func MakeSha256OfFile(s string) (string, error) {
	file, err := os.Open(s)
	if err != nil {
		return "", err
	}

	defer file.Close()

	h := sha256.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

// 对字符串计算 MD5
func MakeMD5OfString(s string) string {
	h := md5.New()
	h.Write([]byte(s))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// 对文件计算 MD5
func MakeMD5OfFile(s string) (string, error) {
	file, err := os.Open(s)
	if err != nil {
		return "", nil
	}

	defer file.Close()

	h := md5.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return "", nil
	}

	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
