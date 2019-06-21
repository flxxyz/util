package convert

import (
	. "github.com/flxxyz/util/convert"
	"testing"
)

func TestStrToInt(t *testing.T) {
	t.Logf("空格 = %#v", StrToInt(""))
	t.Logf("英文字符 = %#v", StrToInt("xyz"))
	t.Logf("中文字符 = %#v", StrToInt("你好"))
	t.Logf("英文特殊符号 = %#v", StrToInt("!@*#￥%"))
	t.Logf("中文特殊符号 = %#v", StrToInt("。’~！“，‘？"))
	t.Logf("八进制 = %#v", StrToInt("01023"))
	t.Logf("十六进制 = %#v", StrToInt("0x1023"))
	t.Logf("小数点 = %#v", StrToInt("0.1023"))
	t.Logf("'0' = %#v", StrToInt("0"))
	t.Logf("'-123' = %#v", StrToInt("-123"))
	t.Logf("'-0' = %#v", StrToInt("-0"))
	t.Logf("数字字符串 = %#v", StrToInt("123"))
}

func TestStrTo64(t *testing.T) {
	v := "64"
	c, _ := StrTo64(v)
	t.Logf("%s type = %T", v, c)
}

func TestStrToBool(t *testing.T) {
	v := "0"
	t.Logf("\"%s\" => %#v", v, StrToBool(v))

	v = "-1"
	t.Logf("\"%s\" => %#v", v, StrToBool(v))

	v = "1"
	t.Logf("\"%s\" => %#v", v, StrToBool(v))
}
