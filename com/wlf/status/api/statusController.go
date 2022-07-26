package api

import (
	"encoding/json"
	Common "example.com/ginExampleTest/com/wlf/common"
	Request "example.com/ginExampleTest/com/wlf/status/request"
	Service "example.com/ginExampleTest/com/wlf/status/service"
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

	jwtToken := Common.TokenTest()
	body := c.Request.Body
	value, _ := ioutil.ReadAll(body)
	var driverCard Request.DriverCard
	json.Unmarshal([]byte(value), &driverCard)

	resultRe := Service.DriverCardStateService(jwtToken, driverCard)

	c.JSON(http.StatusOK, gin.H{
		"success":        resultRe.Success,
		"message":        resultRe.Message,
		"error_code":     resultRe.ErrorCode,
		"transaction_id": resultRe.TransactionId,
	})

}

// 진위 확인 :: 한국 여권
func PassportStatus(c *gin.Context) {

	jwtToken := Common.TokenTest()
	body := c.Request.Body
	value, _ := ioutil.ReadAll(body)

	var passport Request.PassPort
	json.Unmarshal([]byte(value), &passport)

	resultRe := Service.PassportStateService(jwtToken, passport)

	c.JSON(http.StatusOK, gin.H{
		"success":        resultRe.Success,
		"message":        resultRe.Message,
		"error_code":     resultRe.ErrorCode,
		"transaction_id": resultRe.TransactionId,
	})

}

// 진위 확인 :: 외국인 여권
func PassportOverSeas(c *gin.Context) {

	jwtToken := Common.TokenTest()
	body := c.Request.Body
	value, _ := ioutil.ReadAll(body)

	var passportOverSeas Request.PassPortOverSeas
	json.Unmarshal([]byte(value), &passportOverSeas)

	resultRe := Service.PassportOverSeasService(jwtToken, passportOverSeas)

	c.JSON(http.StatusOK, gin.H{
		"success":        resultRe.Success,
		"message":        resultRe.Message,
		"error_code":     resultRe.ErrorCode,
		"transaction_id": resultRe.TransactionId,
	})

}

// 진위 확인 :: 외국인 등록증
func Alien(c *gin.Context) {

	jwtToken := Common.TokenTest()
	body := c.Request.Body
	value, _ := ioutil.ReadAll(body)

	var alien Request.Alien
	json.Unmarshal([]byte(value), &alien)

	resultRe := Service.Alien(jwtToken, alien)

	c.JSON(http.StatusOK, gin.H{
		"success":        resultRe.Success,
		"message":        resultRe.Message,
		"error_code":     resultRe.ErrorCode,
		"transaction_id": resultRe.TransactionId,
	})

}

// 진위 확인 :: 사업자 등록 및 휴폐업 조회
func BusinessRegis(c *gin.Context) {
	jwtToken := Common.TokenTest()
	body := c.Request.Body
	value, _ := ioutil.ReadAll(body)

	var alien Request.Alien
	json.Unmarshal([]byte(value), &alien)

	resultRe := Service.BusinessRegis(jwtToken, alien)

	c.JSON(http.StatusOK, gin.H{
		"success":        resultRe.Success,
		"message":        resultRe.Message,
		"data":           resultRe.Data,
		"transaction_id": resultRe.TransactionId,
	})

}
