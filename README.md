# go-tushare-sdk

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