package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"log"
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

// test request Structure
type testRequest struct {
	Age  string `json:"Age"`
	Name string `json:"Name"`
}

// token 값
var jwtToken string

func main() {

	//gin 선언 하기
	r := gin.Default()

	////요청 처리
	//token
	r.GET("/getToken", httpHandler)

	//진위 확인 (신분증)
	r.POST("/checkIdcard", statusTest)

	//진위 확인 (운전 면허증)
	r.POST("/checkDriverCard", driverCard)

	r.Run(":8084")

}

// 핸들러 함수로 전달된 컨텍스트를 통해 클라이언트에게 http응답을 할 수 있다.
func httpHandler(c *gin.Context) {

	//token check
	createToken(c)

}

// 진위 확인
func statusTest(c *gin.Context) {
	checkIdcard(c)
}

// 진위 확인
func driverCard(c *gin.Context) {
	fmt.Println("***************** start check ****************")
	fmt.Println(c.Request.Body)
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	fmt.Println(value)

	var testReq testRequest
	json.Unmarshal(value, &testReq)

	fmt.Println("console. please ::", testReq)

	if err != nil {
		fmt.Println(err.Error())
	}

	//map으로 처리 된 경우
	var data map[string]interface{}
	json.Unmarshal([]byte(value), &data) // JSON을 Go언어 자료형으로 변환(여기서는 map으로 변환)

	//map to struct
	resultT := &testRequest{}
	mapstructure.Decode(data, &resultT)
	fmt.Println("resultT : ", resultT)

	c.JSON(http.StatusOK, gin.H{
		"name": data["name"],
		"age":  data["age"],
		"c":    string(value),
	})

	fmt.Println("data : ", data)

	real := []byte(string(value))
	err34 := json.Unmarshal(real, &resultT)
	if err34 != nil {
		log.Println(err34)
	}
	fmt.Println("please please : ", resultT.Age)

	///
	var t1 testRequest
	json.NewDecoder(c.Request.Body).Decode(&t1)
	fmt.Println("Test t1 ::", t1)

	var testReq1 testRequest
	json.Unmarshal([]byte(value), &testReq1)
	fmt.Println("testReq1 :", testReq1)

	fmt.Println("string req : ", string(value))
	fmt.Println("[]byte(value) : ", []byte(value))

	fmt.Println("***************** end check ****************")

}

type Post struct {
}

// 진위 확인 여부 (운전 면허증)

// 진위 확인 여부 (주민등록증)
func checkIdcard(c *gin.Context) {

	fmt.Println("*********** check id ************")

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

	//request body
	postBody1, _ := json.Marshal(map[string]string{
		"identity":            "8811211056911",
		"issueDate":           "20000301",
		"userName":            "홍길동",
		"secret_modeoptional": "true",
	})
	requestBody := bytes.NewBuffer(postBody1)

	bearers := "Bearer " + jwtToken.(string)

	r, err := http.NewRequest("POST", "https://api3.useb.co.kr/status/idcard", requestBody)
	if err != nil {
		panic(err)
	}
	r.Header.Set("Authorization", bearers)

	//client객체에서  request 실행
	client := &http.Client{}
	resp1, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp1.Body.Close()

	//response 체크
	respBodyCh, err := ioutil.ReadAll(resp1.Body)

	result12 := string(respBodyCh)
	var m1 map[string]interface{}
	json.Unmarshal([]byte(result12), &m1)
	success := m1["success"]
	message := m1["message"]
	error_code := m1["error_code"]
	transaction_id := m1["transaction_id"]

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success ":       success,
			"message":        message,
			"error_code":     error_code,
			"transaction_id": transaction_id,
		})

	}
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
