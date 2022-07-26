package service

import (
	"bytes"
	"encoding/json"
	Request "example.com/ginExampleTest/com/wlf/status/request"
	Response "example.com/ginExampleTest/com/wlf/status/response"
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
* 진위 확인 service
**/

// 진위 확인 : 주민 등록증
func IdCardStatusService(jwtToken interface{}, idCard Request.IdCard) Response.CommonResponse {

	/*bearer token */
	bearers := "Bearer " + jwtToken.(string)
	fmt.Println("totalToken : ", bearers)
	fmt.Println("idCardInfo : ", idCard)

	/*request Body*/
	postBody1, _ := json.Marshal(idCard)
	requestBody := bytes.NewBuffer(postBody1)

	fmt.Println("request Body : ", requestBody)

	respone, err := http.NewRequest("POST", "https://api3.useb.co.kr/status/idcard", requestBody)
	if err != nil {
		panic(err)
	}
	respone.Header.Set("Authorization", bearers)

	//client 객체에서 request실행
	client := &http.Client{}
	resAgain, error := client.Do(respone)
	if error != nil {
		panic(error)
	}
	defer resAgain.Body.Close()

	//response 체크
	resBodyCheck, _ := ioutil.ReadAll(resAgain.Body)
	result := string(resBodyCheck)

	//var m1 map[string]interface{}
	//json.Unmarshal([]byte(result), &m1)

	var responInfo Response.CommonResponse
	json.Unmarshal([]byte(result), &responInfo)

	return responInfo
}

// 진위 확인 :: 운전 면허증
func DriverCardStateService(jwtToken interface{}, driverCard Request.DriverCard) Response.CommonResponse {

	fmt.Println("jwtToken :::", jwtToken)
	fmt.Println("DriverCard :::", driverCard)

	/*bearer token */
	bearers := "Bearer " + jwtToken.(string)
	postBody1, _ := json.Marshal(driverCard)
	requestBody := bytes.NewBuffer(postBody1)

	respone, err := http.NewRequest("POST", "https://api3.useb.co.kr/status/driver", requestBody)
	if err != nil {
		panic(err)
	}
	respone.Header.Set("Authorization", bearers)

	//client 객체에서 request실행
	client := &http.Client{}
	resAgain, error := client.Do(respone)
	if error != nil {
		panic(error)
	}
	defer resAgain.Body.Close()

	//response 체크
	resBodyCheck, _ := ioutil.ReadAll(resAgain.Body)
	result := string(resBodyCheck)

	//var m1 map[string]interface{}
	//json.Unmarshal([]byte(result), &m1)

	var responInfo Response.CommonResponse
	json.Unmarshal([]byte(result), &responInfo)

	return responInfo

}
