package crypto

import (
	"com.lh.service/tools"
	"crypto/hmac"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"io"
	"strings"
	"time"
)

// 加密密码
func Encrypt(p string, pathname string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		config, _ := tools.Language(pathname)
		return "", errors.New(config["bcrypt"]["error"])
	}
	return string(password), err
}

// 对加密校验
func Decrypt(p string, enP string, pathname string) error {
	err := bcrypt.CompareHashAndPassword([]byte(enP), []byte(p))
	if err != nil {
		config, _ := tools.Language(pathname)
		return errors.New(config["bcrypt"]["errPW"])
	}
	return err
}

// UUID生成解码
func UUID(pathname string) (string, error) {
	v4 := uuid.NewV4()
	id, err := uuid.FromString(v4.String())
	if err != nil {
		config, _ := tools.Language(pathname)
		return "", errors.New(config["uuid"]["error"])
	}
	return id.String(), err
}

// md5加密
func Md5(value string) string {
	if value == "" {
		now := time.Now()
		dates := []string{string(now.UnixNano()), string(now.Nanosecond())}
		value = strings.Join(dates, "")
	}
	md := md5.New()
	io.WriteString(md, value)
	id := md.Sum(nil)
	return hex.EncodeToString(id)
}

// base64转码
func EnBase(value string) string {
	input := []byte(value)
	return base64.StdEncoding.EncodeToString(input)
}

// base64解码
func DeBase(value string, pathname string) (string, error) {
	base, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		config, _ := tools.Language(pathname)
		return "", errors.New(config["base64"]["error"])
	}
	return string(base), err
}

// hmac加密
func Hmac(value string, key string, dir string) string {
	hma := hmac.New(md5.New, []byte(key))
	hma.Write([]byte(value))
	return hex.EncodeToString(hma.Sum([]byte("")))
}
