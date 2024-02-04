# go-tushare-sdk

## 用法
主要是对 tushare 接口的 golang 封装

## 内容
涉及到所有 tushare 类型的接口类型，因为鄙人token 权限过期，部分接口未封装，如果有需要，可以联系我补充接口


## 使用样例
导入 sdk 后，使用代码补齐查询所有接口
```go
import (
    "github.com/yushikuann/go-tushare-sdk/client"
)

func main()  {
	token := "your tushare token"
	tushareClient := client.New(token)

	params := make(map[string]string)
	params["list_status"] = "L"
	var fields []string
	fieldStr := "ts_code,symbol,name,area,industry,market,list_date"
	fields = strings.Split(fieldStr, ",")
	
	data,err := tushareClient.StockBasic(params,fields)
	// your code
}

```

## 谨慎炒股，珍惜生活！！！
