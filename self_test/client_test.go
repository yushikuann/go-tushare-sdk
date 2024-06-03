package self_test

import (
	"fmt"
	"github.com/yushikuann/go-tushare-sdk/client"
	"github.com/yushikuann/go-tushare-sdk/config"
	"strings"
	"testing"
	"time"
)

var token = ""

func getToken() string {
	if token == "" {
		token = config.GetConfig("/..")["TS_TOKEN"]
	}

	return token
}

func TestClient(t *testing.T) {
	share := client.New(getToken())
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

type FiledMapping struct {
	theType string
	trans   string
}

func TestStkMins(t *testing.T) {
	share := client.New(getToken())
	params := make(map[string]string)
	params["ts_code"] = "000005.SZ"
	params["start_date"] = "2015-01-06 15:00:00"
	params["end_date"] = "2015-06-01 15:00:00"
	params["freq"] = "1min"
	params["limit"] = "5"
	data, err := share.StkMins(params, []string{})
	if err != nil {
		t.Error(err)
	}
	if len(data.Data.Items) != 5 {
		t.Errorf("data count not has %d pieces", len(data.Data.Items))
	}

	if len(data.Data.Items) > 0 {
		mapping := []FiledMapping{
			{"string", "symbol"},
			{"string", "tradeDate"},
			{"float64", "open"},
			{"float64", "sClose"},
			{"float64", "high"},
			{"float64", "low"},
			{"float64", "vol"},
			{"float64", "amount"},
		}
		for _, item := range data.Data.Items {
			for i2, slot := range mapping {
				var ok bool

				switch slot.theType {
				case "float64":
					_, ok = item[i2].(float64)
				case "string":
					_, ok = item[i2].(string)
				default:
					t.Errorf("Noun type: %s", slot.theType)
				}
				if !ok {
					fmt.Println(item)
					t.Errorf("No.%d: %s format is failed, the symbol %s", i2, slot.trans, item[0])
				}
			}
		}
	}
}
