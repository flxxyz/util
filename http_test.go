package util

import (
	. "github.com/flxxyz/util"
	"testing"
)

func TestRequest(t *testing.T) {
	body, _ := Request("GET", "http://baidu.com")
	t.Log(string(body))
}
