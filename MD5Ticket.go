package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
)

func getpassword(user string) {
	// use MD5 + salt as user password.
	userMD5 := md5.New()
	io.WriteString(userMD5, user)

	salt := "@#$%^&*()"
	buf := bytes.NewBufferString("")
	io.WriteString(buf, fmt.Sprintf("%x", userMD5.Sum(nil)))
	io.WriteString(buf, salt)

	p := md5.New()
	io.WriteString(p, buf.String())
	Password := fmt.Sprintf("%x", p.Sum(nil))
	fmt.Println(Password)
}
