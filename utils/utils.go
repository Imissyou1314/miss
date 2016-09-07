package utils

/**
 * 工具模块
 */
import (
    "crypto/md5"
    "encoding/hex"
)

/**
 *   进行MD5加密
 */
func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}
