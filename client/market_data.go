package client

import (
	"fmt"
)

// Daily 获取股票行情数据, 日线
func (api *TuShare) Daily(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]

	// ts_code & trade_date required
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, ERR_ARGUEMENT
	}

	if dateFormat := IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]); !dateFormat {
		return nil, ERR_DATE_FORMAT
	}

	body := map[string]interface{}{
		"api_name": "daily",
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}

	return api.postData(body)
}

// Weekly 获取股票行情数据, 周线
func (api *TuShare) Weekly(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]

	// ts_code & trade_date required
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, ERR_ARGUEMENT
	}

	if dateFormat := IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]); !dateFormat {
		return nil, ERR_DATE_FORMAT
	}

	body := map[string]interface{}{
		"api_name": "weekly",
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}

	return api.postData(body)
}

// Monthly 获取A股月线数据
func (api *TuShare) Monthly(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]

	// ts_code & trade_date required
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, ERR_ARGUEMENT
	}

	if dateFormat := IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]); !dateFormat {
		return nil, ERR_DATE_FORMAT
	}

	body := map[string]interface{}{
		"api_name": "monthly",
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}

	return api.postData(body)
}

// DailyBasic 获取全部股票每日重要的基本面指标，可用于选股分析、报表展示等
func (api *TuShare) DailyBasic(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]

	// ts_code & trade_date required
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, ERR_ARGUEMENT
	}

	if dateFormat := IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]); !dateFormat {
		return nil, ERR_DATE_FORMAT
	}

	body := map[string]interface{}{
		"api_name": "daily_basic",
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}

	return api.postData(body)
}

// AdjFactor 获取股票复权因子，可提取单只股票全部历史复权因子，也可以提取单日全部股票的复权因子
func (api *TuShare) AdjFactor(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]

	// ts_code & trade_date required
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, ERR_ARGUEMENT
	}

	if dateFormat := IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]); !dateFormat {
		return nil, ERR_DATE_FORMAT
	}

	body := map[string]interface{}{
		"api_name": "adj_factor",
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}

	return api.postData(body)
}

// Suspend 获取股票每日停复牌信息
func (api *TuShare) Suspend(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["suspend_date"]
	_, hasResumeDate := params["resume_date"]

	// ts_code & trade_date & resume_date only required
	argsCount := 0
	if hasTsCode {
		argsCount++
	}
	if hasTradeDate {
		argsCount++
	}
	if hasResumeDate {
		argsCount++
	}

	if argsCount != 1 {
		return nil, fmt.Errorf("Need one argument among ts_code, suspend_date, resume_date")
	}

	if dateFormat := IsDateFormat(params["suspend_date"], params["resume_date"]); !dateFormat {
		return nil, ERR_DATE_FORMAT
	}

	body := map[string]interface{}{
		"api_name": "suspend",
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}

	return api.postData(body)
}

// StkMins 股票行情数据, 分钟
// params: ts_code / start_date / end_date / freq / offset / limit
// params example: 601965.SH / 2024-05-27 09:30:00 / 2024-05-27 15:00:00 / 1min / 0 / 10
// return: code股票代码 / tradetime交易时间 / open开盘价 / close收盘价 / high最高价 / low最低价 / vol成交量 / amount成交额(元)
// return example: [[601965.SH 2024-05-28 15:00:00 18.69 18.69 18.69 18.69 40700 760683]]
func (api *TuShare) StkMins(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "stk_mins",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}
