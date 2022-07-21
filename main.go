package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {

	//gin 선언 하기
	r := gin.Default()

	createToken()

	//요청 처리
	r.GET("/", httpHandler)
	r.Run(":8082")
}

// 핸들러 함수로 전달된 컨텍스트를 통해 클라이언트에게 http응답을 할 수 있다.
func httpHandler(c *gin.Context) {

	//응답 처리
	c.JSON(http.StatusOK, gin.H{
		"message": "helloworld",
	})

}

/*
*
토큰 생성
*/
func createToken() {
	reqBody := bytes.NewBufferString("Post plain text")
	resp, err := http.Post("https://auth.useb.co.kr/oauth/token", "text/plain", reqBody)

	if err != nil {
		panic(err)
	}

	//response 체크
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		print(str)
	}
}
