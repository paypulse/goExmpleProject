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

// 진위 확인 :: 한국여권
func PassportStateService(jwtToken interface{}, passportCard Request.PassPort) Response.CommonResponse {
	fmt.Println("jwtToken :::", jwtToken)
	fmt.Println("DriverCard :::", passportCard)

	/*bearer token */
	bearers := "Bearer " + jwtToken.(string)
	postBody1, _ := json.Marshal(passportCard)
	requestBody := bytes.NewBuffer(postBody1)

	respone, err := http.NewRequest("POST", "https://api3.useb.co.kr/status/passport", requestBody)
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

// 진위 확인 :: 외국인 여권
func PassportOverSeasService(jwtToken interface{}, passPortOverseas Request.PassPortOverSeas) Response.CommonResponse {
	fmt.Println("jwtToken :::", jwtToken)
	fmt.Println("DriverCard :::", passPortOverseas)

	/*bearer token */
	bearers := "Bearer " + jwtToken.(string)
	postBody1, _ := json.Marshal(passPortOverseas)
	requestBody := bytes.NewBuffer(postBody1)

	respone, err := http.NewRequest("POST", "https://api3.useb.co.kr/status/passport-overseas", requestBody)
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

// 진위 확인 :: 외국인 등록증
func Alien(jwtToken interface{}, alien Request.Alien) Response.CommonResponse {
	fmt.Println("jwtToken :::", jwtToken)
	fmt.Println("DriverCard :::", alien)

	/*bearer token */
	bearers := "Bearer " + jwtToken.(string)
	postBody1, _ := json.Marshal(alien)
	requestBody := bytes.NewBuffer(postBody1)

	respone, err := http.NewRequest("POST", "https://api3.useb.co.kr/status/alien", requestBody)
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

// 진위 확인 :: 사업자 등록 및 휴폐업 조회
func BusinessRegis(jwtToken interface{}, business Request.BusinessRegistraction) Response.BusinessRegis {
	fmt.Println("jwtToken :::", jwtToken)
	fmt.Println("DriverCard :::", business)

	/*bearer token */
	bearers := "Bearer " + jwtToken.(string)
	postBody1, _ := json.Marshal(business)
	requestBody := bytes.NewBuffer(postBody1)

	respone, err := http.NewRequest("POST", "https://api3.useb.co.kr/status-doc/business-registration", requestBody)
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

	var m1 map[string]interface{}
	json.Unmarshal([]byte(result), &m1)
	fmt.Println("m1 :::", m1)

	var responInfo Response.BusinessRegis
	json.Unmarshal([]byte(result), &responInfo)

	fmt.Println("responseInfo :::", responInfo)

	return responInfo

}
