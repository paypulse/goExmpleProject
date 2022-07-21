package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	//gin 선언 하기
	r := gin.Default()
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
