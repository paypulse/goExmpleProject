package response

/**
* 진위 확인
 */

// Common response
type CommonResponse struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	ErrorCode     string `json:"error_code"'`
	TransactionId string `json:"transaction_id"`
}

// 사업자 등록증
type BusinessRegis struct {
	Success       bool         `json:"success"`
	Message       string       `json:"message"`
	Data          DataResponse `json:"data"`
	TransactionId string       `json:"transaction_id"`
}

type DataResponse struct {
	TaxTypeCode       string `json:"tax_type_code"`
	TaxTypeName       string `json:"tax_type_name"`
	ClosingDate       string `json:"closing_date"`
	TaxTypeChangeDate string `json:"tax_type_change_date"`
}
