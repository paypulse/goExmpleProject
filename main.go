package main

import (
	"encoding/json"
	Common "example.com/ginExampleTest/com/wlf/common"
	Controller "example.com/ginExampleTest/com/wlf/status/api"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// test request Structure
type testRequest struct {
	Age  string `json:"Age"`
	Name string `json:"Name"`
}

func main() {

	//gin 선언 하기
	r := gin.Default()
	////요청 처리
	//token
	r.GET("/getToken", httpHandler)

	//진위 확인 (신분증)
	r.POST("/checkIdcard", Controller.IdCardStatus)

	//진위 확인 (운전 면허증)
	r.POST("/checkDriverCard", Controller.DriverCardStatus)

	//진위 확인 (한국 여권)
	r.POST("/checkPassport", Controller.PassportStatus)

	//진위 확인 (외국인 여권)
	r.POST("/checkPassportOverSeas", Controller.PassportOverSeas)

	//진위 확인 (외국인 등록증)
	r.POST("/checkAlien", Controller.Alien)

	//사업자 등록 및 휴 폐업 조회
	r.POST("/checkBusinessRegi", Controller.BusinessRegis)

	r.Run(":8084")

}

// 핸들러 함수로 전달된 컨텍스트를 통해 클라이언트에게 http응답을 할 수 있다.
func httpHandler(c *gin.Context) {
	//token check
	createToken(c)
}

// 토큰 확인
func createToken(c *gin.Context) {
	////////////////request 확인
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	//token
	var tokenInfo Common.TokenCreate
	json.Unmarshal([]byte(value), &tokenInfo)
	jwtToken := Common.GetToken(tokenInfo.Email, tokenInfo.Password)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"token": jwtToken,
	})

}
