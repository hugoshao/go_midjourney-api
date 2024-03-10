package Util

import (
	"crypto/md5"
	"encoding/hex"
)

// 根据提示词（string）生成一个唯一的ID
func GenerateID(hint string) string {
	// 使用 MD5 哈希函数计算哈希值
	hasher := md5.New()
	hasher.Write([]byte(hint))
	hashed := hasher.Sum(nil)

	// 将哈希值转换为十六进制字符串
	hashStr := hex.EncodeToString(hashed)
	SendLog(hashStr)
	return hashStr
}
