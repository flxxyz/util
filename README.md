# util
一些工具类库

## 依赖要求
没有

## 安装
使用`go`命令获取类库

```bash
go get github.com/flxxyz/util
```

## 例子
```go
package main

import (
    "fmt"
    "github.com/flxxyz/util"
)

func main() {
    //发起一个http请求
    body, _ := util.Request("GET", "http://baidu.com")
    fmt.Println(string(body))

    //判断字符串是否为数字
    if util.IsNumeric("99999") {
        fmt.Println("是的")
    }

    //判断元素是否存在集合内
    if util.InArray("4", []interface{}{"0", "1", "2", "3", "4", "5", "6"}) {
        fmt.Println("存在！")
    }

    //格式化当前时间
    now := util.Now("2006-01-02 15:04:05")
    fmt.Printf("现在时间是 %s", now)

    //qq号判断
    if util.QQNumberCheck("10000") {
        fmt.Println("是的")
    }

    //判断ip地址合法
    if util.IpAddrCheck("127.0.0.1") {
        fmt.Println("是的")
    }
}
```

## 文档
[文档点这里](http://godoc.org/github.com/flxxyz/util)

## 版权
util包在MIT License下发布。有关详细信息，请参阅LICENSE。