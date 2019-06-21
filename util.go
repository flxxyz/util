package util

import (
	"net"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"time"
)

// 处理相对的ip地址与端口
func ResolveAddress(host string, port string) string {
	if port == "0" {
		port = "8080"
	}

	if envPort := os.Getenv("PORT"); envPort != "" {
		port = ":" + envPort
	}

	return host + ":" + port
}

// 判断是否存在
func IsExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// php函数in_array实现
func InArray(needle interface{}, hayStack []interface{}) bool {
	for _, v := range hayStack {
		if needle == v {
			return true
		}
	}

	return false
}

// ip地址检查
func IpAddrCheck(ip string) bool {
	var (
		match bool
		err   error
	)

	if a := net.ParseIP(ip); a == nil {
		return false
	}

	if match, err = regexp.MatchString("(2(5[0-5]{1}|[0-4]\\d{1})|[0-1]?\\d{1,2})(\\.(2(5[0-5]{1}|[0-4]\\d{1})|[0-1]?\\d{1,2})){3}", ip); err != nil {
		return false
	}

	if match == false {
		return false
	}

	return true
}

// qq号检查
func QQNumberCheck(qq string) (is bool) {
	is, _ = regexp.MatchString("^[1-9][0-9]{4,11}$", qq)

	return
}

//判断数字字符串
func IsNumeric(v string) bool {
	if v == "0" {
		return true
	}

	n, _ := strconv.Atoi(v)
	if n != 0 {
		return true
	}

	return false
}

// 操作系统的换行符
func EOL() string {
	switch runtime.GOOS {
	case "windows":
		return "\\"
	case "darwin":
		return "/"
	case "linux":
		return "/"
	default:
		return "/"
	}
}

// 格式化当前时间
func Now(format string) string {
	return time.Unix(time.Now().Unix(), 0).Format(format)
}
