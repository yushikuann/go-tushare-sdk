package client

// MoneyFlow 获取股权质押统计数据
func (api *TuShare) MoneyFlow(params map[string]string, fields []string) (*APIResponse, error) {
	// Check param
	_, hasTradeDate := params["trade_date"]
	_, hasStartDate := params["start_date"]

	// trade_date & start_date required
	if (!hasTradeDate && !hasStartDate) || (hasTradeDate && hasStartDate) {
		return nil, ERR_ARGUEMENT
	}

	body := map[string]interface{}{
		"api_name": "moneyflow",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}
