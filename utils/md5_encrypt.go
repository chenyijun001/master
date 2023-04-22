package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
)

// md5String 根据字符串进行加密
func md5String(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

// MD5ByByte 根据文件内容进行加密
func MD5ByByte(file *multipart.FileHeader) (string, error) {
	md5file, err2 := file.Open()
	if err2 != nil {
		return "", err2
	}
	md5bytes, err2 := io.ReadAll(md5file)
	hasher := md5.New()
	if _, err := hasher.Write(md5bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

//// MD5ByByte 根据文件内容进行加密
//func MD5ByByte(content []byte) (string, error) {
//	hasher := md5.New()
//	if _, err := hasher.Write(content); err != nil {
//		return "", err
//	}
//
//	return hex.EncodeToString(hasher.Sum(nil)), nil
//}

// Md5File 根据真实文件的路径计算文件的 MD5 哈希值
func Md5File(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
