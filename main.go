package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// auth
type Auth struct {
	success        string
	message        string
	jwt            string
	expires_in     string
	transaction_id string
}

// requestField
type RequestField struct {
	email    string
	password string
}

// 진위 확인 request
type statusReq struct {
	identity       string
	issueDate      string
	userName       string
	secret_Mode    bool
	birthDate      string
	licenseNo      string
	passportNo     string
	expirationDate string
	nationaliy     string
	biz_no         string
}

// 진위 확인 response
type statusRes struct {
	success        string
	message        string
	transaction_id string
}

// token 값
var jwtToken string

func main() {

	//gin 선언 하기
	r := gin.Default()

	////요청 처리
	//token
	r.GET("/getToken", httpHandler)

	//진위 확인
	r.GET("/checkIdcard", httpHandler)

	r.Run(":8084")
}

// 핸들러 함수로 전달된 컨텍스트를 통해 클라이언트에게 http응답을 할 수 있다.
func httpHandler(c *gin.Context) {

	//token check
	createToken(c)
	//진위 확인
	checkIdcard(c)
}

// 진위 확인 여부
func checkIdcard(c *gin.Context) {

	//request header
	//Authorization
	//bears := "Bearer " + jwtToken

	//request body

}

// 토큰 확인
func createToken(c *gin.Context) {
	//	requestField := RequestField{"ableman82@useb.co.kr", "useb"}
	//	pbytes, _ := xml.Marshal(requestField)
	//	buff := bytes.NewBuffer(pbytes)

	postBody, _ := json.Marshal(map[string]string{
		"email":    "ableman82@useb.co.kr",
		"password": "useb",
	})

	responseBody := bytes.NewBuffer(postBody)
	fmt.Println(responseBody)

	//reqBody := bytes.NewBufferString("Post plain text")
	resp, err := http.Post("https://auth.useb.co.kr/oauth/token", "application/json", responseBody)

	if err != nil {
		panic(err)
	}

	//response 체크
	respBody, err := ioutil.ReadAll(resp.Body)
	result := string(respBody)
	var m map[string]interface{}
	json.Unmarshal([]byte(result), &m)
	jwtToken := m["jwt"]
	fmt.Println(jwtToken)

	if err == nil {
		//응답 처리
		c.JSON(http.StatusOK, gin.H{
			"result": jwtToken,
		})
	}
}
