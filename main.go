package main

import (
	"fmt"
	"gthub.com/skingford/gouti/hash"
)

func main()  {
	s:=hash.StringMd5("123456")
	fmt.Println(s)
	d := hash.Md5String("e10adc3949ba59abbe56e057f20f883e")
	fmt.Println(d)
}
