package self_test

import (
	"fmt"
	"github.com/yushikuann/go-tushare-sdk/client"
	"strings"
	"testing"
	"time"
)

var token = "XXX"

func TestClient(t *testing.T) {
	share := client.New(token)
	params := make(map[string]string)
	params["list_status"] = "L"
	var fields []string
	fieldStr := "ts_code,symbol,name,area,industry,market,list_date"
	fields = strings.Split(fieldStr, ",")
	data, _ := share.StockBasic(params, fields)

	fmt.Println(data.Data.Items)

	for _, item := range data.Data.Items[1:20] {
		fmt.Println(time.Parse("20060102", item[6].(string)))
	}
}
