// Package utils
/**
  * @Author:manguanghui
  * @Date: 2021/11/22
  * @Desc:
**/
package utils

import (
	"bytes"
	"compress/zlib"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/segmentio/ksuid"
)

var iw, _ = NewIdWorker(1)

// GoId 获取雪花算法ID
func GoId() int64 {
	id, _ := iw.NextId()
	return id
}

// GeneratedID 获取ksuid
func GeneratedID() string {
	return ksuid.New().String()
}

// GetOutboundIP 获取本地IP
func GetOutboundIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

func String(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

// NowDateTime 系统时间
func NowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func NowDate() string {
	return time.Now().Format("2006-01-02")
}

func NowParseDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

// NowUnix 时间戳
func NowUnix() int64 {
	return time.Now().Unix()
}

// Base64Encode 编码
func Base64Encode(str []byte) string {
	return base64.StdEncoding.EncodeToString(str)
}

// Base64Decode 解码
func Base64Decode(str string) []byte {
	decoded, _ := base64.StdEncoding.DecodeString(str)
	return decoded
}

// IsString 判断字符串是否为空
func IsString(str string) bool {
	return strings.Trim(str, " ") == ""
}

func IsNotString(str string) bool {
	return !IsString(str)
}

// Md5
// @param str 字符串
func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

// SHA256 算法
func SHA256(message []byte) string {
	bytes2 := sha256.Sum256(message)
	hashcode2 := hex.EncodeToString(bytes2[:])
	return hashcode2
}

// FileSuffix 获取文件后缀
func FileSuffix(path string) string {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[i:]
		}
	}
	return ""
}

// CreateFile 创建新文件
func CreateFile(filepath string) {
	//  如果没有文件目录就创建一个
	if _, err := os.Stat(filepath); err != nil {
		if !os.IsExist(err) {
			_ = os.MkdirAll(filepath, os.ModePerm)
		}
	}
}

// ZlibUnCompress 解压zlib
func ZlibUnCompress(compressSrc []byte) ([]byte, error) {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	_, err := io.Copy(&out, r)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
