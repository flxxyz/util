package util

import (
	. "github.com/flxxyz/util"
	"testing"
)

func TestIsNumeric(t *testing.T) {
	t.Logf("空格'' = %#v", IsNumeric(""))
	t.Logf("英文字符'xyz' = %#v", IsNumeric("xyz"))
	t.Logf("中文字符'你好' = %#v", IsNumeric("你好"))
	t.Logf("英文特殊符号'!@#$&*(_+)=' = %#v", IsNumeric("!@#$&*(_+)="))
	t.Logf("中文特殊符号'。’~！“，‘？' = %#v", IsNumeric("。’~！“，‘？"))
	t.Logf("八进制'01023' = %#v", IsNumeric("01023"))
	t.Logf("十六进制'0x123' = %#v", IsNumeric("0x1023"))
	t.Logf("小数点'0.123' = %#v", IsNumeric("0.1023"))
	t.Logf("零值'0' = %#v", IsNumeric("0"))
	t.Logf("负数'-123' = %#v", IsNumeric("-123"))
	t.Logf("负零值'-0' = %#v", IsNumeric("-0"))
	t.Logf("数字字符串'123' = %#v", IsNumeric("123"))
}

func TestInArray(t *testing.T) {
	var s string

	needle := "4"
	hayStack := []interface{}{"0", "1", "2", "3", "4", "5", "6"}
	if InArray(needle, hayStack) {
		s = "存在"
	} else {
		s = "不存在"
	}
	t.Logf("%#v %s %s", hayStack, s, needle)

	needle = "9"
	if InArray(needle, hayStack) {
		s = "存在"
	} else {
		s = "不存在"
	}
	t.Logf("%#v %s %s", hayStack, s, needle)
}

func TestEOL(t *testing.T) {
	t.Logf("当前操作系统的换行符: %s", EOL())
}

func TestNow(t *testing.T) {
	t.Log(Now("20060102"))
	t.Log(Now("2006-01-02"))
	t.Log(Now("2006/01/02"))
	t.Log(Now("2006-01-02 15:04:05"))
}

func TestQQNumberCheck(t *testing.T) {
	var s string

	number := "10000"
	if QQNumberCheck(number) {
		s = "是"
	} else {
		s = "不是"
	}
	t.Logf("%s %sqq号", number, s)

	number = "TestQQNumberCheck"
	if QQNumberCheck(number) {
		s = "是"
	} else {
		s = "不是"
	}
	t.Logf("%s %sqq号", number, s)
}

func TestIpAddrCheck(t *testing.T) {
	var s string

	ip := "127.0.0.1"
	if IpAddrCheck(ip) {
		s = "合法"
	} else {
		s = "不合法"
	}
	t.Logf("%s ip地址%s", ip, s)

	ip = "0.0.0.0"
	if IpAddrCheck(ip) {
		s = "合法"
	} else {
		s = "不合法"
	}
	t.Logf("%s ip地址%s", ip, s)

	ip = "127.0.0.1:8080"
	if IpAddrCheck(ip) {
		s = "合法"
	} else {
		s = "不合法"
	}
	t.Logf("%s ip地址%s", ip, s)

	ip = "8.8.8.8"
	if IpAddrCheck(ip) {
		s = "合法"
	} else {
		s = "不合法"
	}
	t.Logf("%s ip地址%s", ip, s)
}

func TestResolveAddress(t *testing.T) {
	addr := ResolveAddress("", "2000")
	t.Log(addr)
}
