package api

import (
	"encoding/json"
	Common "example.com/ginExampleTest/com/wlf/common"
	Request "example.com/ginExampleTest/com/wlf/status/request"
	Service "example.com/ginExampleTest/com/wlf/status/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

/**
* 진위 확인 여부 controller
**/

// 진위 확인 :: 주민 등록증
func IdCardStatus(c *gin.Context) {

	jwtToken := Common.TokenTest()
	/**
	Real request
	*/
	body := c.Request.Body
	value, _ := ioutil.ReadAll(body)
	var idCardInfo Request.IdCard
	json.Unmarshal([]byte(value), &idCardInfo)

	resultRe := Service.IdCardStatusService(jwtToken, idCardInfo)

	c.JSON(http.StatusOK, gin.H{
		"success":        resultRe.Success,
		"message":        resultRe.Message,
		"error_code":     resultRe.ErrorCode,
		"transaction_id": resultRe.TransactionId,
	})

}

// 진위 확인 :: 운전 면허증
func DriverCardStatus(c *gin.Context) {

	token := Common.TokenTest()

	fmt.Println("token Test : ", token)

}
