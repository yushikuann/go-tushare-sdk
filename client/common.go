package client

import "fmt"

const (
	PostMethod = "POST"
	GetMethod  = "GET"

	Domain = "http://api.tushare.pro"
)

var (
	ERR_PERMISSION  = fmt.Errorf("your token has no permission to use this api")
	ERR_ARGUEMENT   = fmt.Errorf("argurment error")
	ERR_DATE_FORMAT = fmt.Errorf("please input right date format YYYYMMDD")
)
