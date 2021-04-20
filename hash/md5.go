package hash

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"unsafe"
)

func StringMd5(s string) string {
	md5S := md5.New()
	md5S.Write([]byte(s))
	return hex.EncodeToString(md5S.Sum(nil))
}

func Md5String(s string) string {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	return String(decoded)
}

func String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
