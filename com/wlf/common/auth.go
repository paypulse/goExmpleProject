package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TokenCreate struct {
	Email    string `json:Email`
	Password string `json:Password`
}

/**
* Token 처리
**/
func GetToken(email string, password string) (jwtToken interface{}) {

	//test계정 email/password
	fmt.Println("test email and password")
	postBody, _ := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})

	responseBody := bytes.NewBuffer(postBody)
	fmt.Println(responseBody)

	resp, err := http.Post("https://auth.useb.co.kr/oauth/token", "application/json", responseBody)
	if err != nil {
		panic(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	result := string(respBody)
	var m map[string]interface{}
	json.Unmarshal([]byte(result), &m)

	jwtToken = m["jwt"]

	return

}

/*
*
gorem 저장 전  token  공통
*
*/
func TokenTest() string {

	/***
	* 추후 gorm 으로 token , user name , email 저장 될 예정
	**/
	postBody, _ := json.Marshal(map[string]string{
		"email":    "ableman82@useb.co.kr",
		"password": "useb",
	})

	responseBody := bytes.NewBuffer(postBody)

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
	jwtToken := m["jwt"].(string)
	fmt.Println(jwtToken)

	return jwtToken

}
