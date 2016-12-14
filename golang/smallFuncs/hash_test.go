package main

import (
	"fmt"
	"testing"
)

func Test_MakeHash(t *testing.T) {

	sha1OfString := MakeSha1OfString(filePath)
	sha1OfFile, _ := MakeSha1OfFile(filePath)
	sha256OfString := MakeSha256OfString(filePath)
	sha256OfFile, _ := MakeSha256OfFile(filePath)
	MD5OfString := MakeMD5OfString(filePath)
	MD5OfFile, _ := MakeMD5OfFile(filePath)
	etag, _ := GetEtag(filePath)

	fmt.Printf("sha1OfString=%v\n", sha1OfString)
	fmt.Printf("sha1OfFile=%v\n\n", sha1OfFile)
	fmt.Printf("sha256OfString=%v\n", sha256OfString)
	fmt.Printf("sha256OfFile=%v\n\n", sha256OfFile)
	fmt.Printf("MD5OfString=%v\n", MD5OfString)
	fmt.Printf("MD5OfFile=%v\n\n", MD5OfFile)
	fmt.Printf("etag=%v\n", etag)
}
