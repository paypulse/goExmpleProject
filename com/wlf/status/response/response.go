package response

/**
* 진위 확인
 */

// idCard response
type IdCardResponse struct {
	Success       string `json:"success"`
	Message       string `json:"message"`
	ErrorCode     string `json:"error_code"'`
	TransactionId string `json:"transaction_id"`
}
