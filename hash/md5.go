/*
 * @Author: kingford
 * @Date: 2021-04-20 12:39:08
 * @LastEditTime: 2021-04-20 12:39:24
 */

import (
	"crypto/md5"
	"encoding/hex"
)

func StringMd5(s string) string {
	md5S := md5.New()
	md5S.Write([]byte(s))
	return hex.EncodeToString(md5S.Sum(nil))
}
